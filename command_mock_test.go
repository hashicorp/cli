// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: MPL-2.0

package cli

import (
	"testing"
)

func TestMockCommand_implements(t *testing.T) {
	var _ Command = new(MockCommand)
}
