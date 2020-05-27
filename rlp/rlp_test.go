package rlp

import (
	"encoding/hex"
	"testing"
)

func TestDecodeSingleString1(t *testing.T) {
	data, _ := hex.DecodeString("00")
	Decode(data)
}

func TestDecodeSingleString2(t *testing.T) {
	data, _ := hex.DecodeString("7f")
	Decode(data)
}

func TestDecodeMultiString1(t *testing.T) {
	data, _ := hex.DecodeString("80")
	Decode(data)
}

func TestDecodeMultiString2(t *testing.T) {
	data, _ := hex.DecodeString("a00000000000000000000000000000000000000000000000000000000000000000")
	Decode(data)
}

func TestDecodeMultiString3(t *testing.T) {
	data, _ := hex.DecodeString("b711223344556677889900112233445566778899001122334455667788990011223344556677889900112233445566778899001122334455")
	Decode(data)
}

func TestDecodeSingleSingle(t *testing.T) {
	data, _ := hex.DecodeString("007f")
	Decode(data)
}

func TestDecodeSingleMulti(t *testing.T) {
	data, _ := hex.DecodeString("00b711223344556677889900112233445566778899001122334455667788990011223344556677889900112233445566778899001122334455")
	Decode(data)
}

func TestDecodeMultiSingle(t *testing.T) {
	data, _ := hex.DecodeString("807f")
	Decode(data)
}

func TestDecodeMultiMulti(t *testing.T) {
	data, _ := hex.DecodeString("80b711223344556677889900112233445566778899001122334455667788990011223344556677889900112233445566778899001122334455")
	Decode(data)
}

func TestDecodeMultiLong1(t *testing.T) {
	data, _ := hex.DecodeString("b8381122334455667788990011223344556677889900112233445566778899001122334455667788990011223344556677889900112233445566")
	Decode(data)
}

func TestDecodeList1(t *testing.T) {
	data, _ := hex.DecodeString("f83ea00000000000000000000000000000000000000000000000000000000000000000d5946cc5fb56e43b3d023a6f43de915fab26c4c1ec96808400000000c0")
	Decode(data)
}
