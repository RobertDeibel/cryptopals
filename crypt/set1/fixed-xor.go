package crypt

import (
	"encoding/hex"
	"fmt"
)

type lengthMismatchError struct {
	first, second int
}

func (e lengthMismatchError) Error() string {
	return fmt.Sprintf("lengths %v and %v do not match", e.first, e.second)
}

// FixedXorStr takes hex strings of same length and xors them after decoding
// returns non nil error when strings not match or decoding fails
func FixedXorStr(first, second string) ([]byte, string, error) {

	decodedFirst, err := hex.DecodeString(first)
	if err != nil {
		return nil, "", err
	}
	decodedSecond, err := hex.DecodeString(second)
	if err != nil {
		return nil, "", err
	}

	xor, err := FixedXor(decodedFirst, decodedSecond)
	if err != nil {
		return nil, "", err
	}
	return xor, hex.EncodeToString(xor), nil
}

// FixedXor takes hex byte arrays of same length and xors them after decoding
func FixedXor(first, second []byte) ([]byte, error) {
	if len(first) != len(second) {
		return nil, lengthMismatchError{len(first), len(second)}
	}
	xor := make([]byte, len(first))
	for i := range first {
		xor[i] = first[i] ^ second[i]
	}
	return xor, nil
}
