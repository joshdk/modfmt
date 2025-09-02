// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

package modfmt

import (
	"golang.org/x/mod/modfile"
)

// sectionTool formats the `tool (…)` section for `go.mod` files. Returns an
// empty string if the section contains no directives.
//
// See https://go.dev/ref/mod#go-mod-file-tool
func sectionTool(directives []*modfile.Tool) string {
	items := make([]item, 0, len(directives))

	for _, directive := range directives {
		i := item{
			comments: extractComments(directive.Syntax.Before, directive.Syntax.Suffix),
			line:     directive.Path,
		}

		items = append(items, i)
	}

	return block("tool", items)
}
