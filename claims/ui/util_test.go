package main

import (
	"fmt"
	"testing"
)

func TestVinToNum(t *testing.T) {
	vin := "4T1CE30P17U755329"
	num := VinToUint128(vin)
	fmt.Println(num.String())
}

func TestPosition(t *testing.T) {
	fmt.Println(PositionToInt(10.324342, -7.123456))
	//10324342 -7123456
	fmt.Println(IntToPosition(10324342, -7123456))
}
