// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

// Package cmd provides the command line handler for modfmt.
package cmd

import (
	"github.com/joshdk/buildversion"
	"github.com/spf13/cobra"
)

// Command returns a complete command line handler for modfmt.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "modfmt [directory|file]",
		Long:    "modfmt - formatter for go.mod and go.work files",
		Version: "-",

		SilenceUsage:  true,
		SilenceErrors: true,
	}

	// Add a custom usage footer template.
	cmd.SetUsageTemplate(cmd.UsageTemplate() + "\n" + buildversion.Template(usageTemplate))

	// Set a custom version template.
	cmd.SetVersionTemplate(buildversion.Template(versionTemplate))

	cmd.RunE = func(*cobra.Command, []string) error {
		return nil
	}

	return cmd
}
