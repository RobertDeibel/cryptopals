package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

type argumentNError uint

var target []byte = []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")

func (e argumentNError) Error() string {
	return fmt.Sprintf("Error: wrong number of arguments: %v", uint(e))
}

func main() {
	args := os.Args[:]

	if len(args) != 2 {
		log.Fatal(argumentNError(len(args)))
		return
	}
	hexDec, err := hex.DecodeString(args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("len hexDec: %d len string: %d len target: %d\n", len(hexDec), len(args[1]), len(target))
	fmt.Printf("hex1: %v\n", hexDec)
	var base64 []byte
	for i := 0; i < (len(hexDec)-1)/3; i++ {
		j := i * 3
		k := i * 4
		fourByte := hexDec[j : j+3]
		base64[k] = (fourByte[0] & (128 + 64 + 32 + 16 + 8 + 4)) >> 2
		base64[k+1] = ((fourByte[0] & (2 + 1)) << 4) + ((fourByte[1] & (128 + 64 + 32 + 16)) >> 4)
		base64[k+2] = ((fourByte[1] & (8 + 4 + 2 + 1)) << 2) + ((fourByte[2] & (128 + 64)) >> 6)
		base64[k+3] = (fourByte[2] & (32 + 16 + 8 + 4 + 2 + 1))
	}
	fmt.Printf("base64Dec: %v", base64)

	fmt.Println("Hex stream: ")
}
