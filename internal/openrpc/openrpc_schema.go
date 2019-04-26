// Copyright 2019 The go-ethereum Authors
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

package openrpc

// File contains auto-generated constant(s) containing schema data for OpenRPC.
// Their content is a JSON string.
// Use mkopenrpc.go to create/update them.

// OpenRPCSchema defines the default full suite of possibly available go-ethereum RPC
// methods.
const OpenRPCSchema = `
{
  "openrpc": "1.0.0",
  "info": {
    "version": "1.0.0",
    "title": "Multi Geth RPC",
    "description": "This API lets you interact with the blockchain via JSON-RPC",
    "license": {
      "name": "Apache 2.0",
      "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
    }
  },
  "servers": [
    {
      "url": "http://localhost:8545"
    }
  ],
  "methods": [
    {
      "name": "eth_getBlockByHash",
      "summary": "Gets a block for a given hash",
      "params": [
        {
          "name": "blockHash",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/BlockHash"
          }
        },
        {
          "name": "includeTransactions",
          "description": "If ` + "`" + `true` + "`" + ` it returns the full transaction objects, if ` + "`" + `false` + "`" + ` only the hashes of the transactions.",
          "required": true,
          "schema": {
            "type": "boolean"
          }
        }
      ],
      "result": {
        "name": "getBlockByHashResult",
        "schema": {
          "oneOf": [
            {
              "$ref": "#/components/schemas/Block"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      }
    },
    {
      "name": "eth_getBlockByNumber",
      "summary": "Gets a block for a given number",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/BlockNumber"
        },
        {
          "name": "includeTransactions",
          "description": "If ` + "`" + `true` + "`" + ` it returns the full transaction objects, if ` + "`" + `false` + "`" + ` only the hashes of the transactions.",
          "required": true,
          "schema": {
            "type": "boolean"
          }
        }
      ],
      "result": {
        "name": "getBlockByNumberResult",
        "schema": {
          "oneOf": [
            {
              "$ref": "#/components/schemas/Block"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      }
    },
    {
      "name": "eth_blockNumber",
      "summary": "Returns the number of most recent block.",
      "params": [],
      "result": {
        "$ref": "#/components/contentDescriptors/BlockNumber"
      }
    },
    {
      "name": "eth_getStorageAt",
      "summary": "Gets a storage value from a contract address, a position, and an optional blockNumber",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Address"
        },
        {
          "$ref": "#/components/contentDescriptors/Position"
        },
        {
          "$ref": "#/components/contentDescriptors/BlockNumber"
        }
      ],
      "result": {
        "name": "dataWord",
        "schema": {
          "$ref": "#/components/schemas/DataWord"
        }
      }
    },
    {
      "name": "eth_getTransactionCount",
      "summary": "Returns the number of transactions sent from an address",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Address"
        },
        {
          "$ref": "#/components/contentDescriptors/BlockNumber"
        }
      ],
      "result": {
        "name": "transactionCount",
        "schema": {
          "oneOf": [
            {
              "$ref": "#/components/schemas/Nonce"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      }
    },
    {
      "name": "eth_getTransactionByHash",
      "summary": "Returns the information about a transaction requested by transaction hash.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/TransactionHash"
        }
      ],
      "result": {
        "$ref": "#/components/contentDescriptors/TransactionResult"
      }
    },
    {
      "name": "eth_getTransactionByBlockHashAndIndex",
      "summary": "Returns the information about a transaction requested by the block hash and index of which it was mined.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/BlockHash"
        },
        {
          "name": "index",
          "description": "The ordering in which a transaction is mined within its block.",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/Integer"
          }
        }
      ],
      "result": {
        "$ref": "#/components/contentDescriptors/TransactionResult"
      },
      "examples": [
        {
          "name": "nullExample",
          "params": [
            {
              "name": "blockHashExample",
              "value": "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"
            },
            {
              "name": "indexExample",
              "value": "0x0"
            }
          ],
          "result": {
            "name": "nullResultExample",
            "value": null
          }
        }
      ]
    },
    {
      "name": "eth_getTransactionByBlockNumberAndIndex",
      "summary": "Returns the information about a transaction requested by the block hash and index of which it was mined.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/BlockNumber"
        },
        {
          "name": "index",
          "description": "The ordering in which a transaction is mined within its block.",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/Integer"
          }
        }
      ],
      "result": {
        "$ref": "#/components/contentDescriptors/TransactionResult"
      }
    },
    {
      "name": "eth_getTransactionReceipt",
      "summary": "Returns the receipt information of a transaction by its hash.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/TransactionHash"
        }
      ],
      "result": {
        "name": "transactionReceiptResult",
        "description": "returns either a receipt or null",
        "schema": {
          "oneOf": [
            {
              "$ref": "#/components/schemas/Receipt"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      }
    },
    {
      "name": "eth_getUncleByBlockHashAndIndex",
      "summary": "Returns information about a uncle of a block by hash and uncle index position.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/BlockHash"
        },
        {
          "name": "index",
          "description": "The ordering in which a uncle is included within its block.",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/Integer"
          }
        }
      ],
      "result": {
        "name": "uncle",
        "schema": {
          "oneOf": [
            {
              "$ref": "#/components/schemas/Uncle"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      }
    },
    {
      "name": "eth_getUncleByBlockNumberAndIndex",
      "summary": "Returns information about a uncle of a block by hash and uncle index position.",
      "params": [
        {
          "name": "uncleBlockNumber",
          "description": "The block in which the uncle was included",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/BlockNumber"
          }
        },
        {
          "name": "index",
          "description": "The ordering in which a uncle is included within its block.",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/Integer"
          }
        }
      ],
      "result": {
        "name": "uncleResult",
        "description": "returns an uncle or null",
        "schema": {
          "oneOf": [
            {
              "$ref": "#/components/schemas/Uncle"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      },
      "examples": [
        {
          "name": "nullResultExample",
          "params": [
            {
              "name": "uncleBlockNumberExample",
              "value": "0x0"
            },
            {
              "name": "uncleBlockNumberIndexExample",
              "value": "0x0"
            }
          ],
          "result": {
            "name": "nullResultExample",
            "value": null
          }
        }
      ]
    },
    {
      "name": "eth_newFilter",
      "summary": "Creates a filter object, based on filter options, to notify when the state changes (logs). To check if the state has changed, call eth_getFilterChanges.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Filter"
        }
      ],
      "result": {
        "name": "filterId",
        "schema": {
          "description": "The filter ID for use in ` + "`" + `eth_getFilterChanges` + "`" + `",
          "$ref": "#/components/schemas/Integer"
        }
      }
    },
    {
      "name": "eth_newBlockFilter",
      "summary": "Creates a filter in the node, to notify when a new block arrives. To check if the state has changed, call eth_getFilterChanges.",
      "params": [],
      "result": {
        "$ref": "#/components/contentDescriptors/FilterId"
      }
    },
    {
      "name": "eth_newPendingTransactionFilter",
      "summary": "Creates a filter in the node, to notify when new pending transactions arrive. To check if the state has changed, call eth_getFilterChanges.",
      "params": [],
      "result": {
        "$ref": "#/components/contentDescriptors/FilterId"
      }
    },
    {
      "name": "eth_uninstallFilter",
      "summary": "Uninstalls a filter with given id. Should always be called when watch is no longer needed. Additonally Filters timeout when they aren't requested with eth_getFilterChanges for a period of time.",
      "params": [
        {
          "name": "filterId",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/FilterId"
          }
        }
      ],
      "result": {
        "name": "filterUninstalledSuccess",
        "schema": {
          "type": "boolean",
          "description": "Whether of not the filter was successfully uninstalled"
        }
      }
    },
    {
      "name": "eth_getFilterChanges",
      "summary": "Polling method for a filter, which returns an array of logs which occurred since last poll.",
      "params": [
        {
          "name": "filterId",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/FilterId"
          }
        }
      ],
      "result": {
        "name": "logResult",
        "schema": {
          "type": "array",
          "items": {
            "$ref": "#/components/schemas/Log"
          }
        }
      }
    },
    {
      "name": "eth_getFilterLogs",
      "summary": "Returns an array of all logs matching filter with given id.",
      "params": [
        {
          "name": "filterId",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/FilterId"
          }
        }
      ],
      "result": {
        "$ref": "#/components/contentDescriptors/Logs"
      }
    },
    {
      "name": "eth_getLogs",
      "summary": "Returns an array of all logs matching a given filter object.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Filter"
        }
      ],
      "result": {
        "$ref": "#/components/contentDescriptors/Logs"
      }
    },
    {
      "name": "eth_getWork",
      "summary": "Returns the hash of the current block, the seedHash, and the boundary condition to be met ('target').",
      "params": [],
      "result": {
        "name": "work",
        "schema": {
          "type": "array",
          "items": [
            {
              "$ref": "#/components/schemas/PowHash"
            },
            {
              "$ref": "#/components/schemas/SeedHash"
            },
            {
              "$ref": "#/components/schemas/Difficulty"
            }
          ]
        }
      }
    },
    {
      "name": "eth_submitWork",
      "summary": "Used for submitting a proof-of-work solution.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Nonce"
        },
        {
          "name": "powHash",
          "schema": {
            "$ref": "#/components/schemas/PowHash"
          }
        },
        {
          "name": "mixHash",
          "schema": {
            "$ref": "#/components/schemas/MixHash"
          }
        }
      ],
      "result": {
        "name": "solutionValid",
        "description": "returns true if the provided solution is valid, otherwise false.",
        "schema": {
          "type": "boolean",
          "description": "Whether or not the provided solution is valid"
        }
      },
      "examples": [
        {
          "name": "submitWorkExample",
          "params": [
            {
              "name": "nonceExample",
              "description": "example of a number only used once",
              "value": "0x0000000000000001"
            },
            {
              "name": "powHashExample",
              "description": "proof of work to submit",
              "value": "0x6bf2cAE0dE3ec3ecA5E194a6C6e02cf42aADfe1C2c4Fff12E5D36C3Cf7297F22"
            },
            {
              "name": "mixHashExample",
              "description": "the mix digest example",
              "value": "0xD1FE5700000000000000000000000000D1FE5700000000000000000000000000"
            }
          ],
          "result": {
            "name": "solutionInvalidExample",
            "description": "this example should return ` + "`" + `false` + "`" + ` as it is not a valid pow to submit",
            "value": false
          }
        }
      ]
    },
    {
      "name": "eth_submitHashrate",
      "summary": "Returns an array of all logs matching a given filter object.",
      "params": [
        {
          "name": "hashRate",
          "schema": {
            "$ref": "#/components/schemas/DataWord"
          }
        },
        {
          "name": "id",
          "description": "String identifiying the client",
          "schema": {
            "$ref": "#/components/schemas/DataWord"
          }
        }
      ],
      "result": {
        "name": "submitHashRateSuccess",
        "schema": {
          "type": "boolean",
          "description": "whether of not submitting went through succesfully"
        }
      }
    },
    {
      "name": "eth_getProof",
      "summary": "Returns the account- and storage-values of the specified account including the Merkle-proof.",
      "params": [
        {
          "name": "address",
          "description": "The address of the account or contract",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/Address"
          }
        },
        {
          "name": "storageKeys",
          "schema": {
            "description": "The storage keys of all the storage slots being requested",
            "items": {
              "description": "A storage key is indexed from the solidity compiler by the order it is declaired. For mappings it uses the keccak of the mapping key with its position (and recursively for X-dimentional mappings)",
              "$ref": "#/components/schemas/Integer"
            }
          }
        },
        {
          "$ref": "#/components/contentDescriptors/BlockNumber"
        }
      ],
      "result": {
        "name": "account",
        "schema": {
          "oneOf": [
            {
              "type": "object",
              "description": "The merkle proofs of the specified account connecting them to the blockhash of the block specified",
              "properties": {
                "address": {
                  "description": "The address of the account or contract of the request",
                  "$ref": "#/components/schemas/Address"
                },
                "accountProof": {
                  "$ref": "#/components/schemas/AccountProof"
                },
                "balance": {
                  "description": "The Ether balance of the account or contract of the request",
                  "$ref": "#/components/schemas/Integer"
                },
                "codeHash": {
                  "description": "The code hash of the contract of the request (keccak(NULL) if external account)",
                  "$ref": "#/components/schemas/Keccak"
                },
                "nonce": {
                  "description": "The transaction count of the account or contract of the request",
                  "$ref": "#/components/schemas/Nonce"
                },
                "storageHash": {
                  "description": "The storage hash of the contract of the request (keccak(rlp(NULL)) if external account)",
                  "$ref": "#/components/schemas/Keccak"
                },
                "storageProof": {
                  "$ref": "#/components/schemas/StorageProof"
                }
              }
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      }
    },
    {
      "name": "eth_getCode",
      "summary": "Returns code at a given contract address",
      "params": [
        {
          "name": "address",
          "description": "The address of the contract",
          "schema": {
            "$ref": "#/components/schemas/Address"
          }
        },
        {
          "name": "blockNumber",
          "description": "A BlockNumber of which the code existed",
          "schema": {
            "$ref": "#/components/schemas/BlockNumber"
          }
        }
      ],
      "result": {
        "name": "bytes",
        "schema": {
          "$ref": "#/components/schemas/Bytes"
        }
      }
    },
    {
      "name": "eth_getBalance",
      "summary": "Returns Ether balance of a given or account or contract",
      "params": [
        {
          "name": "address",
          "description": "The address of the acccount or contract",
          "schema": {
            "$ref": "#/components/schemas/Address"
          }
        },
        {
          "name": "blockNumber",
          "description": "A BlockNumber at which to request the balance",
          "schema": {
            "$ref": "#/components/schemas/BlockNumber"
          }
        }
      ],
      "result": {
        "name": "getBalanceResult",
        "schema": {
          "oneOf": [
            {
              "$ref": "#/components/schemas/Integer"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      }
    },
    {
      "name": "eth_sign",
      "summary": "The sign method calculates an Ethereum specific signature with: sign(keccak256( '\\x19Ethereum Signed Message:\\n' + len(message) + message))).",
      "params": [
        {
          "name": "address",
          "description": "The address of the account whos signature to use.",
          "schema": {
            "$ref": "#/components/schemas/Address"
          }
        },
        {
          "name": "message",
          "description": "The message to sign",
          "schema": {
            "$ref": "#/components/schemas/Bytes"
          }
        }
      ],
      "result": {
        "name": "signatureData",
        "schema": {
          "description": "The signature data.",
          "$ref": "#/components/schemas/Bytes"
        }
      }
    },
    {
      "name": "eth_accounts",
      "summary": "Returns a list of accounts owned by the client",
      "params": [],
      "result": {
        "name": "addresses",
        "schema": {
          "description": "An array of addresses owned by the client",
          "type": "array",
          "items": {
            "$ref": "#/components/schemas/Address"
          }
        }
      }
    },
    {
      "name": "eth_gasPrice",
      "summary": "Returns the current price per gas in wei",
      "params": [],
      "result": {
        "$ref": "#/components/contentDescriptors/GasPrice"
      }
    },
    {
      "name": "eth_hashrate",
      "summary": "Returns the number of hashes per second that the node is mining with.",
      "params": [],
      "result": {
        "name": "hashesPerSecond",
        "schema": {
          "description": "Integer of the number of hashes per second",
          "$ref": "#/components/schemas/Integer"
        }
      }
    },
    {
      "name": "eth_mining",
      "summary": "Returns true if client is actively mining new blocks.",
      "params": [],
      "result": {
        "name": "mining",
        "schema": {
          "description": "Whether of not the client is mining",
          "type": "boolean"
        }
      }
    },
    {
      "name": "eth_coinbase",
      "summary": "Returns the client coinbase address.",
      "params": [],
      "result": {
        "name": "address",
        "schema": {
          "description": "The address owned by the client that is used as default for things like the mining reward",
          "$ref": "#/components/schemas/Address"
        }
      }
    },
    {
      "name": "eth_protocolVersion",
      "summary": "Returns the current ethereum protocol version.",
      "params": [],
      "result": {
        "name": "protocolVersion",
        "schema": {
          "description": "The current ethereum protocol version",
          "$ref": "#/components/schemas/Integer"
        }
      }
    },
    {
      "name": "eth_getBlockTransactionCountByHash",
      "summary": "Returns the number of transactions in a block from a block matching the given block hash.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/BlockHash"
        }
      ],
      "result": {
        "name": "blockTransactionCountByHash",
        "description": "The Number of total transactions in the given block",
        "schema": {
          "oneOf": [
            {
              "$ref": "#/components/schemas/Integer"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      }
    },
    {
      "name": "eth_getBlockTransactionCountByNumber",
      "summary": "Returns the number of transactions in a block from a block matching the given block number.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/BlockNumber"
        }
      ],
      "result": {
        "name": "blockTransactionCountByHash",
        "description": "The Number of total transactions in the given block",
        "schema": {
          "oneOf": [
            {
              "$ref": "#/components/schemas/Integer"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      }
    },
    {
      "name": "eth_getUncleCountByBlockHash",
      "summary": "Returns the number of uncles in a block from a block matching the given block hash.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/BlockHash"
        }
      ],
      "result": {
        "name": "uncleCountResult",
        "description": "The Number of total uncles in the given block",
        "schema": {
          "oneOf": [
            {
              "$ref": "#/components/schemas/Integer"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      }
    },
    {
      "name": "eth_getUncleCountByBlockNumber",
      "summary": "Returns the number of uncles in a block from a block matching the given block number.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/BlockNumber"
        }
      ],
      "result": {
        "name": "uncleCountResult",
        "schema": {
          "oneOf": [
            {
              "description": "The Number of total uncles in the given block",
              "$ref": "#/components/schemas/Integer"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      }
    },
    {
      "name": "eth_sendTransaction",
      "summary": "Creates new message call transaction or a contract creation, if the data field contains code.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Transaction"
        }
      ],
      "result": {
        "name": "transactionHash",
        "schema": {
          "description": "The transaction hash, or the zero hash if the transaction is not yet available.",
          "$ref": "#/components/schemas/Keccak"
        }
      }
    },
    {
      "name": "eth_sendRawTransaction",
      "summary": "Creates new message call transaction or a contract creation for signed transactions.",
      "params": [
        {
          "name": "signedTransactionData",
          "description": "The signed transaction data",
          "schema": {
            "$ref": "#/components/schemas/Bytes"
          }
        }
      ],
      "result": {
        "name": "transactionHash",
        "schema": {
          "description": "The transaction hash, or the zero hash if the transaction is not yet available.",
          "$ref": "#/components/schemas/Keccak"
        }
      }
    },
    {
      "name": "eth_call",
      "summary": "Executes a new message call (locally) immediately without creating a transaction on the block chain.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Transaction"
        },
        {
          "$ref": "#/components/contentDescriptors/BlockNumber"
        }
      ],
      "result": {
        "name": "returnValue",
        "schema": {
          "description": "The return value of the executed contract",
          "$ref": "#/components/schemas/Bytes"
        }
      }
    },
    {
      "name": "eth_estimateGas",
      "summary": "Generates and returns an estimate of how much gas is necessary to allow the transaction to complete. The transaction will not be added to the blockchain. Note that the estimate may be significantly more than the amount of gas actually used by the transaction, for a variety of reasons including EVM mechanics and node performance.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Transaction"
        }
      ],
      "result": {
        "name": "gasUsed",
        "schema": {
          "description": "The amount of gas used",
          "$ref": "#/components/schemas/Integer"
        }
      }
    },
    {
      "name": "eth_syncing",
      "summary": "Returns an object with data about the sync status or false.",
      "params": [],
      "result": {
        "name": "syncing",
        "schema": {
          "oneOf": [
            {
              "description": "An object with sync status data",
              "type": "object",
              "properties": {
                "startingBlock": {
                  "description": "Block at which the import started (will only be reset, after the sync reached his head)",
                  "$ref": "#/components/schemas/Integer"
                },
                "currentBlock": {
                  "description": "The current block, same as eth_blockNumber",
                  "$ref": "#/components/schemas/Integer"
                },
                "highestBlock": {
                  "description": "The estimated highest block",
                  "$ref": "#/components/schemas/Integer"
                },
                "knownStates": {
                  "description": "The known states",
                  "$ref": "#/components/schemas/Integer"
                },
                "pulledStates": {
                  "description": "The pulled states",
                  "$ref": "#/components/schemas/Integer"
                }
              }
            },
            {
              "type": "boolean",
              "description": "The value ` + "`" + `false` + "`" + ` indicating that syncing is complete"
            }
          ]
        }
      }
    },
    {
      "name": "admin_exportChain",
      "description": "Exports a blockchain into the specified file, appending to the file if data already exists in it",
      "params": [
        {
          "name": "fp",
          "description": "path to file",
          "schema": {
            "type": "string"
          }
        }
      ],
      "result": {
        "name": "result",
        "description": "result of chain exporting",
        "schema": {
          "description": "` + "`" + `true` + "`" + ` if exported or ` + "`" + `false` + "`" + ` failed",
          "type": "boolean"
        }
      }
    },
    {
      "name": "admin_importChain",
      "description": "Import full chain",
      "params": [
        {
          "name": "chain",
          "description": "path to file containing chain",
          "schema": {
            "type": "string"
          }
        }
      ],
      "result": {
        "name": "result",
        "description": "result of chain import",
        "schema": {
          "description": "` + "`" + `true` + "`" + ` if imported or ` + "`" + `false` + "`" + ` failed",
          "type": "boolean"
        }
      }
    },
    {
      "name": "debug_chaindbCompact",
      "description": "Compact chain DB. Does not work for memory databases",
      "params": [],
      "result": {
        "name": "null",
        "schema": {
          "type": "null"
        }
      }
    },
    {
      "name": "debug_chaindbProperty",
      "description": "returns leveldb properties of the chain database.",
      "params": [
        {
          "name": "property",
          "schema": {
            "type": "string"
          }
        }
      ],
      "result": {
        "name": "leveldb property",
        "schema": {
          "type": "string"
        }
      }
    },
    {
      "name": "debug_dumpBlock",
      "description": "Retrieves the state that corresponds to the block number and returns a list of accounts (including storage and code)",
      "params": [
        {
          "name": "blockNumber",
          "schema": {
            "type": "string",
            "pattern": "^0x[a-fA-F\\d]+$"
          }
        }
      ],
      "result": {
        "name": "Block state",
        "schema": {
          "description": "",
          "$ref": "#/components/schemas/Block"
        }
      }
    },
    {
      "name": "debug_getBadBlocks",
      "description": "returns a list of the last 'bad blocks' that the client has seen on the network and returns them as a JSON list of block-hashes",
      "params": [],
      "result": {
        "name": "badBlocks",
        "schema": {
          "type": "array",
          "items": {
            "$ref": "#/components/schemas/BlockHash"
          }
        }
      }
    },
    {
      "name": "debug_getBlockRlp",
      "description": "Retrieves and returns the RLP encoded block by number",
      "params": [
        {
          "name": "blockNumber",
          "description": "number of target block",
          "schema": {
            "type": "integer",
            "minimum": 0
          }
        }
      ],
      "result": {
        "name": "blockRLP",
        "schema": {
          "$ref": "#/components/schemas/BlockRLP"
        }
      }
    },
    {
      "name": "debug_getModifiedAccountsByHash",
      "description": "returns all accounts that have changed between the two blocks specified.\nA change is defined as a difference in nonce, balance, code hash, or storage hash.\n With one parameter, returns the list of accounts modified in the specified block.",
      "params": [
        {
          "name": "startHash",
          "description": "hash of start block",
          "schema": {
            "$ref": "#/components/schemas/BlockHash"
          }
        },
        {
          "name": "endHash",
          "description": "hash of end block",
          "schema": {
            "oneOf": [
              {
                "$ref": "#/components/schemas/BlockHash"
              },
              {
                "$ref": "#/components/schemas/Null"
              }
            ]
          }
        }
      ],
      "result": {
        "name": "changedAccounts",
        "schema": {
          "description": "accounts that have changed",
          "type": "array",
          "items": {
            "$ref": "#/components/schemas/Address"
          }
        }
      }
    },
    {
      "name": "debug_getModifiedAccountsByNumber",
      "description": "returns all accounts that have changed between the two blocks specified.\n A change is defined as a difference in nonce, balance, code hash, or storage hash. \nWith one parameter, returns the list of accounts modified in the specified block.",
      "params": [
        {
          "name": "startHash",
          "description": "hash of start block",
          "schema": {
            "type": "integer",
            "minimum": 0
          }
        },
        {
          "name": "endHash",
          "description": "hash of end block",
          "schema": {
            "oneOf": [
              {
                "type": "integer",
                "minimum": 0
              },
              {
                "$ref": "#/components/schemas/Null"
              }
            ]
          }
        }
      ],
      "result": {
        "name": "changedAccounts",
        "schema": {
          "description": "accounts that have changed",
          "type": "array",
          "items": {
            "$ref": "#/components/schemas/Address"
          }
        }
      }
    },
    {
      "name": "debug_preimage",
      "description": "preimage retrieves a cached trie node pre-image from memory. \nIf it cannot be found cached, the method queries the persistent database for the content",
      "params": [
        {
          "name": "nodeHash",
          "description": "hash for required tree node",
          "schema": {
            "$ref": "#/components/schemas/Keccak"
          }
        }
      ],
      "result": {
        "name": "preimage",
        "schema": {
          "description": "requested pre-image",
          "$ref": "#/components/schemas/Bytes"
        }
      }
    },
    {
      "name": "debug_printBlock",
      "description": "PrintBlock retrieves a block and returns its pretty printed form",
      "params": [
        {
          "name": "blockNumber",
          "description": "number of target block",
          "schema": {
            "type": "integer",
            "minimum": 0
          }
        }
      ],
      "result": {
        "name": "result",
        "description": "pretty printed block",
        "schema": {
          "type": "string"
        }
      }
    },
    {
      "name": "debug_seedHash",
      "description": "Fetches and retrieves the seed hash of the block by number",
      "params": [
        {
          "name": "blockNumber",
          "description": "number of target block",
          "schema": {
            "type": "integer",
            "minimum": 0
          }
        }
      ],
      "result": {
        "name": "seed hash of the block",
        "schema": {
          "$ref": "#/components/schemas/Bytes"
        }
      }
    },
    {
      "name": "debug_setHead",
      "description": "Sets the current head of the local chain by block number. \nNote: this is a destructive action and may severely damage your chain. Use with extreme caution.",
      "params": [
        {
          "name": "blockNumber",
          "description": "number of target block",
          "schema": {
            "$ref": "#/components/schemas/BlockNumber"
          }
        }
      ],
      "result": {
        "name": "head is set to the new block",
        "description": "` + "`" + `true` + "`" + ` if head set or ` + "`" + `null` + "`" + ` if failed",
        "schema": {
          "oneOf": [
            {
              "type": "boolean"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      }
    },
    {
      "name": "debug_standardTraceBadBlockToFile",
      "description": "StandardTraceBadBlockToFile dumps the structured logs created during the execution of EVM against a block pulled from the pool of bad ones to the local file system and returns a list of files to the caller",
      "params": [
        {
          "name": "badBlockHash",
          "description": "hash for a bad block",
          "schema": {
            "$ref": "#/components/schemas/BlockHash"
          }
        },
        {
          "name": "StdTraceConfig",
          "description": "holds extra paramters to standard-json trace functions",
          "schema": {
            "type": "object",
            "properties": {
              "logConfig": {
                "$ref": "#/components/schemas/LogConfig"
              },
              "reexec": {
                "type": "integer"
              },
              "txHash": {
                "$ref": "#/components/schemas/TransactionHash"
              }
            }
          }
        }
      ],
      "result": {
        "name": "result",
        "description": "bad block traced to file",
        "schema": {
          "description": "` + "`" + `true` + "`" + ` if traced to file or ` + "`" + `false` + "`" + ` failed",
          "type": "boolean"
        }
      }
    },
    {
      "name": "debug_standardTraceBlockToFile",
      "description": "standardTraceBlockToFile configures a new tracer which uses standard JSON output,\n and traces either a full block or an individual transaction.\n The return value will be one filename per transaction traced.",
      "params": [
        {
          "name": "blockHash",
          "description": "hash for a block",
          "schema": {
            "$ref": "#/components/schemas/BlockHash"
          }
        }
      ],
      "result": {
        "name": "block traced to file",
        "schema": {
          "description": "` + "`" + `true` + "`" + ` if traced to file or ` + "`" + `false` + "`" + ` failed",
          "type": "boolean"
        }
      }
    },
    {
      "name": "debug_storageRangeAt",
      "description": "StorageRangeAt returns the storage at the given block height and transaction index.",
      "params": [
        {
          "name": "blockHash",
          "description": "hash for a block",
          "schema": {
            "$ref": "#/components/schemas/BlockHash"
          }
        },
        {
          "name": "txIndex",
          "description": "transaction index",
          "schema": {
            "type": "integer"
          }
        },
        {
          "$ref": "#/components/contentDescriptors/Address"
        },
        {
          "name": "startKey",
          "description": "Start Key",
          "schema": {
            "type": "string",
            "pattern": "^0x[a-fA-F\\d]{40}$"
          }
        },
        {
          "name": "limit",
          "summary": "number of storage entries",
          "description": "Number of storage entries to return.",
          "schema": {
            "type": "integer"
          }
        }
      ],
      "result": {
        "name": "storage range",
        "schema": {
          "$ref": "#/components/schemas/StorageRange"
        }
      }
    },
    {
      "name": "debug_traceBadBlock",
      "description": "TraceBadBlockByHash returns the structured logs created during the execution of EVM against a block pulled from the pool of bad ones and returns them as a JSON object.",
      "params": [
        {
          "name": "badBlockHash",
          "schema": {
            "$ref": "#/components/schemas/BlockHash"
          }
        }
      ],
      "result": {
        "name": "badBlockResult",
        "schema": {
          "description": "",
          "$ref": "#/components/schemas/Bytes"
        }
      }
    },
    {
      "name": "debug_traceBlock",
      "description": "Will return a full stack trace of all invoked opcodes of all transaction that were included included in this block. \nNote: the parent of this block must be present or it will fail",
      "params": [
        {
          "name": "blockRLP",
          "description": "RLP representation of the block",
          "schema": {
            "$ref": "#/components/schemas/BlockRLP"
          }
        }
      ],
      "result": {
        "name": "stackTrace",
        "schema": {
          "description": "",
          "$ref": "#/components/schemas/BlockRLP"
        }
      }
    },
    {
      "name": "debug_traceBlockByHash",
      "description": "Similar to ` + "`" + `debug_traceBlock` + "`" + `, accepts a block hash and will replay the block that is already present in the database.",
      "params": [
        {
          "name": "blockHash",
          "description": "target block hash",
          "schema": {
            "$ref": "#/components/schemas/BlockHash"
          }
        }
      ],
      "result": {
        "name": "trace",
        "schema": {
          "$ref": "#/components/schemas/Bytes"
        }
      }
    },
    {
      "name": "debug_traceBlockByNumber",
      "description": "Similar to debug_traceBlock, accepts a block number and will replay the block that is already present in the database.",
      "params": [
        {
          "name": "blockNumber",
          "description": "target block number",
          "schema": {
            "$ref": "#/components/schemas/BlockNumber"
          }
        }
      ],
      "result": {
        "name": "trace",
        "schema": {
          "$ref": "#/components/schemas/Bytes"
        }
      }
    },
    {
      "name": "debug_traceBlockFromFile",
      "description": "Similar to ` + "`" + `debug_traceBlock` + "`" + `, accepts a file containing the RLP of the block",
      "params": [
        {
          "name": "filename",
          "description": "file containing with block's RLP",
          "schema": {
            "type": "string"
          }
        }
      ],
      "result": {
        "name": "traceBlockResult",
        "schema": {
          "description": "",
          "$ref": "#/components/schemas/Bytes"
        }
      }
    },
    {
      "name": "debug_traceTransaction",
      "description": "The traceTransaction debugging method will attempt to run the transaction in the exact same manner as it was executed on the network.\n It will replay any transaction that may have been executed prior to this one before it will finally attempt to execute the transaction that corresponds to the given hash.",
      "params": [
        {
          "name": "debugTransactionHash",
          "description": "Transaction hash to debug",
          "schema": {
            "$ref": "#/components/schemas/TransactionHash"
          }
        }
      ],
      "result": {
        "name": "trace transaction object",
        "schema": {
          "type": "object"
        }
      }
    },
    {
      "name": "miner_getHashrate",
      "summary": "Returns the number of hashes per second that the node is mining with.",
      "params": [],
      "result": {
        "name": "hashesPerSecond",
        "schema": {
          "type": "integer",
          "description": "Integer of the number of hashes per second"
        }
      }
    },
    {
      "name": "miner_setEtherbase",
      "description": "In order to earn ether you must have your etherbase (or coinbase) address set. This etherbase defaults to your primary account. If you don't have an etherbase address, then geth --mine will not start up.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Address"
        }
      ],
      "result": {
        "name": "setEtherBaseSuccess",
        "description": "returns true if success",
        "schema": {
          "type": "boolean"
        }
      }
    },
    {
      "name": "miner_setExtra",
      "description": "There is an option to add extra Data (32 bytes only) to your mined blocks. By convention this is interpreted as a unicode string, so you can set your short vanity tag.",
      "params": [
        {
          "name": "extraData",
          "description": ". By convention this is interpreted as a unicode string, so you can set your short vanity tag.",
          "schema": {
            "type": "string",
            "maxLength": 64
          }
        }
      ],
      "result": {
        "name": "setExtraSuccess",
        "description": "returns true if success",
        "schema": {
          "type": "boolean"
        }
      }
    },
    {
      "name": "miner_setGasPrice",
      "description": "Sets the minimum accepted gas price for the miner.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/GasPrice"
        }
      ],
      "result": {
        "name": "setGasPriceSuccess",
        "description": "returns true if success",
        "schema": {
          "type": "boolean"
        }
      }
    },
    {
      "name": "miner_setRecommitInterval",
      "description": "updates the interval for miner sealing work recommitting.",
      "params": [
        {
          "name": "interval",
          "description": "interval for miner sealing work recommitting.",
          "schema": {
            "type": "integer"
          }
        }
      ],
      "result": {
        "name": "null",
        "schema": {
          "$ref": "#/components/schemas/Null"
        }
      }
    },
    {
      "name": "miner_start",
      "description": "Start starts the miner with the given number of threads. If threads is not defined, the number of workers started is equal to the number of logical CPUs that are usable by this process. If mining is already running, this method adjust the number of threads allowed to use and updates the minimum price required by the transaction pool.",
      "params": [
        {
          "name": "threads",
          "description": "number of given threads",
          "schema": {
            "type": "integer"
          }
        }
      ],
      "result": {
        "name": "null",
        "schema": {
          "$ref": "#/components/schemas/Null"
        }
      }
    },
    {
      "name": "miner_stop",
      "description": "Stop terminates the miner, both at the consensus engine level as well as at the block creation level.",
      "params": [],
      "result": {
        "name": "null",
        "schema": {
          "$ref": "#/components/schemas/Null"
        }
      }
    },
    {
      "name": "net_listening",
      "description": "Determines if this client is listening for new network connections.",
      "params": [],
      "result": {
        "name": "netListeningResult",
        "description": "` + "`" + `true` + "`" + ` if listening is active or ` + "`" + `false` + "`" + ` if listening is not active",
        "schema": {
          "type": "boolean"
        }
      },
      "examples": [
        {
          "name": "netListeningTrueExample",
          "description": "example of true result for net_listening",
          "params": [],
          "result": {
            "name": "netListeningExampleFalseResult",
            "value": true
          }
        }
      ]
    },
    {
      "name": "net_peerCount",
      "description": "Returns the number of peers currently connected to this client.",
      "params": [],
      "result": {
        "name": "quantity",
        "description": "number of connected peers.",
        "schema": {
          "description": "Hex representation of number of connected peers",
          "type": "string"
        }
      }
    },
    {
      "name": "net_version",
      "description": "Returns the chain ID associated with the current network.",
      "params": [],
      "result": {
        "name": "chainID",
        "description": "chain ID associated with the current network",
        "schema": {
          "type": "string"
        }
      }
    },
    {
      "name": "personal_deriveAccount",
      "description": "requests a HD wallet to derive a new account, optionally pinning it for later reuse.",
      "params": [
        {
          "name": "url",
          "description": "base url",
          "required": true,
          "schema": {
            "type": "string",
            "format": "uri"
          }
        },
        {
          "name": "path",
          "required": true,
          "description": "path to derive a new account",
          "schema": {
            "type": "string"
          }
        },
        {
          "name": "pin",
          "description": "pin for later use",
          "schema": {
            "type": "boolean"
          }
        }
      ],
      "result": {
        "$ref": "#/components/contentDescriptors/Address"
      }
    },
    {
      "name": "personal_ecRecover",
      "description": "returns the address associated with the private key that was used to calculate the signature in personal_sign.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Message"
        },
        {
          "$ref": "#/components/contentDescriptors/Signature"
        }
      ],
      "result": {
        "$ref": "#/components/contentDescriptors/Address"
      }
    },
    {
      "name": "personal_importRawKey",
      "description": "Imports the given unencrypted private key (hex string) into the key store, encrypting it with the passphrase.",
      "params": [
        {
          "name": "keydata",
          "description": "unencrypted private key",
          "schema": {
            "type": "string",
            "pattern": "[a-fA-F\\d]{64}$"
          }
        },
        {
          "$ref": "#/components/contentDescriptors/Passphrase"
        }
      ],
      "result": {
        "name": "Address of imported private key",
        "schema": {
          "$ref": "#/components/schemas/Address"
        }
      }
    },
    {
      "name": "personal_listAccounts",
      "description": "Returns all the Ethereum account addresses of all keys in the key store.",
      "params": [],
      "result": {
        "name": "addresses",
        "schema": {
          "description": "ethereum account addresses in the key store.",
          "type": "array",
          "items": {
            "$ref": "#/components/schemas/Address"
          }
        }
      }
    },
    {
      "name": "personal_listWallets",
      "description": "ListWallets will return a list of wallets this node manages.",
      "params": [],
      "result": {
        "name": "rawWallet",
        "schema": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "url": {
                "type": "string"
              },
              "status": {
                "type": "string"
              },
              "accounts": {
                "description": "An array of addresses owned by the client",
                "type": "array",
                "items": {
                  "type": "object",
                  "properties": {
                    "address": {
                      "$ref": "#/components/schemas/Address"
                    },
                    "url": {
                      "type": "string",
                      "format": "uri"
                    }
                  }
                }
              },
              "failure": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    {
      "name": "personal_lockAccount",
      "description": "lock the account associated with the given address when it's unlocked.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Address"
        }
      ],
      "result": {
        "name": "lockAccountSuccess",
        "schema": {
          "description": "returns true if successfully locked account",
          "type": "boolean"
        }
      }
    },
    {
      "name": "personal_newAccount",
      "description": "create a new account and returns the address for the new account.",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Passphrase"
        }
      ],
      "result": {
        "$ref": "#/components/contentDescriptors/Address"
      }
    },
    {
      "name": "personal_openWallet",
      "description": "OpenWallet initiates a hardware wallet opening procedure, establishing a USB connection and attempting to authenticate via the provided passphrase. Note, the method may return an extra challenge requiring a second open (e.g. the Trezor PIN matrix challenge).",
      "params": [
        {
          "name": "url",
          "description": "hd path for hardware wallet",
          "schema": {
            "type": "string",
            "format": "uri"
          }
        },
        {
          "$ref": "#/components/contentDescriptors/Passphrase"
        }
      ],
      "result": {
        "name": "null",
        "schema": {
          "$ref": "#/components/schemas/Null"
        }
      }
    },
    {
      "name": "personal_sign",
      "description": "Signs data using a specific account. This data is before UTF-8 HEX decoded and enveloped as follows: ` + "`" + `'\\x19Ethereum Signed Message:\n' + message.length + message` + "`" + `.",
      "params": [
        {
          "name": "message",
          "required": true,
          "description": "The message to sign",
          "schema": {
            "$ref": "#/components/schemas/Bytes"
          }
        },
        {
          "name": "address",
          "required": true,
          "description": "The address of the account who's signature to use.",
          "schema": {
            "$ref": "#/components/schemas/Address"
          }
        },
        {
          "$ref": "#/components/contentDescriptors/Passphrase"
        }
      ],
      "result": {
        "name": "signatureData",
        "schema": {
          "description": "The signature data.",
          "$ref": "#/components/schemas/Bytes"
        }
      }
    },
    {
      "name": "personal_sendTransaction",
      "description": "Validate the given passphrase and submit transaction. ",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Transaction"
        },
        {
          "$ref": "#/components/contentDescriptors/Passphrase"
        }
      ],
      "result": {
        "$ref": "#/components/contentDescriptors/TransactionHash"
      }
    },
    {
      "name": "personal_signAndSendTransaction",
      "description": "Validate the given passphrase and submit transaction. ",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Transaction"
        },
        {
          "$ref": "#/components/contentDescriptors/Passphrase"
        }
      ],
      "result": {
        "$ref": "#/components/contentDescriptors/TransactionHash"
      }
    },
    {
      "name": "personal_signTransaction",
      "description": "SignTransaction will create a transaction from the given arguments and tries to sign it with the key associated with args.To. If the given passwd isn't able to decrypt the key it fails. The transaction is returned in RLP-form, not broadcast to other nodes",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Transaction"
        },
        {
          "$ref": "#/components/contentDescriptors/Passphrase"
        }
      ],
      "result": {
        "name": "signedTransaction",
        "description": "transaction returned in RLP-form. not broadcast to other nodes.",
        "schema": {
          "type": "string"
        }
      }
    },
    {
      "name": "personal_unlockAccount",
      "description": "UnlockAccount will unlock the account associated with the given address with the given password for duration seconds. If duration is nil it will use a default of 300 seconds. It returns an indication if the account was unlocked. ",
      "params": [
        {
          "$ref": "#/components/contentDescriptors/Address"
        },
        {
          "$ref": "#/components/contentDescriptors/Passphrase"
        },
        {
          "name": "duration",
          "schema": {
            "type": "integer",
            "minimum": 0
          }
        }
      ],
      "result": {
        "name": "unlockAccountSuccess",
        "description": "returns true if unlock success",
        "schema": {
          "type": "boolean"
        }
      }
    },
    {
      "name": "txpool_inspect",
      "description": "The inspect inspection property can be queried to list a textual summary of all the transactions currently pending for inclusion in the next block(s),\n as well as the ones that are being scheduled for future execution only.",
      "params": [],
      "result": {
        "name": "inspection",
        "schema": {
          "$ref": "#/components/schemas/TxPoolInspection"
        }
      }
    },
    {
      "name": "txpool_status",
      "description": "The status inspection property can be queried for the number of transactions currently pending for inclusion in the next block(s),\n as well as the ones that are being scheduled for future execution only.",
      "params": [],
      "result": {
        "name": "status",
        "schema": {
          "$ref": "#/components/schemas/TxPoolStatus"
        }
      }
    }
  ],
  "components": {
    "schemas": {
      "LogConfig": {
        "type": "object",
        "description": "Logconfig are the configuration options for the structured logger for the EVM.",
        "properties": {
          "disableMemory": {
            "type": "boolean"
          },
          "disableStack": {
            "type": "boolean"
          },
          "disableStorage": {
            "type": "boolean"
          },
          "debug": {
            "description": "print debug during capture end",
            "type": "boolean"
          },
          "limit": {
            "description": "maximum length of output, but zero means unlimited",
            "type": "integer"
          }
        }
      },
      "CodeResponse": {
        "description": "An object containing information about the code.",
        "type": "object",
        "properties": {
          "code": {
            "description": "The compiled Byte code",
            "$ref": "#/components/schemas/Bytes"
          },
          "info": {
            "description": "An object contianing information about the code compilation.",
            "type": "object",
            "properties": {
              "source": {
                "type": "string",
                "description": "The sorce code that was compiled"
              },
              "language": {
                "type": "string",
                "description": "The language of the code that was compiled"
              },
              "languageVersion": {
                "type": "string",
                "description": "The language version number"
              },
              "compilerVersion": {
                "type": "string",
                "description": "The sorce code that was compiled"
              },
              "abiDefinition": {
                "type": "object",
                "description": "The application binary interface definitions of the code"
              }
            }
          }
        }
      },
      "ProofNode": {
        "type": "string",
        "description": "An indiviual node used to prove a path down a merkle-patricia-tree",
        "$ref": "#/components/schemas/Bytes"
      },
      "AccountProof": {
        "$ref": "#/components/schemas/ProofNodes"
      },
      "StorageProof": {
        "type": "array",
        "description": "Current block header PoW hash.",
        "items": {
          "type": "object",
          "description": "Object proving a relationship of a storage value to an account's storageHash.",
          "properties": {
            "key": {
              "description": "The key used to get the storage slot in its account tree",
              "$ref": "#/components/schemas/Integer"
            },
            "value": {
              "description": "The value of the storage slot in its account tree",
              "$ref": "#/components/schemas/Integer"
            },
            "proof": {
              "$ref": "#/components/schemas/ProofNodes"
            }
          }
        }
      },
      "ProofNodes": {
        "type": "array",
        "description": "The set of node values needed to traverse a patricia merkle tree (from root to leaf) to retrieve a value",
        "items": {
          "$ref": "#/components/schemas/ProofNode"
        }
      },
      "PowHash": {
        "description": "Current block header PoW hash.",
        "$ref": "#/components/schemas/DataWord"
      },
      "SeedHash": {
        "description": "The seed hash used for the DAG.",
        "$ref": "#/components/schemas/DataWord"
      },
      "TxPoolInspection": {
        "description": "Transaction pool inspection is an object with two fields pending and queued.\n Each of these fields are associative arrays,\n in which each entry maps an origin-address to a batch of scheduled transactions.\n These batches themselves are maps associating nonces with transactions summary strings.",
        "schema": {
          "type": "object",
          "properties": {
            "pending": {
              "description": "Pending transactions list",
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/components/schemas/TxBatch"
                }
              }
            },
            "queued": {
              "description": "Queued transactions list",
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/components/schemas/TxBatch"
                }
              }
            }
          }
        }
      },
      "TxBatch": {
        "description": "Maps an origin-address to a batch of scheduled transactions",
        "schema": {
          "type": "object",
          "properties": {
            "address": {
              "description": "Account address to which batch of transactions relates",
              "schema": {
                "$ref": "#/components/schemas/Address"
              }
            },
            "batch": {
              "description": "Batch of transactions for current address",
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/components/schemas/TxBatchEntry"
                }
              }
            }
          }
        }
      },
      "TxBatchEntry": {
        "description": "Map associating nonces with transactions summary strings.",
        "schema": {
          "type": "object",
          "properties": {
            "nonce": {
              "description": "nonce",
              "schema": {
                "type": "integer"
              }
            },
            "tx": {
              "description": "Transaction summary strings",
              "schema": {
                "type": "string"
              }
            }
          }
        }
      },
      "TxPoolStatus": {
        "description": " Transaction pool status. An object with two fields ` + "`" + `pending` + "`" + ` and ` + "`" + `queued` + "`" + `, \neach of which is a counter representing the number of transactions in that particular state",
        "schema": {
          "type": "object",
          "properties": {
            "pending": {
              "description": "Pending transactions count",
              "schema": {
                "type": "integer"
              }
            },
            "queued": {
              "description": "Queued transactions count",
              "schema": {
                "type": "integer"
              }
            }
          }
        }
      },
      "MixHash": {
        "description": "The mix digest.",
        "$ref": "#/components/schemas/DataWord"
      },
      "Difficulty": {
        "description": "The boundary condition ('target'), 2^256 / difficulty.",
        "$ref": "#/components/schemas/DataWord"
      },
      "FilterId": {
        "type": "string",
        "description": "An identifier used to reference the filter."
      },
      "BlockHash": {
        "type": "string",
        "pattern": "^0x[a-fA-F\\d]{64}$",
        "description": "The hex representation of the Keccak 256 of the RLP encoded block"
      },
      "BlockNumber": {
        "type": "string",
        "pattern": "^0x[a-fA-F\\d]+$",
        "description": "The hex representation of the block's height"
      },
      "BlockRLP": {
        "type": "string",
        "pattern": "^[a-fA-F\\d]+$",
        "description": "The hex representation of the block's RLP"
      },
      "BlockNumberTag": {
        "type": "string",
        "description": "The optional block height description",
        "enum": [
          "earliest",
          "latest",
          "pending"
        ]
      },
      "StorageRange": {
        "type": "object",
        "description": "allows downloading storage in chunks. debug_storageRangeAt takes parameters blockHash, txIndex, account, startKey, limit. The startKey applies to the hashed key. The returned storage values contain the preimage if available. The returned object also contains the nextKey field, which will be non-null if there are more keys after the range that was returned. ",
        "properties": {
          "storage": {
            "patternProperties": {
              "^0x[a-fA-F\\d]{40}$": {
                "type": "object",
                "properties": {
                  "key": {
                    "type": "string",
                    "pattern": "^0x[a-fA-F\\d]+$"
                  },
                  "value": {
                    "type": "string",
                    "pattern": "^0x[a-fA-F\\d]+$"
                  }
                }
              }
            }
          },
          "nextKey": {
            "oneOf": [
              {
                "type": "string",
                "pattern": "^0x[a-fA-F\\d]+$"
              },
              {
                "$ref": "#/components/schemas/Null"
              }
            ]
          }
        }
      },
      "Receipt": {
        "type": "object",
        "description": "The receipt of a transaction",
        "required": [
          "blockHash",
          "blockNumber",
          "contractAddress",
          "cumulativeGasUsed",
          "from",
          "gasUsed",
          "logs",
          "logsBloom",
          "to",
          "transactionHash",
          "transactionIndex"
        ],
        "properties": {
          "blockHash": {
            "description": "BlockHash of the block in which the transaction was mined",
            "$ref": "#/components/schemas/BlockHash"
          },
          "blockNumber": {
            "description": "BlockNumber of the block in which the transaction was mined",
            "$ref": "#/components/schemas/BlockNumber"
          },
          "contractAddress": {
            "description": "The contract address created, if the transaction was a contract creation, otherwise null",
            "$ref": "#/components/schemas/Address"
          },
          "cumulativeGasUsed": {
            "description": "The gas units used by the transaction",
            "$ref": "#/components/schemas/Integer"
          },
          "from": {
            "description": "The sender of the transaction",
            "$ref": "#/components/schemas/Address"
          },
          "gasUsed": {
            "description": "The total gas used by the transaction",
            "$ref": "#/components/schemas/Integer"
          },
          "logs": {
            "type": "array",
            "description": "An array of all the logs triggered during the transaction",
            "items": {
              "$ref": "#/components/schemas/Log"
            }
          },
          "logsBloom": {
            "$ref": "#/components/schemas/BloomFilter"
          },
          "to": {
            "description": "Destination address of the transaction",
            "$ref": "#/components/schemas/Address"
          },
          "transactionHash": {
            "description": "Keccak 256 of the transaction",
            "$ref": "#/components/schemas/Keccak"
          },
          "transactionIndex": {
            "description": "An array of all the logs triggered during the transaction",
            "$ref": "#/components/schemas/BloomFilter"
          },
          "postTransactionState": {
            "description": "The intermediate stateRoot directly after transaction execution.",
            "$ref": "#/components/schemas/Keccak"
          },
          "status": {
            "description": "Whether or not the transaction threw an error.",
            "type": "boolean"
          }
        }
      },
      "BloomFilter": {
        "type": "string",
        "description": "A 2048 bit bloom filter from the logs of the transaction. Each log sets 3 bits though taking the low-order 11 bits of each of the first three pairs of bytes in a Keccak 256 hash of the log's byte series"
      },
      "Log": {
        "type": "object",
        "description": "An indexed event generated during a transaction",
        "properties": {
          "address": {
            "description": "Sender of the transaction",
            "$ref": "#/components/schemas/Address"
          },
          "blockHash": {
            "description": "BlockHash of the block in which the transaction was mined",
            "$ref": "#/components/schemas/BlockHash"
          },
          "blockNumber": {
            "description": "BlockNumber of the block in which the transaction was mined",
            "$ref": "#/components/schemas/BlockNumber"
          },
          "data": {
            "description": "The data/input string sent along with the transaction",
            "$ref": "#/components/schemas/Bytes"
          },
          "logIndex": {
            "description": "The index of the event within its transaction, null when its pending",
            "$ref": "#/components/schemas/Integer"
          },
          "removed": {
            "schema": {
              "description": "Whether or not the log was orphaned off the main chain",
              "type": "boolean"
            }
          },
          "topics": {
            "type": "array",
            "items": {
              "topic": {
                "description": "32 Bytes DATA of indexed log arguments. (In solidity: The first topic is the hash of the signature of the event (e.g. Deposit(address,bytes32,uint256))",
                "$ref": "#/components/schemas/DataWord"
              }
            }
          },
          "transactionHash": {
            "description": "The hash of the transaction in which the log occurred",
            "$ref": "#/components/schemas/Keccak"
          },
          "transactionIndex": {
            "description": "The index of the transaction in which the log occurred",
            "$ref": "#/components/schemas/Integer"
          }
        }
      },
      "Uncle": {
        "type": "object",
        "description": "Orphaned blocks that can be included in the chain but at a lower block reward. NOTE: An uncle doesn’t contain individual transactions.",
        "properties": {
          "number": {
            "description": "The block number or null when its the pending block",
            "$ref": "#/components/schemas/IntOrPending"
          },
          "hash": {
            "description": "The block hash or null when its the pending block",
            "$ref": "#/components/schemas/KeccakOrPending"
          },
          "parentHash": {
            "description": "Hash of the parent block",
            "$ref": "#/components/schemas/Keccak"
          },
          "nonce": {
            "description": "Randomly selected number to satisfy the proof-of-work or null when its the pending block",
            "$ref": "#/components/schemas/IntOrPending"
          },
          "sha3Uncles": {
            "description": "Keccak hash of the uncles data in the block",
            "$ref": "#/components/schemas/Keccak"
          },
          "logsBloom": {
            "type": "string",
            "description": "The bloom filter for the logs of the block or null when its the pending block",
            "pattern": "^0x[a-fA-F\\d]+$"
          },
          "transactionsRoot": {
            "description": "The root of the transactions trie of the block.",
            "$ref": "#/components/schemas/Keccak"
          },
          "stateRoot": {
            "description": "The root of the final state trie of the block",
            "$ref": "#/components/schemas/Keccak"
          },
          "receiptsRoot": {
            "description": "The root of the receipts trie of the block",
            "$ref": "#/components/schemas/Keccak"
          },
          "miner": {
            "description": "The address of the beneficiary to whom the mining rewards were given or null when its the pending block",
            "oneOf": [
              {
                "$ref": "#/components/schemas/Address"
              },
              {
                "$ref": "#/components/schemas/Null"
              }
            ]
          },
          "difficulty": {
            "type": "string",
            "description": "Integer of the difficulty for this block"
          },
          "totalDifficulty": {
            "description": "Integer of the total difficulty of the chain until this block",
            "$ref": "#/components/schemas/IntOrPending"
          },
          "extraData": {
            "type": "string",
            "description": "The 'extra data' field of this block"
          },
          "size": {
            "type": "string",
            "description": "Integer the size of this block in bytes"
          },
          "gasLimit": {
            "type": "string",
            "description": "The maximum gas allowed in this block"
          },
          "gasUsed": {
            "type": "string",
            "description": "The total used gas by all transactions in this block"
          },
          "timestamp": {
            "type": "string",
            "description": "The unix timestamp for when the block was collated"
          },
          "uncles": {
            "description": "Array of uncle hashes",
            "type": "array",
            "items": {
              "description": "Block hash of the RLP encoding of an uncle block",
              "$ref": "#/components/schemas/Keccak"
            }
          }
        }
      },
      "Block": {
        "type": "object",
        "properties": {
          "number": {
            "description": "The block number or null when its the pending block",
            "$ref": "#/components/schemas/IntOrPending"
          },
          "hash": {
            "description": "The block hash or null when its the pending block",
            "$ref": "#/components/schemas/KeccakOrPending"
          },
          "parentHash": {
            "description": "Hash of the parent block",
            "$ref": "#/components/schemas/Keccak"
          },
          "nonce": {
            "description": "Randomly selected number to satisfy the proof-of-work or null when its the pending block",
            "$ref": "#/components/schemas/IntOrPending"
          },
          "sha3Uncles": {
            "description": "Keccak hash of the uncles data in the block",
            "$ref": "#/components/schemas/Keccak"
          },
          "logsBloom": {
            "type": "string",
            "description": "The bloom filter for the logs of the block or null when its the pending block",
            "pattern": "^0x[a-fA-F\\d]+$"
          },
          "transactionsRoot": {
            "description": "The root of the transactions trie of the block.",
            "$ref": "#/components/schemas/Keccak"
          },
          "stateRoot": {
            "description": "The root of the final state trie of the block",
            "$ref": "#/components/schemas/Keccak"
          },
          "receiptsRoot": {
            "description": "The root of the receipts trie of the block",
            "$ref": "#/components/schemas/Keccak"
          },
          "miner": {
            "description": "The address of the beneficiary to whom the mining rewards were given or null when its the pending block",
            "oneOf": [
              {
                "$ref": "#/components/schemas/Address"
              },
              {
                "$ref": "#/components/schemas/Null"
              }
            ]
          },
          "difficulty": {
            "type": "string",
            "description": "Integer of the difficulty for this block"
          },
          "totalDifficulty": {
            "description": "Integer of the total difficulty of the chain until this block",
            "$ref": "#/components/schemas/IntOrPending"
          },
          "extraData": {
            "type": "string",
            "description": "The 'extra data' field of this block"
          },
          "size": {
            "type": "string",
            "description": "Integer the size of this block in bytes"
          },
          "gasLimit": {
            "type": "string",
            "description": "The maximum gas allowed in this block"
          },
          "gasUsed": {
            "type": "string",
            "description": "The total used gas by all transactions in this block"
          },
          "timestamp": {
            "type": "string",
            "description": "The unix timestamp for when the block was collated"
          },
          "transactions": {
            "description": "Array of transaction objects, or 32 Bytes transaction hashes depending on the last given parameter",
            "type": "array",
            "items": {
              "oneOf": [
                {
                  "$ref": "#/components/schemas/Transaction"
                },
                {
                  "$ref": "#/components/schemas/TransactionHash"
                }
              ]
            }
          },
          "uncles": {
            "description": "Array of uncle hashes",
            "type": "array",
            "items": {
              "description": "Block hash of the RLP encoding of an uncle block",
              "$ref": "#/components/schemas/Keccak"
            }
          }
        }
      },
      "Transaction": {
        "type": "object",
        "required": [
          "gas",
          "gasPrice",
          "nonce"
        ],
        "properties": {
          "blockHash": {
            "description": "Hash of the block where this transaction was in. null when its pending",
            "$ref": "#/components/schemas/KeccakOrPending"
          },
          "blockNumber": {
            "description": "Block number where this transaction was in. null when its pending",
            "$ref": "#/components/schemas/IntOrPending"
          },
          "from": {
            "description": "Address of the sender",
            "$ref": "#/components/schemas/Address"
          },
          "gas": {
            "type": "string",
            "description": "The gas limit provided by the sender in Wei"
          },
          "gasPrice": {
            "type": "string",
            "description": "The gas price willing to be paid by the sender in Wei"
          },
          "hash": {
            "$ref": "#/components/schemas/TransactionHash"
          },
          "input": {
            "type": "string",
            "description": "The data field sent with the transaction"
          },
          "nonce": {
            "description": "The total number of prior transactions made by the sender",
            "$ref": "#/components/schemas/Nonce"
          },
          "to": {
            "description": "address of the receiver. null when its a contract creation transaction",
            "$ref": "#/components/schemas/Address"
          },
          "transactionIndex": {
            "description": "Integer of the transaction's index position in the block. null when its pending",
            "$ref": "#/components/schemas/IntOrPending"
          },
          "value": {
            "description": "Value of Ether being transferred in Wei",
            "$ref": "#/components/schemas/Keccak"
          },
          "v": {
            "type": "string",
            "description": "ECDSA recovery id"
          },
          "r": {
            "type": "string",
            "description": "ECDSA signature r"
          },
          "s": {
            "type": "string",
            "description": "ECDSA signature s"
          }
        }
      },
      "TransactionHash": {
        "type": "string",
        "description": "Keccak 256 Hash of the RLP encoding of a transaction",
        "$ref": "#/components/schemas/Keccak"
      },
      "KeccakOrPending": {
        "oneOf": [
          {
            "$ref": "#/components/schemas/Keccak"
          },
          {
            "$ref": "#/components/schemas/Null"
          }
        ]
      },
      "IntOrPending": {
        "oneOf": [
          {
            "$ref": "#/components/schemas/Integer"
          },
          {
            "$ref": "#/components/schemas/Null"
          }
        ]
      },
      "Keccak": {
        "type": "string",
        "description": "Hex representation of a Keccak 256 hash",
        "pattern": "^0x[a-fA-F\\d]{64}$"
      },
      "Nonce": {
        "description": "A number only to be used once",
        "pattern": "^0x[a-fA-F0-9]+$",
        "type": "string"
      },
      "Null": {
        "type": "null",
        "description": "Null"
      },
      "Integer": {
        "type": "string",
        "pattern": "^0x[a-fA-F0-9]+$",
        "description": "Hex representation of the integer"
      },
      "Address": {
        "type": "string",
        "pattern": "^0x[a-fA-F\\d]{40}$"
      },
      "Position": {
        "type": "string",
        "description": "Hex representation of the storage slot where the variable exists",
        "pattern": "^0x([a-fA-F0-9]?)+$"
      },
      "DataWord": {
        "type": "string",
        "description": "Hex representation of a 256 bit unit of data",
        "pattern": "^0x([a-fA-F\\d]{64})?$"
      },
      "Bytes": {
        "type": "string",
        "description": "Hex representation of a variable length byte array",
        "pattern": "^0x([a-fA-F0-9]?)+$"
      }
    },
    "contentDescriptors": {
      "Block": {
        "name": "block",
        "summary": "A block",
        "description": "A block object",
        "schema": {
          "$ref": "#/components/schemas/Block"
        }
      },
      "Null": {
        "name": "Null",
        "description": "JSON Null value",
        "summary": "Null value",
        "schema": {
          "type": "null",
          "description": "Null value"
        }
      },
      "Signature": {
        "name": "signature",
        "summary": "The signature.",
        "required": true,
        "schema": {
          "$ref": "#/components/schemas/Bytes",
          "pattern": "0x^([A-Fa-f0-9]{2}){65}$"
        }
      },
      "GasPrice": {
        "name": "gasPrice",
        "required": true,
        "schema": {
          "description": "Integer of the current gas price",
          "$ref": "#/components/schemas/Integer"
        }
      },
      "Passphrase": {
        "name": "passphrase",
        "required": true,
        "description": "passphrase for the new account",
        "schema": {
          "type": "string"
        }
      },
      "Transaction": {
        "required": true,
        "name": "transaction",
        "schema": {
          "$ref": "#/components/schemas/Transaction"
        }
      },
      "TransactionResult": {
        "name": "transactionResult",
        "description": "Returns a transaction or null",
        "schema": {
          "oneOf": [
            {
              "$ref": "#/components/schemas/Transaction"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        }
      },
      "Message": {
        "name": "message",
        "required": true,
        "schema": {
          "$ref": "#/components/schemas/Bytes"
        }
      },
      "Filter": {
        "name": "filter",
        "required": true,
        "schema": {
          "type": "object",
          "description": "A filter used to monitor the blockchain for log/events",
          "properties": {
            "fromBlock": {
              "description": "Block from which to begin filtering events",
              "$ref": "#/components/schemas/BlockNumber"
            },
            "toBlock": {
              "description": "Block from which to end filtering events",
              "$ref": "#/components/schemas/BlockNumber"
            },
            "address": {
              "oneOf": [
                {
                  "type": "string",
                  "description": "Address of the contract from which to monitor events",
                  "$ref": "#/components/schemas/Address"
                },
                {
                  "type": "array",
                  "description": "List of contract addresses from which to monitor events",
                  "items": {
                    "$ref": "#/components/schemas/Address"
                  }
                }
              ]
            },
            "topics": {
              "type": "array",
              "description": "Array of 32 Bytes DATA topics. Topics are order-dependent. Each topic can also be an array of DATA with 'or' options",
              "items": {
                "description": "Indexable 32 bytes piece of data (made from the event's function signature in solidity)",
                "$ref": "#/components/schemas/DataWord"
              }
            }
          }
        }
      },
      "Address": {
        "name": "address",
        "required": true,
        "schema": {
          "$ref": "#/components/schemas/Address"
        }
      },
      "BlockHash": {
        "name": "blockHash",
        "required": true,
        "schema": {
          "$ref": "#/components/schemas/BlockHash"
        }
      },
      "Nonce": {
        "name": "nonce",
        "schema": {
          "$ref": "#/components/schemas/Nonce"
        }
      },
      "Position": {
        "name": "key",
        "required": true,
        "schema": {
          "$ref": "#/components/schemas/Position"
        }
      },
      "Logs": {
        "name": "logs",
        "description": "An array of all logs matching filter with given id.",
        "schema": {
          "type": "array",
          "items": {
            "$ref": "#/components/schemas/Log"
          }
        }
      },
      "FilterId": {
        "name": "filterId",
        "schema": {
          "description": "The filter ID for use in ` + "`" + `eth_getFilterChanges` + "`" + `",
          "$ref": "#/components/schemas/Integer"
        }
      },
      "BlockNumber": {
        "name": "blockNumber",
        "required": true,
        "schema": {
          "oneOf": [
            {
              "$ref": "#/components/schemas/BlockNumber"
            },
            {
              "$ref": "#/components/schemas/BlockNumberTag"
            }
          ]
        }
      },
      "TransactionHash": {
        "name": "transactionHash",
        "required": true,
        "schema": {
          "$ref": "#/components/schemas/TransactionHash"
        }
      }
    }
  }
}
`