// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cli

import (
	"syscall/js"
)

func (u *BasicUi) ask(query string, secret bool) (string, error) {
	line := js.Global().Call("prompt", query).String()
	return line, nil
}
