// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

package modfmt

import (
	"golang.org/x/mod/modfile"
)

// sectionGo formats the `go …` section for `go.mod` and `go.work` files.
// Returns an empty string if the section directive has no value.
//
// https://go.dev/ref/mod#go-mod-file-go
// https://go.dev/ref/mod#go-work-file-go
func sectionGo(directive *modfile.Go) string {
	if directive == nil {
		return ""
	}

	i := item{
		comments: extractComments(directive.Syntax.Before, directive.Syntax.Suffix),
		line:     directive.Version,
	}

	return value("go", i)
}
