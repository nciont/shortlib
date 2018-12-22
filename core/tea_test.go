package core_test

import (
	"math"
	"math/big"
	"shortlib/core"
	"testing"
)

var TestKey64 = []uint32{0, 0, 0, 0}
var TestKey32 = []uint16{0, 0, 0, 0}

func TestTeaEncode64(t *testing.T) {
	var value uint64 = 0x0
	encrypted := core.TeaEncode64(value, TestKey64)
	decrypted := core.TeaDecode64(encrypted, TestKey64)
	t.Logf("%x -> %x -> %x", value, encrypted, decrypted)
}

func TestTeaEncode32(t *testing.T) {
	var value uint32 = 0x0
	encrypted := core.TeaEncode32(value, TestKey32)
	decrypted := core.TeaDecode32(encrypted, TestKey32)
	t.Logf("%x -> %x -> %x", value, encrypted, decrypted)
}

func TestTeaConstants(t *testing.T) {
	// Note: float precision of math.Pow(2, 64) / ((math.Sqrt(5) + 1) / 2)
	// is not enough to get correct result. Using the big lib instead.

	var s = "9e3779b97f4a7c15f39cc0605cedc834"
	//      |____________ 128 bit ___________|
	//      |____ 64 bit ____|
	//      |__ 32 __|
	//      |_16_|

	phi := big.NewFloat(5)
	phi.SetPrec(128)
	phi.Sqrt(phi)
	phi.Add(phi, big.NewFloat(1))
	phi.Quo(phi, big.NewFloat(2))
	phi.Quo(big.NewFloat(math.Pow(2, 128)), phi)

	expected, _ := new(big.Int).SetString(s, 16)
	if actual, _ := phi.Int(new(big.Int)); actual.Cmp(expected) != 0 {
		t.Errorf("Unexpected value: %x", actual)
	}
}
