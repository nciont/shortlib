package core_test

import (
	"shortlib/core"
	"testing"
)

func TestGenKey16(t *testing.T) {
	var (
		key []uint16
		e   error
	)

	key, e = core.GenKey16(core.TeaKeyLength)

	if e != nil {
		t.Fatal(e)
	}

	t.Log("key", key)

	b64 := core.Stringify16(key)

	t.Log("b64", b64)

	parsed, e := core.Parse16(b64)

	if e != nil {
		t.Fatal(e)
	}

	if len(parsed) != len(key) {
		t.Fatal("Key parse failed")
	}

	for i, sub := range parsed {
		if sub != key[i] {
			t.Fatal("Keys not equal")
		}
	}
}

func TestGenKey32(t *testing.T) {
	var (
		key []uint32
		e   error
	)

	key, e = core.GenKey32(core.TeaKeyLength)

	if e != nil {
		t.Fatal(e)
	}

	t.Log("key", key)

	b64 := core.Stringify32(key)

	t.Log("b64", b64)

	parsed, e := core.Parse32(b64)

	if e != nil {
		t.Fatal(e)
	}

	if len(parsed) != len(key) {
		t.Fatal("Key parse failed")
	}

	for i, sub := range parsed {
		if sub != key[i] {
			t.Fatal("Keys not equal")
		}
	}
}
