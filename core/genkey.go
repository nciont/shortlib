package core

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
)

const TeaKeyLength = 4

// GenKey32 generates random int slices of specified length
func GenKey32(length int) ([]uint32, error) {
	key := make([]uint32, length, length)
	keyByteLength := length * 4
	keyBytes := make([]byte, keyByteLength, keyByteLength)

	if _, e := rand.Read(keyBytes); e != nil {
		return nil, e
	}

	for i := 0; i < length; i++ {
		key[i] = binary.LittleEndian.Uint32(keyBytes[4*i:])
	}

	return key, nil
}

// Parse32 parses base64 representation of 32-bit version of the key
func Parse32(s string) ([]uint32, error) {
	var (
		bytes []byte
		e     error
	)

	if bytes, e = base64.StdEncoding.DecodeString(s); e != nil {
		return nil, e
	}

	if len(bytes)%4 != 0 {
		return nil, errors.New("invalid length")
	}

	out := make([]uint32, len(bytes)/4, len(bytes)/4)

	for i := 0; i < len(out); i++ {
		out[i] = binary.LittleEndian.Uint32(bytes[4*i:])
	}

	return out, nil
}

// Stringify32 converts key into base64 string
func Stringify32(key []uint32) string {
	bytes := make([]byte, len(key)*4, len(key)*4)

	for i, n := range key {
		binary.LittleEndian.PutUint32(bytes[4*i:], n)
	}

	return base64.StdEncoding.EncodeToString(bytes)
}

// GenKey16 generates random int slices of specified length
func GenKey16(length int) ([]uint16, error) {
	key := make([]uint16, length, length)
	keyByteLength := length * 2
	keyBytes := make([]byte, keyByteLength, keyByteLength)

	if _, e := rand.Read(keyBytes); e != nil {
		return nil, e
	}

	for i := 0; i < length; i++ {
		key[i] = binary.LittleEndian.Uint16(keyBytes[2*i:])
	}

	return key, nil
}

// Parse16 parses base64 representation of 16-bit version of the key
func Parse16(s string) ([]uint16, error) {
	var (
		bytes []byte
		e     error
	)

	if bytes, e = base64.StdEncoding.DecodeString(s); e != nil {
		return nil, e
	}

	if len(bytes)%2 != 0 {
		return nil, errors.New("invalid length")
	}

	out := make([]uint16, len(bytes)/2, len(bytes)/2)

	for i := 0; i < len(out); i++ {
		out[i] = binary.LittleEndian.Uint16(bytes[2*i:])
	}

	return out, nil
}

// Stringify16 converts key into base64 string
func Stringify16(key []uint16) string {
	bytes := make([]byte, len(key)*2, len(key)*2)

	for i, n := range key {
		binary.LittleEndian.PutUint16(bytes[2*i:], n)
	}

	return base64.StdEncoding.EncodeToString(bytes)
}
