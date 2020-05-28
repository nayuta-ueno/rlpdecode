package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"rlpdecode/rlp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%s <RLP string>\n", os.Args[0])
		os.Exit(1)
	}
	param := os.Args[1]
	if param[0:2] == "0x" {
		param = param[2:]
	}
	data, _ := hex.DecodeString(param)
	rlp.Decode(data)
}
