package shortlib_test

import (
	"math/rand"
	"shortlib"
	"shortlib/core"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	m.Run()
}

func TestTransform32(t *testing.T) {
	n := rand.Uint32()

	key, e := core.GenKey16(core.TeaKeyLength)
	if e != nil {
		t.Fatal(e)
	}

	tr := &shortlib.Transform32{
		Key: key,
		Alphabet: []rune("qnKDF6Qfid43ys2jb95cgJzLxtw7hP8rpGvmBHNMk"),
		Idlen: 6,
	}

	enc := tr.Encode(uint32(n))
	dec, e := tr.Decode(enc)

	if e != nil {
		t.Fatal(e)
	}

	t.Logf("%d -> %s -> %d", n, enc, dec)

	if n != dec {
		t.Error("Unequal", n, dec)
	}
}

func TestTransform64(t *testing.T) {
	n := rand.Uint64()

	key, e := core.GenKey32(core.TeaKeyLength)
	if e != nil {
		t.Fatal(e)
	}

	tr := &shortlib.Transform64{
		Key: key,
		Alphabet: []rune("BNwEvKqdec4M9kXQrCAHpFSRfPj5YaxGDbV2JT8sUWIu7t3nigyhZ6Lmz"),
		Idlen: 11,
	}

	enc := tr.Encode(n)
	dec, e := tr.Decode(enc)

	if e != nil {
		t.Fatal(e)
	}

	t.Logf("%d -> %s -> %d", n, enc, dec)

	if n != dec {
		t.Error("Unequal", n, dec)
	}
}

func TestTransform32Randomness(t *testing.T) {
	key, e := core.GenKey16(core.TeaKeyLength)
	if e != nil {
		t.Fatal(e)
	}

	tr := &shortlib.Transform32{
		Key: key,
		Alphabet: []rune("qnKDF6Qfid43ys2jb95cgJzLxtw7hP8rpGvmBHNMk"),
		Idlen: 6,
	}

	for k := 0; k < 2 << 4; k++ {
		t.Logf("%d -> %s", k, tr.Encode(uint32(k)))
	}
}
