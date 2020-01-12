package main

import (
	"encoding/hex"
	"fmt"
	"github.com/eoscanada/eos-go"
	"testing"
	"time"
)

func TestAddHiAction (t *testing.T){
	api := eos.New(API_PROVIDER)
	senderAcc := "myacc1111111"
	senderPk := "5J4yzEL3mByMXeDUwkorUauF1y8L9S8W8H8E1NT2BQr55UNcmpt"
	contractName := "claim1111112"
	req := &HiActionReq{eos.AN("myacc1111112")}
	broadcastResult, err := SignAndBroadcastTx(contractName, "hi", senderPk, senderAcc, req, api)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(broadcastResult.Processed.ID))
}

func TestAddClaimEventAction (t *testing.T){
	api := eos.New(API_PROVIDER)
	senderAcc := "myacc1111111"
	senderPk := "5J4yzEL3mByMXeDUwkorUauF1y8L9S8W8H8E1NT2BQr55UNcmpt"
	contractName := "claim1111112"
	vin := "1T1CE30P17U755343"

	req := &AddClaimEventActionReq{}
	req.CreatedAt = uint64(time.Now().Unix())
	req.Vin = []byte(vin)
	req.VinInt = VinToUint128(vin)
	req.GpsLa,req.GpsLt = PositionToInt(50.454213, 30.367816)
	broadcastResult, err := SignAndBroadcastTx(contractName, "addevent", senderPk, senderAcc, req, api)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(broadcastResult.Processed.ID))
}

//only self (contract can clean table)
func TestCleanClaimEventTbl (t *testing.T){
	api := eos.New(API_PROVIDER)
	senderAcc := "claim1111112"
	senderPk := "5KgRMDx9D4EXMSmy9vZhkmzN7qmZHkoUDgCkvg29s7TDpbiD2oj"
	contractName := "claim1111112"

	req := &CleanClaimEventActionReq{}
	req.Name = eos.AN(senderAcc)
	broadcastResult, err := SignAndBroadcastTx(contractName, "delevents", senderPk, senderAcc, req, api)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(broadcastResult.Processed.ID))
}

