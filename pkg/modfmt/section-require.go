// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

package modfmt

import (
	"fmt"

	"golang.org/x/mod/modfile"
)

// sectionRequire formats the `require (…)` section for `go.mod` files. Only
// includes require directives where a package is required directly. Returns an
// empty string if the section contains no directives.
//
// See https://go.dev/ref/mod#go-mod-file-require
func sectionRequire(directives []*modfile.Require) string {
	items := make([]item, 0, len(directives))

	for _, directive := range directives {
		if directive.Indirect {
			continue
		}

		i := item{
			comments: extractComments(directive.Syntax.Before, directive.Syntax.Suffix),
			line:     fmt.Sprintf("%s %s", directive.Mod.Path, directive.Mod.Version),
		}

		items = append(items, i)
	}

	return block("require", items)
}

// sectionRequireIndirect formats the `require (…)` section for `go.mod` files.
// Only includes require directives where a package is indirectly directly.
// Returns an empty string if the section contains no directives.
//
// See https://go.dev/ref/mod#go-mod-file-require
func sectionRequireIndirect(directives []*modfile.Require) string {
	items := make([]item, 0, len(directives))

	for _, directive := range directives {
		if !directive.Indirect {
			continue
		}

		i := item{
			comments: extractComments(directive.Syntax.Before),
			line:     fmt.Sprintf("%s %s // indirect", directive.Mod.Path, directive.Mod.Version),
		}

		items = append(items, i)
	}

	return block("require", items)
}
