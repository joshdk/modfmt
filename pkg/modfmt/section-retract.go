// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

package modfmt

import (
	"fmt"

	"golang.org/x/mod/modfile"
)

// sectionRetract formats the `retract (…)` section for `go.mod` files. Returns
// an empty string if the section contains no directives.
//
// See https://go.dev/ref/mod#go-mod-file-retract
func sectionRetract(directives []*modfile.Retract) string {
	items := make([]item, 0, len(directives))

	for _, directive := range directives {
		i := item{
			comments: extractComments(directive.Syntax.Before, directive.Syntax.Suffix),
			line:     stringRetract(directive),
		}

		items = append(items, i)
	}

	return block("retract", items)
}

func stringRetract(directive *modfile.Retract) string {
	switch {
	case directive.Low != directive.High:
		return fmt.Sprintf("[%s, %s]", directive.Low, directive.High)
	default:
		return directive.Low
	}
}
