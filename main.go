// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

// Package main provides the entry point for main.
package main

import (
	"fmt"
	"os"

	"github.com/joshdk/modfmt/cmd"
)

func main() {
	if err := cmd.Command().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "modfmt:", err)
		os.Exit(1)
	}
}
