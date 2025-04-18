// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cli

import (
	"io"
	"testing"
)

func TestMockUi_implements(t *testing.T) {
	var _ Ui = new(MockUi)
}

func TestMockUi_Ask(t *testing.T) {
	tests := []struct {
		name           string
		query, input   string
		expectedResult string
	}{
		{"EmptyString", "Middle Name?", "\n", ""},
		{"NonEmptyString", "Name?", "foo bar\nbaz\n", "foo bar"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in_r, in_w := io.Pipe()
			defer in_r.Close()
			defer in_w.Close()

			ui := &MockUi{
				InputReader: in_r,
			}

			errors := make(chan error, 1)
			go func() {
				_, err := in_w.Write([]byte(tc.input))
				errors <- err
			}()

			result, err := ui.Ask(tc.query)
			if err != nil {
				t.Fatalf("err: %s", err)
			}

			err = <-errors
			if err != nil {
				t.Fatalf("err: %v", err)
			}

			if result != tc.expectedResult {
				t.Fatalf("bad: %#v", result)
			}
		})
	}
}
