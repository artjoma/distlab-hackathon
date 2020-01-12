package main

import (
	"fmt"
	"testing"
)
//111910938596029940661301829439556104761 / 54314345333050313755373535333239
//111910938596029940661301829439556104752 / 54314345333050313755373535333230
func TestVinToNum (t *testing.T){
	vin := "4T1CE30P17U755329"
	num := VinToUint128(vin)
	fmt.Println(num.String())
	//fmt.Println(num.Text(16))
	//fmt.Println(string(num.Bytes()))

	vin = "4T1CE30P17U755320"
	num = VinToUint128(vin)
	fmt.Println(num.String())
	//fmt.Println(num.Text(16))
	//fmt.Println(string(num.Bytes()))
}
//111910938596029940661301829439556104761
//340282366920938463463374607431768211455
func TestPosition (t *testing.T){
	fmt.Println(PositionToInt(10.324342, -7.123456))
	//10324342 -7123456
	fmt.Println(IntToPosition(10324342, -7123456))
}