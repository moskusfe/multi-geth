// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-test/deep"
)

// StateProcessor is a basic Processor, which takes care of transitioning
// state from one point to another.
//
// StateProcessor implements Processor.
type StateProcessor struct {
	config *params.ChainConfig // Chain configuration options
	bc     *BlockChain         // Canonical block chain
	engine consensus.Engine    // Consensus engine used for block rewards
}

// NewStateProcessor initialises a new StateProcessor.
func NewStateProcessor(config *params.ChainConfig, bc *BlockChain, engine consensus.Engine) *StateProcessor {
	return &StateProcessor{
		config: config,
		bc:     bc,
		engine: engine,
	}
}

// Process processes the state changes according to the Ethereum rules by running
// the transaction messages using the statedb and applying any rewards to both
// the processor (coinbase) and any included uncles.
//
// Process returns the receipts and logs accumulated during the process and
// returns the amount of gas that was used in the process. If any of the
// transactions failed to execute due to insufficient gas it will return an error.
func (p *StateProcessor) Process(block *types.Block, statedb *state.StateDB, cfg vm.Config) (types.Receipts, []*types.Log, uint64, error) {
	var (
		receipts types.Receipts
		usedGas  = new(uint64)
		header   = block.Header()
		allLogs  []*types.Log
		gp       = new(GasPool).AddGas(block.GasLimit())
	)
	// Mutate the block and state according to any hard-fork specs
	if p.config.DAOForkSupport && p.config.DAOForkBlock != nil && p.config.DAOForkBlock.Cmp(block.Number()) == 0 {
		misc.ApplyDAOHardFork(statedb)
	}
	// Iterate over and process the individual transactions
	for i, tx := range block.Transactions() {
		statedb.Prepare(tx.Hash(), block.Hash(), i)
		receipt, _, err := ApplyTransaction(p.config, p.bc, nil, gp, statedb, header, tx, usedGas, cfg)
		if err != nil {
			return nil, nil, 0, err
		}
		receipts = append(receipts, receipt)
		allLogs = append(allLogs, receipt.Logs...)
	}
	// Finalize the block, applying any consensus engine specific extras (e.g. block rewards)
	p.engine.Finalize(p.bc, header, statedb, block.Transactions(), block.Uncles())

	return receipts, allLogs, *usedGas, nil
}

var errEVMMismatch = errors.New("mismatched evm results")

func evmMismatchErr(f string, in ...interface{}) error {
	return fmt.Errorf("%s: %s", errEVMMismatch, fmt.Sprintf(f, in...))
}
func isEVMMismatchError(err error) bool {
	return strings.HasPrefix(err.Error(), errEVMMismatch.Error())
}

func compareEVMs(config *params.ChainConfig, bc ChainContext, author *common.Address, gp *GasPool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *uint64, cfg vm.Config) error {
	s1 := statedb.Copy()
	s2 := statedb.Copy()

	var ug1, ug2 uint64
	ug1 = *usedGas
	ug2 = *usedGas

	gp1 := new(GasPool)
	gp2 := new(GasPool)
	*gp1 = *gp
	*gp2 = *gp

	if ug1 != *usedGas {
		panic("invalid ug1")
	} else if ug2 != *usedGas {
		panic("invalid ug2")
	}

	r1, g1, e1 := ApplySputnikTransaction(config, bc, author, gp1, s1, header, tx, &ug1, cfg)
	r2, g2, e2 := applyTransaction(config, bc, author, gp2, s2, header, tx, &ug2, cfg)

	var err error
	var stateDiffs []string
	if g1 != g2 {
		err = evmMismatchErr("gasUsed - svm: %v, gvm: %v", g1, g2)
	} else if !reflect.DeepEqual(r1, r2) {
		err = evmMismatchErr("receipts - svm: %v, gvm: %v", r1, r2)
	} else if (e1 == nil && e2 != nil) || (e1 != nil && e2 == nil) {
		err = evmMismatchErr("err yes/no - svm: %v, gvm: %v", e1, e2)
	} else if (e1 != nil && e2 != nil) && (e1.Error() != e2.Error()) {
		err = evmMismatchErr("err diff - svm: %v, gvm: %v", e1.Error(), e2.Error())
	} else if s1.IntermediateRoot(config.IsEIP161F(header.Number)) != s2.IntermediateRoot(config.IsEIP161F(header.Number)) {
		err = evmMismatchErr("stateroot - svm: %v, gvm: %v", s1.IntermediateRoot(config.IsEIP161F(header.Number)), s2.IntermediateRoot(config.IsEIP161F(header.Number)))
	} else if !reflect.DeepEqual(s1, s2) {
		err = evmMismatchErr("state - svm: %v, gvm: %v", "state1", "state2")
		stateDiffs = deep.Equal(s1, s2)
	}

	if err == nil {
		return nil
	}

	stateDiffsS := strings.Join(stateDiffs, "\n")

	// Finalise the changes on both states concurrently
	done := make(chan struct{})
	go func() {
		s1.Finalise(config.IsEIP161F(header.Number))
		close(done)
	}()
	s2.Finalise(config.IsEIP161F(header.Number))
	<-done

	errResult := map[string]string{
		"g1":         fmt.Sprintf("%v", g1),
		"g2":         fmt.Sprintf("%v", g2),
		"e1":         fmt.Sprintf("%v", e1),
		"e2":         fmt.Sprintf("%v", e2),
		"error":      err.Error(),
		"statediffs": stateDiffsS,
	}

	outDir := filepath.Join(os.Getenv("HOME"),
		"mgsvm-debug",
		fmt.Sprintf("%v", header.Number),
		tx.Hash().Hex(),
	)
	err = os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		return err
	}

	eb, err := json.MarshalIndent(errResult, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(outDir, "err.json"), eb, os.ModePerm)
	if err != nil {
		return err
	}

	// These are much big, very long take.
	// err = ioutil.WriteFile(filepath.Join(outDir, "state-pre.json"), statedb.Dump(), os.ModePerm)
	// if err != nil {
	// 	return err
	// }

	// err = ioutil.WriteFile(filepath.Join(outDir, "state-post-svm.json"), s1.Dump(), os.ModePerm)
	// if err != nil {
	// 	return err
	// }

	// err = ioutil.WriteFile(filepath.Join(outDir, "state-post-gvm.json"), s2.Dump(), os.ModePerm)
	// if err != nil {
	// 	return err
	// }

	headerB, err := header.MarshalJSON()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(outDir, "header.json"), headerB, os.ModePerm)
	if err != nil {
		return err
	}

	txB, err := tx.MarshalJSON()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(outDir, "tx.json"), txB, os.ModePerm)
	if err != nil {
		return err
	}

	for i, l := range s1.Logs() {
		lb, err := l.MarshalJSON()
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filepath.Join(outDir, fmt.Sprintf("logs-post-svm-%d.json", i)), lb, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for i, l := range s2.Logs() {
		lb, err := l.MarshalJSON()
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filepath.Join(outDir, fmt.Sprintf("logs-post-gvm-%d.json", i)), lb, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if r1 != nil {
		r1b, err := r1.MarshalJSON()
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filepath.Join(outDir, "receipt-svm.json"), r1b, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if r2 != nil {
		r2b, err := r2.MarshalJSON()
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filepath.Join(outDir, "receipt-gvm.json"), r2b, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return err
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns the receipt
// for the transaction, gas used and an error if the transaction failed,
// indicating the block was invalid.
func ApplyTransaction(config *params.ChainConfig, bc ChainContext, author *common.Address, gp *GasPool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *uint64, cfg vm.Config) (*types.Receipt, uint64, error) {
	err := compareEVMs(config, bc, author, gp, statedb, header, tx, usedGas, cfg)
	if err != nil {
		return nil, 0, err
	}
	if cfg.EVMInterpreter == "svm" {
		return ApplySputnikTransaction(config, bc, author, gp, statedb, header, tx, usedGas, cfg)
	}
	return applyTransaction(config, bc, author, gp, statedb, header, tx, usedGas, cfg)
}

// applyTransaction is the standard transaction application function, using the built in go evm.
func applyTransaction(config *params.ChainConfig, bc ChainContext, author *common.Address, gp *GasPool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *uint64, cfg vm.Config) (*types.Receipt, uint64, error) {
	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return nil, 0, err
	}
	// Create a new context to be used in the EVM environment
	context := NewEVMContext(msg, header, bc, author)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(context, statedb, config, cfg)
	// Apply the transaction to the current state (included in the env)
	_, gas, failed, err := ApplyMessage(vmenv, msg, gp)
	if err != nil {
		return nil, 0, err
	}
	// Update the state with pending changes
	var root []byte
	if config.IsEIP658F(header.Number) {
		statedb.Finalise(config.IsEIP161F(header.Number))
	} else {
		root = statedb.IntermediateRoot(config.IsEIP161F(header.Number)).Bytes()
	}
	*usedGas += gas

	// Create a new receipt for the transaction, storing the intermediate root and gas used by the tx
	// based on the eip phase, we're passing whether the root touch-delete accounts.
	receipt := types.NewReceipt(root, failed, *usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = gas
	// if the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(vmenv.Context.Origin, tx.Nonce())
	}
	// Set the receipt logs and create a bloom for filtering
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
	receipt.BlockHash = statedb.BlockHash()
	receipt.BlockNumber = header.Number
	receipt.TransactionIndex = uint(statedb.TxIndex())

	return receipt, gas, err
}
