package main

import (
	rand2 "crypto/rand"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"shortlib/core"
	"unicode/utf8"
)

const alphabet = "bcdfghijkmnpqrstvwxyz23456789BDFGHJKLMNPQRSTVWXYZCIUEAueaoOl10-_"

type params struct {
	bits   int
	abclen int
	keylen int
}

func init() {
	nBytes := make([]byte, 8, 8)
	if _, e := rand2.Read(nBytes); e != nil {
		panic(e)
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(nBytes)))
}

func main() {
	conf := cmd()

	if e := validate(conf); e != nil {
		fmt.Println(e)
		flag.Usage()
		os.Exit(1)
	}

	if conf.abclen == 0 {
		conf.abclen = getAbcLen(conf.keylen, conf.bits)
	} else {
		conf.keylen = getKeyLen(conf.abclen, conf.bits)
	}

	fmt.Printf("Parameters for %d bits:\n", conf.bits)
	fmt.Println("alphabet length:", conf.abclen)
	if conf.abclen <= maxGenAlphabetLength() {
		fmt.Printf("\t%s\n", genAlphabet(conf.abclen))
	}
	fmt.Println("ID key length:", conf.keylen)
	fmt.Println("Encryption key:")
	fmt.Printf("\t%s\n", genKey(conf.bits))
}

func cmd() *params {
	conf := new(params)

	flag.IntVar(&conf.bits, "bits", 32, "32 or 64 bit length")
	flag.IntVar(&conf.abclen, "abclen", 0, "Target alphabet length")
	flag.IntVar(&conf.keylen, "keylen", 0, "Target key length")
	flag.Parse()

	return conf
}

func genKey(bits int) string {
	if bits == 32 {
		if key, e := core.GenKey16(4); e != nil {
			panic(e)
		} else {
			return core.Stringify16(key)
		}
	}

	if key, e := core.GenKey32(4); e != nil {
		panic(e)
	} else {
		return core.Stringify32(key)
	}
}

func maxGenAlphabetLength() int {
	return utf8.RuneCountInString(alphabet)
}

func genAlphabet(length int) string {
	runes := []rune(alphabet)[:length]
	rand.Shuffle(length, func(i, j int) { runes[i], runes[j] = runes[j], runes[i] })
	return string(runes)
}

func getAbcLen(keylen int, bits int) int {
	exponent := big.NewInt(int64(keylen))
	var base int
	target := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(bits)), nil)
	result := new(big.Int)

	for base = 1; ; base++ {
		if result.Exp(big.NewInt(int64(base)), exponent, nil).Cmp(target) >= 0 {
			break
		}
	}

	return base
}

func getKeyLen(abclen int, bits int) int {
	var exponent int
	base := big.NewInt(int64(abclen))
	target := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(bits)), nil)
	result := new(big.Int)

	for exponent = 1; ; exponent++ {
		if result.Exp(base, big.NewInt(int64(exponent)), nil).Cmp(target) >= 0 {
			break
		}
	}

	return exponent
}

func validate(conf *params) error {
	if (conf.keylen == 0 && conf.abclen == 0) || (conf.keylen > 0 && conf.abclen > 0) {
		return errors.New("either keylen or abclen MUST be set")
	}

	if conf.bits != 32 && conf.bits != 64 {
		return errors.New("acceptable bits value: 32 or 64")
	}

	return nil
}
