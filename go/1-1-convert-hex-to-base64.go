package main

import (
  "encoding/hex"
  "encoding/base64"
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
  fmt.Println(n)
	fmt.Printf("%s\n", rawBytes[:n])

	str := base64.StdEncoding.EncodeToString(rawBytes)
	fmt.Println(str)
}
