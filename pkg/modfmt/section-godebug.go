// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

package modfmt

import (
	"fmt"

	"golang.org/x/mod/modfile"
)

// sectionGodebug formats the `exclude (…)` section for `go.mod` and `go.work`
// files. Returns an empty string if the section contains no directives.
//
// See https://go.dev/ref/mod#go-mod-file-godebug
// See https://go.dev/ref/mod#go-work-file-godebug
func sectionGodebug(directives []*modfile.Godebug) string {
	items := make([]item, 0, len(directives))

	for _, directive := range directives {
		i := item{
			comments: extractComments(directive.Syntax.Before, directive.Syntax.Suffix),
			line:     fmt.Sprintf("%s=%s", directive.Key, directive.Value),
		}

		items = append(items, i)
	}

	return block("godebug", items)
}
