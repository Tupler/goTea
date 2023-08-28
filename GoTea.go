package goTea

import (
	"encoding/binary"
)

const (
	teaRounds = 32
	delta     = 0x9e3779b9
)

func Encrypt(textplain []byte, key []byte) {
	v0 := binary.BigEndian.Uint32(textplain[:4])
	v1 := binary.BigEndian.Uint32(textplain[4:])
	sum := uint32(0)
	key0 := binary.BigEndian.Uint32(key[:4])
	key1 := binary.BigEndian.Uint32(key[4:8])
	key2 := binary.BigEndian.Uint32(key[8:12])
	key3 := binary.BigEndian.Uint32(key[12:])

	for i := 0; i < teaRounds; i++ {
		sum += delta
		v0 += ((v1 << 4) + key0) ^ (v1 + sum) ^ ((v1 >> 5) + key1)
		v1 += ((v0 << 4) + key2) ^ (v0 + sum) ^ ((v0 >> 5) + key3)
	}
	binary.BigEndian.PutUint32(textplain[:4], v0)
	binary.BigEndian.PutUint32(textplain[4:], v1)
}

func Decrypt(textplain []byte, key []byte) {
	v0 := binary.BigEndian.Uint32(textplain[:4])
	v1 := binary.BigEndian.Uint32(textplain[4:])
	sum := uint32(0xc6ef3720)
	key0 := binary.BigEndian.Uint32(key[:4])
	key1 := binary.BigEndian.Uint32(key[4:8])
	key2 := binary.BigEndian.Uint32(key[8:12])
	key3 := binary.BigEndian.Uint32(key[12:])
	for i := 0; i < teaRounds; i++ {
		v1 -= ((v0 << 4) + key2) ^ (v0 + sum) ^ ((v0 >> 5) + key3)
		v0 -= ((v1 << 4) + key0) ^ (v1 + sum) ^ ((v1 >> 5) + key1)
		sum -= delta
	}
	binary.BigEndian.PutUint32(textplain[:4], v0)
	binary.BigEndian.PutUint32(textplain[4:], v1)

}
