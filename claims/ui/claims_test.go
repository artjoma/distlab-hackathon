package main

import (
	"encoding/json"
	"fmt"
	"github.com/eoscanada/eos-go"
	"testing"
)

/*
http://jungle2.cryptolions.io:80
https://jungle2.cryptolions.io:443
https://api.jungle.alohaeos.com:443
http://145.239.133.201:8888
http://jungle.eoscafeblock.com:8888
http://jungle-2.eosgen.io:80
http://jungle2.eosdac.io:8882
https://jungle.eosn.io:443
http://jungle.eosbcn.com:8080
http://jungle.atticlab.net:8888
https://jungleapi.eossweden.se:443
https://jungle.eosdac.io:443
https://jungle.eosphere.io:443
http://jungle2.cryptolions.io:8888
http://12.185.120.20:8888
http://jungle-2.eosgen.io:80

*/

func TestGetInfo(t *testing.T) {
	api := eos.New(API_PROVIDER)

	info, err := api.GetInfo()
	if err != nil {
		panic(fmt.Errorf("get info: %s", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %s", err))
	}

	fmt.Println(string(bytes))
}

//simple test get table
func TestGetTableInfo(t *testing.T) {
	api := eos.New(API_PROVIDER)
	tblRequest := eos.GetTableRowsRequest{}
	tblRequest.Code = "claim1111112"
	tblRequest.Scope = "claim1111112"
	tblRequest.Table = "claimev"
	tblRequest.JSON = true
	response, err := api.GetTableRows(tblRequest)
	if err != nil {
		panic(err)
	}

	data, err := response.Rows.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}

func TestGetTableSimple(t *testing.T) {
	api := eos.New(API_PROVIDER)
	tblRequest := eos.GetTableRowsRequest{}
	tblRequest.Code = "claim1111112"
	tblRequest.Scope = "claim1111112"
	tblRequest.Table = "claimev"
	tblRequest.JSON = true
	response, err := api.GetTableRows(tblRequest)
	if err != nil {
		panic(err)
	}

	data, err := response.Rows.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

//range
func TestGetTableInfoByRange(t *testing.T) {
	api := eos.New(API_PROVIDER)
	tblRequest := eos.GetTableRowsRequest{}
	tblRequest.Code = "claim1111112"
	tblRequest.Scope = "claim1111112"
	tblRequest.Table = "claimev"
	tblRequest.JSON = true
	tblRequest.LowerBound = "2"
	tblRequest.UpperBound = "-1"
	response, err := api.GetTableRows(tblRequest)
	if err != nil {
		panic(err)
	}

	data, err := response.Rows.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println("more:", response.More)
	fmt.Println(string(data))
}

func TestGetById(t *testing.T) {
	api := eos.New(API_PROVIDER)
	tblRequest := eos.GetTableRowsRequest{}
	tblRequest.Code = "claim1111112"
	tblRequest.Scope = "claim1111112"
	tblRequest.Table = "event"
	tblRequest.JSON = true
	tblRequest.Index = "1" //primary index
	tblRequest.LowerBound = "1"
	tblRequest.UpperBound = "1"
	response, err := api.GetTableRows(tblRequest)
	if err != nil {
		panic(err)
	}

	data, err := response.Rows.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println("more:", response.More)
	fmt.Println(string(data))
}

func TestGetByVin(t *testing.T) {
	vin := "4T1CE30P17U755309"
	api := eos.New(API_PROVIDER)
	tblRequest := eos.GetTableRowsRequest{}
	tblRequest.Code = "claim1111112"
	tblRequest.Scope = "claim1111112"
	tblRequest.Table = "claimev"
	tblRequest.JSON = true
	tblRequest.Index = "2" //primary index
	tblRequest.KeyType = "i128"
	tblRequest.LowerBound = VinToUint128(vin).String()
	tblRequest.UpperBound = VinToUint128(vin).String()

	response, err := api.GetTableRows(tblRequest)
	if err != nil {
		panic(err)
	}

	data, err := response.Rows.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println("more:", response.More)
	fmt.Println(string(data))
	//340282366920938463463374607431768211455
	//111910938596029940661301829439556104761
}
