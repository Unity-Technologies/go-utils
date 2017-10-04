package cdata

// Int8sToString takes C style string, []int8 null terminated, and converts it to a go string.
func Int8sToString(data []int8) string {
	b := make([]byte, len(data))
	var i int
	var d int8
	for i, d = range data {
		if d == 0 {
			break
		}
		b[i] = byte(d)
	}
	return string(b[:i])
}

// BytesToString takes C style string, []byte null terminated, and converts it to a go string.
func BytesToString(data []byte) string {
	var i int
	var d byte
	for i, d = range data {
		if d == 0 {
			break
		}
	}
	return string(data[:i])
}
