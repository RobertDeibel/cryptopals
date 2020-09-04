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

// Decode attempts to decode a hex encoded string with letter frequency analysis
// returns decoded byte array, and key in base 64
func Decode(hexString string) ([]byte, byte, int, error) {
	type keyScore struct {
		Key   byte
		Score int
	}

	decodedHex, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, 0, 0, err
	}

	results := make([]keyScore, 256)
	for key := 0; key < 256; key++ {
		keyBuffer := make([]byte, len(decodedHex))
		for i := range keyBuffer {
			keyBuffer[i] = byte(key)
		}
		buf := FixedXor(decodedHex, keyBuffer)
		score := score(buf)
		results[key] = keyScore{byte(key), score}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	best := results[0]
	bestBuffer := make([]byte, len(decodedHex))
	for i := range decodedHex {
		bestBuffer[i] = decodedHex[i] ^ best.Key
	}

	return bestBuffer, best.Key, best.Score, nil
}
