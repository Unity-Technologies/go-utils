package cdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt8sToString(t *testing.T) {
	assert.Equal(t, "test", Int8sToString([]int8{116, 101, 115, 116, 0}))
}

func TestBytesToString(t *testing.T) {
	assert.Equal(t, "test", BytesToString([]byte{116, 101, 115, 116, 0}))
}
