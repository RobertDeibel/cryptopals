package main

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

// FixedXor takes hex strings of same length and xors them after decoding
// returns non nil error when strings not match or decoding fails
func FixedXor(first, second string) ([]byte, string, error) {
	if len(first) != len(second) {
		return nil, "", lengthMismatchError{len(first), len(second)}
	}

	decodedFirst, err := hex.DecodeString(first)
	if err != nil {
		return nil, "", err
	}
	decodedSecond, err := hex.DecodeString(second)
	if err != nil {
		return nil, "", err
	}

	xor := make([]byte, len(decodedFirst))
	for i := range decodedFirst {
		xor[i] = decodedFirst[i] ^ decodedSecond[i]
	}
	return xor, hex.EncodeToString(xor), nil
}
