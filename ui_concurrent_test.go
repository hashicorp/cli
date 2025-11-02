// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: MPL-2.0

package cli

import (
	"testing"
)

func TestConcurrentUi_impl(t *testing.T) {
	var _ Ui = new(ConcurrentUi)
}
