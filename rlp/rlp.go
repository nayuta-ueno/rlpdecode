package rlp

import (
	"encoding/hex"
	"fmt"
)

const indentStr = "  "

// Decode []byte to fmt.Print
// https://github.com/ethereum/wiki/wiki/%5BJapanese%5D-RLP
func Decode(data []byte) {
	decode(data, "")
}

func decode(data []byte, indent string) {
	cnt := 0
	for {
		if cnt >= len(data) {
			break
		}
		var dec []byte
		var dataLen int
		show := true
		switch {
		case data[cnt] < 0x80:
			// single data
			dec = make([]byte, 1)
			dec[0] = data[cnt]
			dataLen = 1
		case data[cnt] < 0xb8:
			// < 55 bytes data
			dataLen = int(data[cnt] - 0x80)
			cnt++
			dec = make([]byte, dataLen)
			copy(dec, data[cnt:])
		case data[cnt] < 0xc0:
			// >= 55 bytes data
			lenLen := int(data[cnt] - 0xb7)
			cnt++
			dda := make([]byte, lenLen)
			copy(dda, data[cnt:])
			cnt += lenLen
			dataLen = bigEndian(dda)
			dec = make([]byte, dataLen)
			copy(dec, data[cnt:])
		case data[cnt] < 0xf7:
			// < 55 bytes list
			dataLen = int(data[cnt] - 0xc0)
			cnt++
			lst := make([]byte, dataLen)
			copy(lst, data[cnt:])
			fmt.Printf("%s[%d]LIST\n", indent, dataLen)
			decode(lst, "  "+indent)
			show = false
		default:
			// >= 55 bytes list
			lenLen := int(data[cnt] - 0xf7)
			cnt++
			dda := make([]byte, lenLen)
			copy(dda, data[cnt:])
			cnt += lenLen
			dataLen = bigEndian(dda)
			lst := make([]byte, dataLen)
			copy(lst, data[cnt:])
			fmt.Printf("%s[%d]LIST\n", indent, dataLen)
			decode(lst, "  "+indent)
			show = false
		}
		if show {
			fmt.Printf("%s[%d] %s\n", indent, dataLen, hex.EncodeToString(dec))
		}
		cnt += dataLen
	}
}

func bigEndian(data []byte) int {
	var val int
	for _, i := range data {
		val <<= 8
		val += int(i)
	}
	return val
}
