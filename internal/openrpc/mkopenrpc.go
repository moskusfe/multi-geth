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

/*

  The mkopenrpc tool creates the OpenRPC JSON schema constant in openrpc_schema.go
  It prints a packaged const declaration that contains a string to stdout.

*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// backtickEscape escape backticks in a string literal
func backtickEscape(s string) string {

	// HAVE: \"description\": \"If `true` it returns
	// WANT: \"description\": \"If ` + "`" + `'true` it returns

	btickS := "`"
	bescS := fmt.Sprintf(`%s + "%s" + %s`, btickS, btickS, btickS)
	s = strings.Replace(s, btickS, bescS, -1)
	return s
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

package openrpc

// File contains auto-generated constant(s) containing schema data for OpenRPC.
// Their content is a JSON string.
// Use mkopenrpc.go to create/update them.

// OpenRPCSchema defines the default full suite of possibly available go-ethereum RPC
// methods.
const OpenRPCSchema = %s
%s
%s
`, "`", backtickEscape(string(bs)), "`")
}
