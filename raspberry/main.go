package main

import (
	"encoding/hex"
	"fmt"
	"github.com/eoscanada/eos-go"
	"github.com/stianeikeland/go-rpio"
	"math/rand"
	"strconv"
	"time"
)

const pin = rpio.Pin(18)
const pinLed = rpio.Pin(23)

const (
	API_PROVIDER = "http://jungle2.cryptolions.io"
)

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	defer rpio.Close()

	for ; ; {
		bumpDetector()
	}
}

func bumpDetector() {
	pin.Input()
	pin.PullDown()
	pin.Detect(rpio.FallEdge)
	pinLed.Output()
	pinLed.PullDown()

	fmt.Println("Booting....")
	pinLed.Toggle()
	time.Sleep(1 * time.Second)
	pinLed.Toggle()

	fmt.Println("Monitoring accident")

	for i := 0; i < 1; {
		if pin.EdgeDetected() {
			fmt.Println("ACCIDENT OCCURRED")
			for x := 0; x < 10; x++ {
				pinLed.Toggle()
				time.Sleep(100 * time.Millisecond)
			}
			api := eos.New(API_PROVIDER)
			senderAcc := "myacc1111111"
			senderPk := "5J4yzEL3mByMXeDUwkorUauF1y8L9S8W8H8E1NT2BQr55UNcmpt"
			contractName := "claim1111112"
			vin := getVIN()

			req := &AddClaimEventActionReq{}
			req.CreatedAt = uint64(time.Now().Unix())
			req.Vin = []byte(vin)
			req.VinInt = VinToUint128(vin)
			req.GpsLa, req.GpsLt = PositionToInt(50.454800, 30.481412)
			req.GpsLa= req.GpsLa + int32(rand.Intn(100))
			req.GpsLt= req.GpsLt - int32(rand.Intn(70))

			broadcastResult, err := SignAndBroadcastTx(contractName, "addevent", senderPk, senderAcc, req, api)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(broadcastResult.Processed.ID))
			for x := 0; x < 10; x++ {
				pinLed.Toggle()
				time.Sleep(100 * time.Millisecond)
			}
			i++
		}
		time.Sleep(500 * time.Millisecond)
	}
	pin.Detect(rpio.NoEdge)
}

func getVIN() string {
	return "4T1CE30P17U75530" + strconv.Itoa(rand.Intn(9))
}