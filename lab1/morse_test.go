package main

import (
	"bytes"
	"testing"
)

func compareByteSlices(a, b []byte) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

var apple = []byte{0x00, 0x17, 0x05, 0xdd, 0x05, 0xdd, 0x01, 0x75, 0x00, 0x01}

func TestEncode(t *testing.T) {
	if len(encode("")) != 0 {
		t.Error(`encode("") != []`)
	}
	if !compareByteSlices(encode("apple"), apple) {
		t.Error(`return value from encode("apple") is not correct`)
	}
}

func TestDecode(t *testing.T) {
	if decode(bytes.NewReader([]byte{})) != "" {
		t.Error(`decode(\*empty reader*\) does not return ""`)
	}
	if decode(bytes.NewReader(apple)) != "apple" {
		t.Error(`decode(\*encoded string "apple"*\) != "apple"`)
	}
}
