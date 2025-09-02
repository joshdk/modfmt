// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

package modfmt

import (
	"golang.org/x/mod/modfile"
)

// sectionModule formats the `module …` section for `go.mod` files. Returns an
// empty string if the section directive has no value.
//
// See https://go.dev/ref/mod#go-mod-file-module
func sectionModule(directive *modfile.Module) string {
	if directive == nil {
		return ""
	}

	i := item{
		comments: extractComments(directive.Syntax.Before, directive.Syntax.Suffix),
		line:     directive.Mod.Path,
	}

	return value("module", i)
}
