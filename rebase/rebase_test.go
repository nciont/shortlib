package rebase_test

import (
	"github.com/nciont/shortlib/rebase"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type RebaseTestCase struct {
	expected string
	number   uint64
	alphabet []rune
}

var alphaBinary = []rune{'0', '1'}
var alphaDecimal = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var alphaHex = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func TestRebase(t *testing.T) {
	testcases := []RebaseTestCase{
		// actual, expected, alphabet
		{"11", 3, alphaBinary},
		{"1111111111111111111111111111111111111111111111111111111111111111", 1<<64 - 1, alphaBinary},
		{"999", 999, alphaDecimal},
		{"3e7", 999, alphaHex},
	}

	for idx, data := range testcases {
		if actual := string(rebase.Rebase(data.number, data.alphabet)); actual != data.expected {
			t.Errorf("Test %d: Expected %s, got %s", idx, data.expected, actual)
		}
	}
}

func TestRebase1000(t *testing.T) {
	var k uint64

	for k = 0; k < 1001; k++ {
		strK := strconv.Itoa(int(k))
		if actual := string(rebase.Rebase(k, alphaDecimal)); actual != strK {
			t.Errorf("Test %d: Expected %s, got %s", k, strK, actual)
		}
	}
}

func TestPadRebase(t *testing.T) {
	actual := string(rebase.PadRebase(452, alphaHex, 8))
	if actual != "000001c4" {
		t.Error("Incorrect result", actual)
	}
}

func TestUnbase(t *testing.T) {
	if actual := rebase.Unbase("ff", alphaHex); actual != 255 {
		t.Error(actual)
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Uint64()
	st := string(rebase.PadRebase(n, alphaHex, 40))
	d := rebase.Unbase(st, alphaHex)

	if n != d {
		t.Error("Unequal", n, d)
	}
}
