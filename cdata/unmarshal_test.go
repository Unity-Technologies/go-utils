package cdata

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalBinary(t *testing.T) {
	expected := Str1{
		A: Str2{
			B: 1,
			C: 2,
		},
		D: Str3{
			E: 3,
			F: 4,
			G: Str4{
				H: 5,
				I: 6,
			},
			J: 7,
			K: 8,
		},
		L: 9,
		M: 10,
	}

	cases := []struct {
		name   string
		buf    []byte
		padded bool
	}{
		{"padded", paddedStr1, true},
		{"unpadded", unpaddedStr1, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data := Str1{}
			err := UnmarshalBinary(tc.buf, &data, binary.LittleEndian, tc.padded)
			if assert.NoError(t, err) {
				assert.Equal(t, expected, data)
			}
		})
	}
}
