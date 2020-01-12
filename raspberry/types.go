package main

import "github.com/eoscanada/eos-go"

type HiActionReq struct {
	Name     eos.AccountName `json:"name"`
}

type AddClaimEventActionReq struct {
	CreatedAt 	uint64
	GpsLt 		int32
	GpsLa 		int32
	VinInt      *eos.Uint128
	Vin 		[]byte
}

type CleanClaimEventActionReq struct {
	Name     eos.AccountName `json:"name"`
}

