package simdjson

import (
	"fmt"
	"testing"
	"encoding/binary"
)

func TestParseString(t *testing.T) {

	stringbuf := make([]byte, 0, 256)

	const str = "key"
	size := parse_string_simd([]byte(fmt.Sprintf(`"%s":`, str)), &stringbuf)

	// First four bytes are size
	length := int(binary.LittleEndian.Uint32(stringbuf[0:4]))
	if length != len(str) {
		t.Errorf("TestParseString: got: %d want: %d", length, len(str))
	}
	// Then comes value of string
	if string(stringbuf[4:4+length]) != str {
		t.Errorf("TestParseString: got: %s want: %s", string(stringbuf[4:4+length]), str)
	}
	// Followed by NULL-character
	if stringbuf[4+length] != 0 {
		t.Errorf("TestParseString: got: 0x%x want: 0x0", stringbuf[4+length])
	}
	if size != 4 + length + 1 {
		t.Errorf("TestParseString: got: %d want: %d", size, 4 + length + 1)
	}

	const str2 = "value"
	size2 := parse_string_simd([]byte(fmt.Sprintf(`"%s":`, str2)), &stringbuf)

	// First four bytes are size
	length = int(binary.LittleEndian.Uint32(stringbuf[size:size+4]))
	if length != len(str2) {
		t.Errorf("TestParseString: got: %d want: %d", length, len(str2))
	}
	// Then comes value of string
	if string(stringbuf[size+4:size+4+length]) != str2 {
		t.Errorf("TestParseString: got: %s want: %s", string(stringbuf[size+4:size+4+length]), str2)
	}
	// Followed by NULL-character
	if stringbuf[size+4+length] != 0 {
		t.Errorf("TestParseString: got: 0x%x want: 0x0", stringbuf[size+4+length])
	}
	if size2 != 4 + length + 1 {
		t.Errorf("TestParseString: got: %d want: %d", size2, 4 + length + 1)
	}
}