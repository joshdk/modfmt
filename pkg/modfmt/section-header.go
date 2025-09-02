// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

package modfmt

import (
	"golang.org/x/mod/modfile"
)

// sectionHeader formats a header comments section for `go.mod` and `go.work`
// files. Returns an empty string if the section directive has no value.
//
// See https://go.dev/ref/mod#go-mod-file-go
// See https://go.dev/ref/mod#go-work-file-go
func sectionHeader(file *modfile.FileSyntax) string {
	var lines []string

	for _, statement := range file.Stmt {
		if commentBlock, ok := statement.(*modfile.CommentBlock); ok {
			lines = append(lines, extractComments(commentBlock.Before)...)
		}
	}

	return comments(lines, "")
}
