package main

import (
	"fmt"
	"log"
	"os"
)

type argumentNError uint

func (e argumentNError) Error() string {
	return fmt.Sprintf("Error: wrong number of arguments: %v", uint(e))
}

// func mainHexToBase64() {
// 	if len(os.Args) != 2 {
// 		log.Fatal(argumentNError(len(os.Args) - 1))
// 	}
// 	hexString := os.Args[1]
// 	var target string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
// 	base64Dec, base64String, err := HexToBase64(hexString)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("base64 byte array %v\n", base64Dec)
// 	fmt.Printf("base64 string: %s\n", base64String)
// 	fmt.Printf("conversion successfull? %t\n", base64String == target)
// }

// func mainFixedXor() {
// 	if len(os.Args) != 3 {
// 		log.Fatal(argumentNError(len(os.Args) - 1))
// 	}

// 	first := os.Args[1]
// 	second := os.Args[2]
// 	var target string = "746865206b696420646f6e277420706c6179"
// 	xor, xorString, err := FixedXorStr(first, second)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("xored bytes: %v\n", xor)
// 	fmt.Printf("xor string: %s\n", xorString)
// 	fmt.Printf("xor successfull? %t\n", target == xorString)
// }

func main() {
	if len(os.Args) != 2 {
		log.Fatal(argumentNError(len(os.Args) - 1))
	}

	hexString := os.Args[1]
	cipher, key, _, err := Decode(hexString)

	if err != nil {
		panic(err)
	}

	fmt.Printf("byte cypher: %v\n", cipher)
	fmt.Printf("byte key: %b\n", key)
	fmt.Printf("encoded cipher: %s\n", string(cipher))
	fmt.Printf("encoded key: %s\n", string([]byte{key}))
}
