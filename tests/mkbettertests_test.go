// This is only a test. Reads test files, and prints them. OK.

package tests

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
- Tests should run respective to chain configurations.
- Chain configurations should be integrated as test data.
- Chain configurations which are intergrated as test data should be spec'd,
  so that any client can use those configurations in their feature switch implementations
  and thus actually use the tests to prove expectations.

Operations:
- Define a chain configuration schema (JSON). This must be expected to play nicely with multi-geth and parity-ethereum.
  + AFAIR there might already exist a tool which can translate between configuration schemas;
    go-ethereum considers chain configuration as a facet of genesis block specification.
  + multi-geth already has JSON tags defined for params.ChainConfig
- Implement "test" suite that can run incorporated chain configurations defined under this schema
  against _existing_ test chain configurations (aka Forks). If tests pass, we know that our test runner
  and chain configurations are playing well together to match existing chain configurations.
- Implement "test runner" that is actually a "test generator," which, instead of check got vs. want test results,
  will write test results to some file in JSON format.
- Using this test generator, create JSON tests against chain configurations beyond the currently hardcoded ETH-only
  magical snake charming tests. These should _be able_ to prove that existing Parity and Go-ethereum (multi-geth)
  implementations pass tests equivalently for existing (live network) chain configurations.
- Implement a chain configuration generator that can create chain X test permutations, where chains include but
  are not limited to any network's canonical specifications. This will aid to ferret out edge cases.

*/

func TestMakeBetterTests(t *testing.T) {
	// t.Parallel()

	st := new(testMatcher)
	// Long tests:
	st.slow(`^stAttackTest/ContractCreationSpam`)
	st.slow(`^stBadOpcode/badOpcodes`)
	st.slow(`^stPreCompiledContracts/modexp`)
	st.slow(`^stQuadraticComplexityTest/`)
	st.slow(`^stStaticCall/static_Call50000`)
	st.slow(`^stStaticCall/static_Return50000`)
	st.slow(`^stStaticCall/static_Call1MB`)
	st.slow(`^stSystemOperationsTest/CallRecursiveBomb`)
	st.slow(`^stTransactionTest/Opcodes_TransactionInit`)
	// Broken tests:
	st.skipLoad(`^stTransactionTest/OverflowGasRequire\.json`) // gasLimit > 256 bits
	st.skipLoad(`^stTransactionTest/zeroSigTransa[^/]*\.json`) // EIP-86 is not supported yet
	// Expected failures:
	st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/Byzantium/0`, "bug in test")
	st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/Byzantium/3`, "bug in test")
	st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/Constantinople/0`, "bug in test")
	st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/Constantinople/3`, "bug in test")
	st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/ConstantinopleFix/0`, "bug in test")
	st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/ConstantinopleFix/3`, "bug in test")

	st.walk(t, stateTestDir, func(t *testing.T, name string, test *StateTest) {
		for _, subtest := range test.Subtests() {
			// subtest := subtest // closure
			// key := fmt.Sprintf("%s/%d", subtest.Fork, subtest.Index)
			// name := name + "/" + key

			x, _ := json.MarshalIndent(test.json, "", "    ")
			fmt.Printf("%s / %s", x, subtest.Fork)

		}
	})
}
