// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	} {
		got := String(c.in)
		assert.Equal(t, c.want, got)
	}
}

func TestInt(t *testing.T) {
	for _, c := range []struct {
		in, want int
	}{
		{1234, 4321},
		{123456789, 987654321},
		{0, 0},
	} {
		got := Int(c.in)
		assert.Equal(t, c.want, got)
	}
}
