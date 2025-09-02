// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

package modfmt

import (
	"fmt"
	"strings"

	"golang.org/x/mod/modfile"
)

func isLocal(name string) bool {
	return strings.HasPrefix(name, "./") || strings.HasPrefix(name, "../")
}

// sectionReplace formats the `replace (…)` section for `go.mod` and `go.work`
// files. Only includes replace directives where a package is being replaced by
// another package. Returns an empty string if the section contains no
// directives.
//
// See https://go.dev/ref/mod#go-mod-file-replace
// See https://go.dev/ref/mod#go-work-file-replace
func sectionReplace(directives []*modfile.Replace) string {
	items := make([]item, 0, len(directives))

	for _, directive := range directives {
		if isLocal(directive.New.Path) {
			continue
		}

		i := item{
			comments: extractComments(directive.Syntax.Before, directive.Syntax.Suffix),
			line:     stringReplace(directive),
		}

		items = append(items, i)
	}

	return block("replace", items)
}

func stringReplace(directive *modfile.Replace) string {
	switch {
	case directive.Old.Version != "":
		return fmt.Sprintf("%s %s => %s %s", directive.Old.Path, directive.Old.Version, directive.New.Path, directive.New.Version) //nolint:lll
	default:
		return fmt.Sprintf("%s => %s %s", directive.Old.Path, directive.New.Path, directive.New.Version)
	}
}

// sectionReplaceLocal formats the `replace (…)` section for `go.mod` and
// `go.work` files. Only includes replace directives where a package is being
// replaced by a local file path. Returns an empty string if the section
// contains no directives.
//
// See https://go.dev/ref/mod#go-mod-file-replace
// See https://go.dev/ref/mod#go-work-file-replace
func sectionReplaceLocal(directives []*modfile.Replace) string {
	items := make([]item, 0, len(directives))

	for _, directive := range directives {
		if !isLocal(directive.New.Path) {
			continue
		}

		i := item{
			comments: extractComments(directive.Syntax.Before, directive.Syntax.Suffix),
			line:     stringReplaceLocal(directive),
		}

		items = append(items, i)
	}

	return block("replace", items)
}

func stringReplaceLocal(directive *modfile.Replace) string {
	switch {
	case directive.Old.Version != "":
		return fmt.Sprintf("%s %s => %s", directive.Old.Path, directive.Old.Version, directive.New.Path)
	default:
		return fmt.Sprintf("%s => %s", directive.Old.Path, directive.New.Path)
	}
}
