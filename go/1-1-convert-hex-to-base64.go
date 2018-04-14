package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	src := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	rawBytes := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(rawBytes, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rawBytes)
	// [73 39 109 32 107 105 108 108 105 110 103 32 121 111 117 114 32 98 114 97 105 110 32 108 105 107 101 32 97 32 112 111 105 115 111 110 111 117 115 32 109 117 115 104 114 111 111 109]

	fmt.Println(n)
	// 48

	fmt.Printf("%s\n", rawBytes[:n])
	// I'm killing your brain like a poisonous mushroom

	str := base64.StdEncoding.EncodeToString(rawBytes)

	fmt.Println(str) // SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
}
