// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

// Package modfmt contains functions for formatting `go.mod` and `go.work` files.
package modfmt

import (
	"bytes"
	"io"
	"slices"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/semver"
)

// Format attempts to parse and format the given data as either a `go.mod` or
// `go.work` file.
func Format(file string, data []byte) ([]byte, error) {
	// First, attempt to parse and format the given data as a go.mod file.
	formatted, errmod := FormatMod(file, data)
	if errmod == nil {
		return formatted, nil
	}

	// Second, attempt to parse and format the given data as a go.work file.
	formatted, errwork := FormatWork(file, data)
	if errwork == nil {
		return formatted, nil
	}

	// If both attempts failed, then return the error from the initial go.mod
	// attempt.
	return nil, errmod
}

// FormatMod attempts to parse and format the given data as a `go.mod` file.
func FormatMod(file string, data []byte) ([]byte, error) {
	mod, err := modfile.Parse(file, data, nil)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	formatMod(mod, &buf)

	return buf.Bytes(), nil
}

// formatMod updates, sorts, & formats the given modfile.File.
//
// See https://go.dev/ref/mod#go-mod-file
func formatMod(mod *modfile.File, w io.Writer) {
	// sort `exclude (…)` directives by module path.
	slices.SortFunc(mod.Exclude, func(a, b *modfile.Exclude) int {
		return strings.Compare(a.Mod.Path, b.Mod.Path)
	})

	// sort `godebug (…)` directives by key.
	slices.SortFunc(mod.Godebug, func(a, b *modfile.Godebug) int {
		return strings.Compare(a.Key, b.Key)
	})

	// sort `ignore (…)` directives by file path.
	slices.SortFunc(mod.Ignore, func(a, b *modfile.Ignore) int {
		return strings.Compare(a.Path, b.Path)
	})

	// sort `replace (…)` directives by module path, then by version.
	slices.SortFunc(mod.Replace, func(a, b *modfile.Replace) int {
		if cmp := strings.Compare(a.Old.Path, b.Old.Path); cmp != 0 {
			return cmp
		}

		return semver.Compare(a.Old.Version, b.Old.Version)
	})

	// sort `require (…)` directives by module path.
	slices.SortFunc(mod.Require, func(a, b *modfile.Require) int {
		return strings.Compare(a.Mod.Path, b.Mod.Path)
	})

	// sort `retract (…)` directives by version.
	slices.SortFunc(mod.Retract, func(a, b *modfile.Retract) int {
		if cmp := semver.Compare(a.Low, b.Low); cmp != 0 {
			return cmp
		}

		return semver.Compare(a.High, b.High)
	})

	// sort `tool (…)` directives by module path.
	slices.SortFunc(mod.Tool, func(a, b *modfile.Tool) int {
		return strings.Compare(a.Path, b.Path)
	})

	joinSections(w,
		sectionHeader(mod.Syntax),
		sectionModule(mod.Module),
		sectionGo(mod.Go),
		sectionToolchain(mod.Toolchain),
		sectionGodebug(mod.Godebug),
		sectionRetract(mod.Retract),
		sectionRequire(mod.Require),
		sectionRequireIndirect(mod.Require),
		sectionIgnore(mod.Ignore),
		sectionExclude(mod.Exclude),
		sectionReplace(mod.Replace),
		sectionReplaceLocal(mod.Replace),
		sectionTool(mod.Tool),
	)
}

// FormatWork attempts to parse and format the given data as a `go.work` file.
func FormatWork(file string, data []byte) ([]byte, error) {
	work, err := modfile.ParseWork(file, data, nil)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	formatWork(work, &buf)

	return buf.Bytes(), nil
}

// formatWork updates, sorts, & formats the given modfile.WorkFile.
//
// See https://go.dev/ref/mod#go-work-file
func formatWork(work *modfile.WorkFile, w io.Writer) {
	// sort `godebug (…)` directives by key.
	slices.SortFunc(work.Godebug, func(a, b *modfile.Godebug) int {
		return strings.Compare(a.Key, b.Key)
	})

	// sort `replace (…)` directives by module path, then by version.
	slices.SortFunc(work.Replace, func(a, b *modfile.Replace) int {
		if cmp := strings.Compare(a.Old.Path, b.Old.Path); cmp != 0 {
			return cmp
		}

		return semver.Compare(a.Old.Version, b.Old.Version)
	})

	// sort `use (…)` directives by file path.
	slices.SortFunc(work.Use, func(a, b *modfile.Use) int {
		return strings.Compare(a.Path, b.Path)
	})

	joinSections(w,
		sectionHeader(work.Syntax),
		sectionGo(work.Go),
		sectionToolchain(work.Toolchain),
		sectionGodebug(work.Godebug),
		sectionUse(work.Use),
		sectionReplace(work.Replace),
		sectionReplaceLocal(work.Replace),
	)
}

// joinSections writes each non-empty section to the given io.Writer with a
// newline between each written section.
func joinSections(w io.Writer, sections ...string) {
	var newline bool

	for _, section := range sections {
		if section == "" {
			continue
		}

		if newline {
			w.Write([]byte("\n")) //nolint:errcheck
		}

		w.Write([]byte(section)) //nolint:errcheck

		newline = true
	}
}
