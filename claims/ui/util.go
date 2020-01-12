package main

import (
	"bytes"
	"encoding/binary"
	"github.com/eoscanada/eos-go"
)

const (
	positionDec = 1000000
)

func VinToUint128(vin string) *eos.Uint128 {
	lo := []byte(vin[1:9])
	hi := []byte(vin[9:17])
	return &eos.Uint128{BytesToUInt64(lo), BytesToUInt64(hi)}
}

func BytesToUInt64(buffer []byte) (ret uint64) {
	buf := bytes.NewBuffer(buffer)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}
func PositionToInt(x, y float32) (int32, int32) {
	return int32(x * positionDec), int32(y * positionDec)
}

func IntToPosition(x, y int32) (float32, float32) {
	return float32(x) / positionDec, float32(y) / positionDec
}
