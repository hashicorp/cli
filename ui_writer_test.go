// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cli

import (
	"io"
	"testing"
)

func TestUiWriter_impl(t *testing.T) {
	var _ io.Writer = new(UiWriter)
}

func TestUiWriter(t *testing.T) {
	ui := new(MockUi)
	w := &UiWriter{
		Ui: ui,
	}

	if _, err := w.Write([]byte("foo\n")); err != nil {
		t.Fatalf("failed to write the bytes: %v", err)
	}
	if _, err := w.Write([]byte("bar\n")); err != nil {
		t.Fatalf("failed to write the bytes: %v", err)
	}

	if ui.OutputWriter.String() != "foo\nbar\n" {
		t.Fatalf("bad: %s", ui.OutputWriter.String())
	}
}

func TestUiWriter_empty(t *testing.T) {
	ui := new(MockUi)
	w := &UiWriter{
		Ui: ui,
	}

	if _, err := w.Write([]byte("")); err != nil {
		t.Fatalf("failed to write the bytes: %v", err)
	}

	if ui.OutputWriter.String() != "\n" {
		t.Fatalf("bad: %s", ui.OutputWriter.String())
	}
}
