package crypt

import (
	"bufio"
	"encoding/hex"
	"log"
	"os"
	"sort"
)

// DetectSingleXorPhrase tries to detect and return a xored cipher from a file of cipers
// uses DetectSingleXor for scoring of the ciphers
// returns
func DetectSingleXorPhrase(file string) ([]byte, byte, int, int, error) {
	type clkv struct {
		Cipher []byte
		Line   int
		Key    byte
		Value  int
	}
	lines, err := ReadLinesFromFile(file)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	ranking := make([]clkv, len(lines))

	for i, line := range lines {
		decodeLine := make([]byte, hex.DecodedLen(len(line)))
		hex.Decode(decodeLine, line)
		cipher, key, score, err := DecodeSingleXor(decodeLine)
		if err != nil {
			return nil, 0, 0, 0, err
		}
		ranking[i] = clkv{cipher, i, key, score}
	}

	sort.Slice(ranking, func(i, j int) bool {
		return ranking[i].Value > ranking[j].Value
	})

	best := ranking[0]

	return best.Cipher, best.Key, best.Value, best.Line, err
}

// ReadLinesFromFile reads the file and returns the lines as fields of a array
func ReadLinesFromFile(file string) ([][]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	fileScanner := bufio.NewScanner(f)
	// destination of the read lines
	var lines [][]byte
	for fileScanner.Scan() {
		currentLine := fileScanner.Bytes()
		newLine := make([]byte, len(currentLine))
		copy(newLine, currentLine)
		lines = append(lines, newLine)
	}

	if err = fileScanner.Err(); err != nil {
		return lines, err
	}

	return lines, err
}
