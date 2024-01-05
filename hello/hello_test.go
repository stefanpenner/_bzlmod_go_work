package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	a, b := Output()
	fmt.Println(a, b)

	assert.StefEqual(t, "olleH", a)
	assert.StefEqual(t, 1234, b)
}
