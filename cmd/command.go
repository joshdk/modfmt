// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

// Package cmd provides the command line handler for modfmt.
package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joshdk/buildversion"
	"github.com/spf13/cobra"

	"github.com/joshdk/modfmt/pkg/modfmt"
)

// Command returns a complete command line handler for modfmt.
func Command() *cobra.Command { //nolint:cyclop,funlen
	cmd := &cobra.Command{
		Use:     "modfmt [directory|file]",
		Long:    "modfmt - formatter for go.mod and go.work files",
		Version: "-",

		SilenceUsage:  true,
		SilenceErrors: true,
	}

	// Set a custom list of examples.
	cmd.Example = strings.TrimRight(exampleText, "\n")

	// Add a custom usage footer template.
	cmd.SetUsageTemplate(cmd.UsageTemplate() + "\n" + buildversion.Template(usageTemplate))

	// Set a custom version template.
	cmd.SetVersionTemplate(buildversion.Template(versionTemplate))

	// Define --check/-c flag.
	check := cmd.Flags().BoolP(
		"check", "c",
		false,
		"exit with code 1 if any files were unformatted")

	// Define --list/-l flag.
	list := cmd.Flags().BoolP(
		"list", "l",
		false,
		"list any files that were unformatted")

	// Define --write/-w flag.
	write := cmd.Flags().BoolP(
		"write", "w",
		false,
		"write result to (source) file instead of stdout")

	cmd.RunE = func(_ *cobra.Command, args []string) error {
		// If no arguments are given, default to recursively searching through
		// the current working directory.
		if len(args) == 0 {
			args = []string{"."}
		}

		// Search for go.mod and go.work files.
		filenames, err := discover(args)
		if err != nil {
			return err
		}

		var unformatted bool

		for _, filename := range filenames {
			// Read the original file.
			original, err := os.ReadFile(filename)
			if err != nil {
				return err
			}

			// Format the file.
			formatted, err := modfmt.Format(filename, original)
			if err != nil {
				return err
			}

			// Did formatting change the file or was it already formatted?
			if bytes.Equal(original, formatted) {
				continue
			}

			unformatted = true

			if *write {
				// If write mode was requested, then silently update the
				// original file.
				if err := os.WriteFile(filename, formatted, 0o644); err != nil {
					return err
				}
			}

			if *list {
				// If list mode was requested, then list files.
				fmt.Println(filename)
			} else if !*write {
				// Otherwise if write mode was not requested, then print
				// filename header and formatted body.
				fmt.Printf("--- %s ---\n", filename)
				fmt.Print(string(formatted))
			}
		}

		if *check && unformatted {
			// If check mode was requested and any files were unformatted, then
			// exit with an error.
			return errors.New("some files were unformatted")
		}

		return nil
	}

	return cmd
}

// discover returns a list of `go.mod` and `go.work` file paths based on the
// given specs. Each spec can be one of the following:
//   - If an explicit file name is given, it will be returned verbatim.
//   - If a directory name is given, any directly contained `go.mod` or
//     `go.work` files are returned.
//   - If the given spec ends with `/...` then it is treated as a directory and
//     walked in search of any `go.mod` or `go.work` files.
func discover(specs []string) ([]string, error) { //nolint:cyclop
	var results []string

	for _, spec := range specs {
		if directory, ok := strings.CutSuffix(spec, "/..."); ok {
			// If the spec ends with `/...` then walk through the directory.
			if err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
				switch {
				case err != nil:
					// There was an actual error.
					return err

				case info.IsDir():
					switch info.Name() {
					case ".git":
						// Ignore `.git/` directory.
						return filepath.SkipDir
					case "vendor":
						// Ignore `vendor/` directory.
						return filepath.SkipDir
					}

				case info.Name() == "go.mod":
					// Found a `go.mod` file!
					results = append(results, path)

				case info.Name() == "go.work":
					// Found a `go.work` file!
					results = append(results, path)
				}

				return nil
			}); err != nil {
				return nil, err
			}

			continue
		}

		stat, err := os.Stat(spec)
		if err != nil {
			return nil, err
		}

		// A file was named explicitly.
		if !stat.IsDir() {
			results = append(results, spec)

			continue
		}

		// A directory was named explicitly. Try checking for a `go.mod` file.
		modpath := filepath.Join(spec, "go.mod")
		if _, err := os.Stat(modpath); err == nil {
			results = append(results, modpath)
		}

		// Finally try checking for a `go.work` file.
		workpath := filepath.Join(spec, "go.work")
		if _, err := os.Stat(workpath); err == nil {
			results = append(results, workpath)
		}
	}

	return results, nil
}
