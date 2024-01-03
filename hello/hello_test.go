package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	a, b := Output()
	fmt.Println(a, b)

	assert.Equal(t, "olleH", a)
	assert.Equal(t, 1234, b)

}
