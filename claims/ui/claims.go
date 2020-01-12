package main

import (
	"encoding/json"
	"fmt"
	"github.com/eoscanada/eos-go"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"math/big"
	"strconv"
	"time"
)

var app = tview.NewApplication()
var searchForm = tview.NewForm()
var resultTable = tview.NewTable()
var resultTableTextView = tview.NewTextView()
var isSearchMode = false

const API_PROVIDER = "http://jungle2.cryptolions.io"

/**
go mod init claims
go mod vendor
*/
func main() {
	if err := buildUi().Run(); err != nil {
		panic(err)
	}
}

func buildDefaultSerchForm() {
	searchForm.
		AddInputField("VIN:", "", 17, nil, nil).
		AddButton("Search", srchHandler).
		AddButton("Reset", resetForm)
}

func resetForm() {
	searchForm.GetFormItemByLabel("VIN:").(*tview.InputField).SetText("")
	app.SetFocus(searchForm.GetFormItemByLabel("VIN:"))
	isSearchMode = false
}

func buildDefaultResultTable() {
	resultTable.Clear()
	resultTable.SetSelectable(false, false)
	resultTable.
		SetCellSimple(0, 0, "#").
		SetCellSimple(0, 1, "id").
		SetCellSimple(0, 2, "vin").
		SetCellSimple(0, 3, "lat.").
		SetCellSimple(0, 4, "lon.").
		SetCellSimple(0, 5, "time").
		SetBorders(true).
		SetFixed(100, 5).
		SetTitle("Search result")

	resultTableTextView.Clear()
}

func buildUi() *tview.Application {
	buildDefaultSerchForm()
	buildDefaultResultTable()
	go func() {
		for {
			time.Sleep(time.Second)
			if !isSearchMode {
				getLastEvents()
			}
		}
	}()

	results := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(resultTableTextView, 0, 1, false).
		AddItem(resultTable, 0, 16, false)

	layout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(searchForm, 0, 1, true).
		AddItem(results, 0, 4, false)

	app := app.SetRoot(layout, true)
	return app
}

func srchHandler() {
	vin := searchForm.GetFormItemByLabel("VIN:").(*tview.InputField).GetText()
	isSearchMode = true
	if len(vin) < 17 {
		resetForm()
		return
	}

	resultTableTextView.SetText(fmt.Sprintf("Looking for VIN: %s", vin))
	if blockchainResponse, err := getBlockchainRespose(vin); err == nil {
		buildAccidentEventView(blockchainResponse)
		app.SetFocus(resultTable)
	} else {
		panic(err)
	}
}

func buildAccidentEventView(blockchainResponse []*AccidentEvent) {
	buildDefaultResultTable()
	for i, data := range blockchainResponse {
		x, y := IntToPosition(data.Lat, data.Lon)
		resultTable.SetCellSimple(i+1, 0, strconv.Itoa(i+1))
		resultTable.SetCellSimple(i+1, 1, strconv.FormatUint(data.Id.Uint64(), 10))
		resultTable.SetCellSimple(i+1, 2, data.Vin)
		resultTable.SetCellSimple(i+1, 3, strconv.FormatFloat(float64(x), 'f', -1, 32))
		resultTable.SetCellSimple(i+1, 4, strconv.FormatFloat(float64(y), 'f', -1, 32))
		resultTable.SetCellSimple(i+1, 5, time.Unix(data.Timestamp, 0).Format("2006-02-01 15:04:05"))
		resultTable.SetSelectable(true, false)
		resultTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == 13 {
				r, _ := resultTable.GetSelection()
				if r != 0 {
					eventId := resultTable.GetCell(r, 1).Text
					handleGetDetails(eventId)
				}
			} else if event.Key() == 27 {
				app.SetFocus(searchForm)
			}
			return event
		})
		app.Draw()
	}
}

func getLastEvents() ([]*AccidentEvent, error) {
	api := eos.New(API_PROVIDER)
	tblRequest := eos.GetTableRowsRequest{}
	tblRequest.Code = "claim1111112"
	tblRequest.Scope = "claim1111112"
	tblRequest.Table = "claimev"
	tblRequest.JSON = true
	response, err := api.GetTableRows(tblRequest)
	if err != nil {
		return nil, err
	}

	data, err := response.Rows.MarshalJSON()
	if err != nil {
		return nil, err
	}

	model := make([]*AccidentEvent, 0)
	err = json.Unmarshal([]byte(data), &model)
	if err != nil {
		return nil, err
	}
	resultTableTextView.Clear()
	resultTable.Clear()
	buildDefaultResultTable()
	resultTableTextView.SetText("Last known events")
	buildAccidentEventView(model)

	return model, nil
}

func buildDetailsView() {

}

func getBlockchainRespose(vin string) ([]*AccidentEvent, error) {
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

	model := make([]*AccidentEvent, 0)
	err = json.Unmarshal([]byte(data), &model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func handleGetDetails(eventId string) {
	resultTableTextView.SetText(fmt.Sprintf("Details for claim: %s", eventId))

	resultTable.Clear()
	resultTable.
		SetCellSimple(0, 0, "id").
		SetCellSimple(0, 1, "country").
		SetCellSimple(0, 2, "city").
		SetCellSimple(0, 3, "street").
		SetCellSimple(0, 4, "number").
		SetCellSimple(0, 5, "time").
		SetCellSimple(0, 6, "officer")

	if detailsResponse, err := getBlockchainDetailsResponse(eventId); err == nil {
		resultTable.SetCellSimple(1, 0, strconv.FormatUint(detailsResponse.Id.Uint64(), 10))
		resultTable.SetCellSimple(1, 1, detailsResponse.Country)
		resultTable.SetCellSimple(1, 2, detailsResponse.City)
		resultTable.SetCellSimple(1, 3, detailsResponse.Street)
		resultTable.SetCellSimple(1, 4, strconv.FormatUint(uint64(detailsResponse.HouseNumber), 10))
		resultTable.SetCellSimple(1, 5, time.Unix(int64(detailsResponse.AccidentTime), 0).Format("2006-02-01 15:04:05"))
		resultTable.SetCellSimple(1, 6, strconv.FormatUint(detailsResponse.OfficerId, 10))
	}
}

func getBlockchainDetailsResponse(accidentId string) (*ClaimDetails, error) {
	id, _ := big.NewInt(0).SetString(accidentId, 0)
	return &ClaimDetails{
		Id:           id,
		Country:      "Ukraine",
		City:         "Kyiv",
		Street:       "Lesy Ukrayinky blvd.",
		HouseNumber:  10,
		AccidentTime: 1568470766,
		OfficerId:    123,
	}, nil
}
