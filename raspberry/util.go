package main

import (
	"bytes"
	"encoding/binary"
	"github.com/eoscanada/eos-go"
)

const(
	positionDec = 1000000
)

func VinToUint128 (vin string) *eos.Uint128{
	lo := []byte(vin[1:9])
	hi := []byte(vin[9:17])
	return &eos.Uint128{BytesToUInt64(lo), BytesToUInt64(hi)}
}

func BytesToUInt64(buffer []byte) (ret uint64) {
	buf := bytes.NewBuffer(buffer)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func PositionToInt (x, y float32)(int32, int32){
	return int32(x * positionDec), int32(y * positionDec)
}

func IntToPosition (x, y int32)(float32, float32){
	return float32(x) / positionDec, float32(y) / positionDec
}

func SignAndBroadcastTx (contractAccStr, actionNameStr, senderPk, senderAccStr string, req interface{},
		api *eos.API) (*eos.PushTransactionFullResp, error){
	keyBag := &eos.KeyBag{}
	senderAcc := eos.AN(senderAccStr)
	err := keyBag.ImportPrivateKey(senderPk)
	if err != nil {
		return nil, err
	}
	api.SetSigner(keyBag)

	if err != nil {
		return nil, err
	}

	txOpts := &eos.TxOptions{}
	if err := txOpts.FillFromChain(api); err != nil {
		return nil, err
	}

	authorization := []eos.PermissionLevel{{Actor: senderAcc, Permission: eos.PN("active")}}
	actList := []*eos.Action{&eos.Action{eos.AN(contractAccStr), eos.ActN(actionNameStr), authorization, eos.NewActionData(req)}}

	tx := eos.NewTransaction(actList, txOpts)
	_, packedTx, err := api.SignTransaction(tx, txOpts.ChainID, eos.CompressionZlib)
	if err != nil {
		return nil, err
	}

	response, err := api.PushTransaction(packedTx)
	if err != nil {
		return nil, err
	}

	return response, nil;
}