package cdata

import (
	"bytes"
	"encoding/binary"
)

// UnmarshalBinary copies the structured binary data from buf into data using
// the given ByteOrder. If padded is true then padding, according to the memory
// layout of data, is removed first.
// This can be used to unmarshal data from a byte slice that represents a native
// C struct into its corresponding golang struct.
// Data must be a pointer to a fixed-size value or a slice of fixed-size values.
// Bytes read from r are decoded using the specified byte order and written to
// successive fields of the data. When decoding boolean values, a zero byte is
// decoded as false, and any other non-zero byte is decoded as true. When filling
// into structs, the field data for fields with blank (_) field names is skipped;
// i.e., blank field names may be used for padding. When reading into a struct,
// all non-blank fields must be exported or Fill may panic.
func UnmarshalBinary(buf []byte, data interface{}, order binary.ByteOrder, padded bool) error {
	if padded {
		buf = Unpad(buf, data)
	}
	return binary.Read(bytes.NewReader(buf), order, data)
}
