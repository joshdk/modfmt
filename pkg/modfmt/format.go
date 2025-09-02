// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

// Package modfmt contains functions for formatting `go.mod` and `go.work` files.
package modfmt

import (
	"bytes"
	"io"
	"sort"

	"golang.org/x/mod/modfile"
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
	sort.Slice(mod.Exclude, func(i, j int) bool {
		return mod.Exclude[i].Mod.Path < mod.Exclude[j].Mod.Path
	})

	sort.Slice(mod.Godebug, func(i, j int) bool {
		return mod.Godebug[i].Key < mod.Godebug[j].Key
	})

	sort.Slice(mod.Ignore, func(i, j int) bool {
		return mod.Ignore[i].Path < mod.Ignore[j].Path
	})

	sort.Slice(mod.Replace, func(i, j int) bool {
		return mod.Replace[i].Old.Path < mod.Replace[j].Old.Path
	})

	sort.Slice(mod.Require, func(i, j int) bool {
		return mod.Require[i].Mod.Path < mod.Require[j].Mod.Path
	})

	sort.Slice(mod.Retract, func(i, j int) bool {
		return mod.Retract[i].Low < mod.Retract[j].Low
	})

	sort.Slice(mod.Tool, func(i, j int) bool {
		return mod.Tool[i].Path < mod.Tool[j].Path
	})

	joinSections(w,
		sectionHeader(mod.Syntax),
		sectionModule(mod.Module),
		sectionGo(mod.Go),
		sectionToolchain(mod.Toolchain),
		sectionGodebug(mod.Godebug),
		sectionTool(mod.Tool),
		sectionRequire(mod.Require),
		sectionRequireIndirect(mod.Require),
		sectionIgnore(mod.Ignore),
		sectionExclude(mod.Exclude),
		sectionReplace(mod.Replace),
		sectionReplaceLocal(mod.Replace),
		sectionRetract(mod.Retract),
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
	sort.Slice(work.Godebug, func(i, j int) bool {
		return work.Godebug[i].Key < work.Godebug[j].Key
	})

	sort.Slice(work.Replace, func(i, j int) bool {
		return work.Replace[i].Old.Path < work.Replace[j].Old.Path
	})

	sort.Slice(work.Use, func(i, j int) bool {
		return work.Use[i].Path < work.Use[j].Path
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
