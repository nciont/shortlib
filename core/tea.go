package core

// https://github.com/HowardHinnant/hash_append/issues/7
// https://softwareengineering.stackexchange.com/a/63605

// TeaNumRounds number of rounds in Tiny Encryption Algorithm
const TeaNumRounds = 32

// TeaDelta64 TEA encryption constant
const TeaDelta64 uint32 = 0x9E3779B9

// TeaDecryptSum64 TEA decryption constant
const TeaDecryptSum64 uint32 = 0xc6ef3720 // 0x9E3779B9 * 32

// TeaDelta32 TEA encryption constant
const TeaDelta32 uint16 = 0x9E37

// TeaDecryptSum32 TEA decryption constant
const TeaDecryptSum32 uint16 = 0xc6e0 // 0x9E37 * 32

// TeaEncode64 encrypts n using Tiny Encryption Algorithm
func TeaEncode64(n uint64, key []uint32) uint64 {
	var sum uint32
	v0 := uint32((n >> 32) & 0xffffffff)
	v1 := uint32(n & 0xffffffff)

	for i := 0; i < TeaNumRounds; i++ {
		sum += TeaDelta64
		v0 += ((v1 << 4) + key[0]) ^ (v1 + sum) ^ ((v1 >> 5) + key[1])
		v1 += ((v0 << 4) + key[2]) ^ (v0 + sum) ^ ((v0 >> 5) + key[3])
	}

	return (uint64(v0) << 32) | uint64(v1)
}

// TeaDecode64 decrypts n using Tiny Encryption Algorithm
func TeaDecode64(n uint64, key []uint32) uint64 {
	var sum = TeaDecryptSum64
	v0 := uint32((n >> 32) & 0xffffffff)
	v1 := uint32(n & 0xffffffff)

	for i := 0; i < TeaNumRounds; i++ {
		v1 -= ((v0 << 4) + key[2]) ^ (v0 + sum) ^ ((v0 >> 5) + key[3])
		v0 -= ((v1 << 4) + key[0]) ^ (v1 + sum) ^ ((v1 >> 5) + key[1])
		sum -= TeaDelta64
	}

	return (uint64(v0) << 32) | uint64(v1)
}

// TeaEncode32 encrypts n using Tiny Encryption Algorithm
func TeaEncode32(n uint32, key []uint16) uint32 {
	var sum uint16
	v0 := uint16((n >> 16) & 0xffff)
	v1 := uint16(n & 0xffff)

	for i := 0; i < TeaNumRounds; i++ {
		sum += TeaDelta32
		v0 += ((v1 << 2) + key[0]) ^ (v1 + sum) ^ ((v1 >> 3) + key[1])
		v1 += ((v0 << 2) + key[2]) ^ (v0 + sum) ^ ((v0 >> 3) + key[3])
	}

	return (uint32(v0) << 16) | uint32(v1)
}

// TeaDecode32 decrypts n using Tiny Encryption Algorithm
func TeaDecode32(n uint32, key []uint16) uint32 {
	var sum = TeaDecryptSum32
	v0 := uint16((n >> 16) & 0xffff)
	v1 := uint16(n & 0xffff)

	for i := 0; i < TeaNumRounds; i++ {
		v1 -= ((v0 << 2) + key[2]) ^ (v0 + sum) ^ ((v0 >> 3) + key[3])
		v0 -= ((v1 << 2) + key[0]) ^ (v1 + sum) ^ ((v1 >> 3) + key[1])
		sum -= TeaDelta32
	}

	return (uint32(v0) << 16) | uint32(v1)
}
