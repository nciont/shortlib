package rebase

import (
	"errors"
	"math"
)

// Rebase converts n to alphabet length base
func Rebase(n uint64, alphabet []rune) (digits []rune) {
	if n == 0 {
		return []rune{alphabet[0]}
	}

	var base = uint64(len(alphabet))

	// Precalc number of digits: log x base b = log x / log b
	digits = make([]rune, int(math.Ceil(math.Log1p(float64(n))/math.Log(float64(base)))))

	for idx := len(digits) - 1; idx >= 0; idx-- {
		digits[idx] = alphabet[n%base]
		n /= base
	}

	return
}

// PadRebase converts n to alphabet-length base with padding to specified width
func PadRebase(n uint64, alphabet []rune, min int) []rune {
	result := Rebase(n, alphabet)
	padLen := min - len(result)

	if padLen <= 0 {
		return result
	}

	paddedResult := make([]rune, padLen, padLen+len(result))
	for k := 0; k < padLen; k++ {
		paddedResult[k] = alphabet[0]
	}
	return append(paddedResult, result...)
}

// Convert a number from alphabet-length base to int
func Unbase(s string, alphabet []rune) (uint64, error) {
	var n uint64
	var base = uint64(len(alphabet))
	m := make(map[rune]int)

	for i, c := range alphabet { // TODO build map once and reuse
		m[c] = i
	}

	sm := len(s) - 1
	for i, c := range s {
		exp := sm - i

		if b, ok := m[c]; false == ok {
			return 0, errors.New("invalid character found")
		} else {
			n += uint64(b) * pow(base, uint64(exp))
		}
	}

	return n, nil
}

func pow(a, b uint64) uint64 {
	var p uint64 = 1

	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}

	return p
}