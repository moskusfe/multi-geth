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
  It outputs a const declaration that contains a string.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ansii = false

// backtickEscape escape backticks in a string literal
func backtickEscape(s string) string {

	// HAVE: \"description\": \"If `true` it returns
	// WANT: \"description\": \"If ` + "`" + `'true` it returns

	btickS := "`"
	bescS := fmt.Sprintf(`%s + "%s" + %s`, btickS, btickS, btickS)
	s = strings.Replace(s, btickS, bescS, -1)
	return s
}

// multilineStringValue should return a string value
// GOTCHA, maybe: string lines MAYBE MUST be shortened to not be 20k+ chars long...
func multilineStringValue(data []string, ansiiQuoted bool) string {
	if len(data) == 0 {
		panic("no lines")
	}

	// NOTE: the string MUST LITERALLY include the backticks; remember that we are writing the value of a string declaration, not a string
	var ls string
	for i, d := range data {
		if ansiiQuoted {
			d = strconv.QuoteToASCII(d)
		}
		if i < len(data)-1 {
			ls += fmt.Sprintf("`%s`+\n", backtickEscape(d))
		} else {
			ls += fmt.Sprintf("`%s`\n", backtickEscape(d))
		}
	}
	return ls
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: mkopenrpc openrpc.json")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
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

package rpc

// File contains auto-generated constant(s) containing schema data for OpenRPC.
// Their content is a JSON string.
// Use mkopenrpc.go to create/update them.

const openRPCSchema = %s
`, multilineStringValue(lines, ansii))
}
