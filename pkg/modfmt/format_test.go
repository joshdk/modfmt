// Copyright Josh Komoroske. All rights reserved.
// Use of this source code is governed by the MIT license,
// a copy of which can be found in the LICENSE.txt file.
// SPDX-License-Identifier: MIT

package modfmt_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/joshdk/modfmt/pkg/modfmt"
)

const testdataDir = "./testdata"

func TestFormat(t *testing.T) {
	t.Parallel()

	entries, err := os.ReadDir(testdataDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".mod") && !strings.HasSuffix(entry.Name(), ".work") {
			continue
		}

		t.Run(entry.Name(), func(t *testing.T) {
			t.Parallel()

			originalFile := filepath.Join(testdataDir, entry.Name())

			originalData, err := os.ReadFile(originalFile)
			if err != nil {
				t.Fatal(err)
			}

			formattedFile := filepath.Join(testdataDir, entry.Name()+".formatted")

			expectedData, err := os.ReadFile(formattedFile)
			if err != nil {
				t.Fatal(err)
			}

			actualData, err := modfmt.Format(originalFile, originalData)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(expectedData, actualData) {
				t.Fatalf("formatted %s differed from %s", originalFile, formattedFile)
			}
		})
	}
}
