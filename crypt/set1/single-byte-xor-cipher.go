package crypt

import (
	"encoding/hex"
	"sort"
)

func score(hexBuffer []byte) int {
	scoring := make(map[byte]int)
	scoring['e'] = 12
	scoring['E'] = 12
	scoring['t'] = 11
	scoring['T'] = 11
	scoring['a'] = 10
	scoring['A'] = 10
	scoring['o'] = 9
	scoring['O'] = 9
	scoring['i'] = 8
	scoring['I'] = 8
	scoring['n'] = 7
	scoring['N'] = 7
	scoring[' '] = 6
	scoring['s'] = 5
	scoring['S'] = 5
	scoring['h'] = 4
	scoring['H'] = 4
	scoring['r'] = 3
	scoring['R'] = 3
	scoring['D'] = 2
	scoring['d'] = 2
	scoring['l'] = 1
	scoring['L'] = 1
	scoring['U'] = 1
	scoring['u'] = 1

	score := 0
	for i := range hexBuffer {
		add, ok := scoring[hexBuffer[i]]
		if ok {
			score += add
		}
	}

	return score
}

// DecodeSingleXorString attempts to decode a hex encoded string XORed with a single byte
// by testing each possible byte and scoring based on letter frequency: 'ETAOIN SHRDLU'
// returns decoded byte array, key as bytes and the score of the best solution
func DecodeSingleXorString(hexString string) ([]byte, byte, int, error) {

	decodedHex, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, 0, 0, err
	}

	return DecodeSingleXor(decodedHex)
}

// DecodeSingleXor attempts to decode a byte array XORed with a single byte
// by testing each possible byte and scoring based on letter frequency: 'ETAOIN SHRDLU'
// returns decoded byte array, key as bytes and the score of the best solution
func DecodeSingleXor(input []byte) ([]byte, byte, int, error) {
	type keyScore struct {
		Key   byte
		Score int
	}

	results := make([]keyScore, 256)
	for key := 0; key < 256; key++ {
		keyBuffer := make([]byte, len(input))
		for i := range keyBuffer {
			keyBuffer[i] = byte(key)
		}
		buf, err := FixedXor(input, keyBuffer)
		if err != nil {
			return nil, 0, 0, err
		}
		score := score(buf)
		results[key] = keyScore{byte(key), score}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	best := results[0]
	bestBuffer := make([]byte, len(input))
	for i := range input {
		bestBuffer[i] = input[i] ^ best.Key
	}

	return bestBuffer, best.Key, best.Score, nil
}
