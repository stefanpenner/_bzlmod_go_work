// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reverse_test

import (
	"golang.org/x/example/hello/reverse"
  "testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleString(t *testing.T) {
	assert.Equal(t, reverse.String("hello"), "olleh")
}
