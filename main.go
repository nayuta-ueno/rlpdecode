package main

import (
	"encoding/hex"
	"rlpdecode/rlp"
)

func main() {
	data, _ := hex.DecodeString("3830613637386634393539383439643539343530653336616239643433383131383061363738663439353938343964353934353065333661623964343338313133663061336231363432306134316363393865366237393834636436646237383262613436646665623463303439333239396364356639336261616163613539000000000000000000000000000000000000000000000000000000000000138800000000000000000000000000000000000000000000000000000000000001206b696e64000000000000000000000000000000000000000000000000000000006e6f6e65000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000000f7b226b696e64223a226e6f6e65227d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000103237323263353364343831663033356200000000000000000000000000000000")
	rlp.Decode(data)
}
