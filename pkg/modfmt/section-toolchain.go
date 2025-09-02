// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

package modfmt

import (
	"golang.org/x/mod/modfile"
)

// sectionToolchain formats the `toolchain …` section for `go.mod` and
// `go.work` files. Returns an empty string if the section directive has no
// value.
//
// See https://go.dev/ref/mod#go-mod-file-toolchain
// See https://go.dev/ref/mod#go-work-file-toolchain
func sectionToolchain(directive *modfile.Toolchain) string {
	if directive == nil {
		return ""
	}

	i := item{
		comments: extractComments(directive.Syntax.Before, directive.Syntax.Suffix),
		line:     directive.Name,
	}

	return value("toolchain", i)
}
