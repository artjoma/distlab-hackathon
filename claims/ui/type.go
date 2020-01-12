package main

import (
	"fmt"
	"math/big"
	"time"
)

type AccidentEvent struct {
	Id        *big.Int `json:"id"`
	Vin       string   `json:"vin"`
	Lat       int32    `json:"gpsLt"`
	Lon       int32    `json:"gpsLa"`
	Timestamp int64    `json:"createdAt"`
}

func (m *AccidentEvent) String() string {
	x, y := IntToPosition(m.Lat, m.Lon)
	ts := time.Unix(m.Timestamp, 0).Format("2006-02-01 15:04:05")
	return fmt.Sprintf("Acident: VIN: %s @%f,%f on %s\n", m.Vin, x, y, ts)
}

type ClaimDetails struct {
	Id            *big.Int
	Country       string
	City          string
	Street        string
	HouseNumber   uint16
	AccidentTime  uint64
	OfficerId     uint64
	External      []byte
	FormatVersion uint16
	Version       uint16
	UpdatedAt     uint64
	UpdatedBy     uint64
	CreatedAt     uint64
	CreatedBy     uint64
}
