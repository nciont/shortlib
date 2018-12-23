package shortlib

import (
	"errors"
	"github.com/nciont/shortlib/core"
	"github.com/nciont/shortlib/rebase"
	"unicode/utf8"
)

type Transform32 struct {
	Alphabet []rune
	Idlen    int
	Key      []uint16
}

func (t *Transform32) Encode(n uint32) string {
	enc := core.TeaEncode32(n, t.Key)
	return string(rebase.PadRebase(uint64(enc), t.Alphabet, t.Idlen))
}

func (t *Transform32) Decode(src string) (uint32, error) {
	if utf8.RuneCountInString(src) != t.Idlen {
		return 0, errors.New("invalid string length")
	}

	enc, e := rebase.Unbase(src, t.Alphabet)

	if e != nil {
		return 0, e
	}

	return core.TeaDecode32(uint32(enc), t.Key), nil
}

type Transform64 struct {
	Alphabet []rune
	Idlen    int
	Key      []uint32
}

func (t *Transform64) Encode(n uint64) string {
	enc := core.TeaEncode64(n, t.Key)
	return string(rebase.PadRebase(enc, t.Alphabet, t.Idlen))
}

func (t *Transform64) Decode(src string) (uint64, error) {
	if utf8.RuneCountInString(src) != t.Idlen {
		return 0, errors.New("invalid string length")
	}

	enc, e := rebase.Unbase(src, t.Alphabet)

	if e != nil {
		return 0, e
	}

	return core.TeaDecode64(enc, t.Key), nil
}
