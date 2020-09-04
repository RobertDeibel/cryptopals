package crypt

import (
	"encoding/base64"
	"encoding/hex"
)

// HexToBase64 converts a Hex string to a Base64 byte array and string
// returns the base64 byte array and string, on error returns a non nil error and
// the zero-values of the byte array and string
func HexToBase64(encodedHex string) ([]byte, string, error) {

	hexDec, err := hex.DecodeString(encodedHex)
	if err != nil {
		return nil, "", err
	}
	base64Dec := Base64(hexDec)
	base64Str := string(base64Dec)
	return base64Dec, base64Str, nil
}

// Base64 converts a byte array to a Base64 byte array
// returns the base64 byte array and string, on error returns a non nil error and
// the zero-values of the byte array and string
func Base64(buffer []byte) []byte {
	base64Dec := make([]byte, base64.StdEncoding.EncodedLen(len(buffer)))
	base64.StdEncoding.Encode(base64Dec, buffer)
	return base64Dec
}
