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

// item represents a single entry to be used either alone in a value directive
// or as a collection in a block directive.
type item struct {
	// comments is an optional set of comments to include with this entry.
	comments []string

	// line is a (potentially formatted) string value for this entry.
	line string
}

// comments formats the given lines as comments with an optional indent prefix.
func comments(lines []string, indent string) string {
	var result string
	for _, line := range lines {
		result += fmt.Sprintf("%s// %s\n", indent, line)
	}

	return result
}

// value formats a single value directive (e.g.`module …`). Returns an empty string
// if the given item has an empty line.
func value(name string, i item) string {
	if i.line == "" {
		return ""
	}

	result := comments(i.comments, "")
	result += fmt.Sprintf("%s %s\n", name, i.line)

	return result
}

// block formats a single block directive (e.g.`require (…)`). Returns an empty
// string if the given item slice is empty.
func block(name string, items []item) string {
	if len(items) == 0 {
		return ""
	}

	result := name + " (\n"
	for _, i := range items {
		result += comments(i.comments, "\t")
		result += fmt.Sprintf("\t%s\n", i.line)
	}

	result += ")\n"

	return result
}

// extractComments extracts, simplifies, and combines comment lines from the
// given modfile.Comment inputs. Returned lines will be stripped of the comment
// prefix (`//`).
func extractComments(sections ...[]modfile.Comment) []string {
	var lines []string

	for _, section := range sections {
		for _, comment := range section {
			if line := strings.TrimSpace(strings.TrimPrefix(comment.Token, "//")); line != "" {
				lines = append(lines, line)
			}
		}
	}

	return lines
}
