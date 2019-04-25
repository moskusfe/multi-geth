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

// +build none

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func encodeOpenRPCSchema(data []byte) string {

	// HAVE: \"description\": \"If `true` it returns
	// WANT: \"description\": \"If ` + "`" + `'true` it returns

	btickS := "`"
	dataS := strconv.QuoteToASCII(string(data))
	bescS := fmt.Sprintf(`%s + "%s" + %s`, btickS, btickS, btickS)
	dataS = strings.Replace(dataS, btickS, bescS, -1)
	return dataS
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: mkopenrpc openrpc.json")
		os.Exit(1)
	}

	bs, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	// schema := make(map[string]interface{})
	// file, err := os.Open(os.Args[1])
	// if err != nil {
	// 	panic(err)
	// }

	// err = json.NewDecoder(file).Decode(schema)
	// if err != nil {
	// 	panic(err)
	// }

	// data, err := rlp.EncodeToBytes(schema)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Printf(`// Copyright 2019 The go-ethereum Authors
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

package rpc

// File contains auto-generated constant(s) containing schema data for OpenRPC.
// Their content is a JSON string.
// Use mkopenrpc.go to create/update them.

// nolint: misspell

const openRPCSchema = %s%s%s

`, "`", encodeOpenRPCSchema(bs), "`")
}
