package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type TransactionRecord struct {
	ID                string          `json:"ID"`
	Status            string          `json:"status"`
	Direction         string          `json:"direction"`
	Created_on        time.Time       `json:"created_on"`
	Finished_on       time.Time       `json:"finished_on"`
	Source_fee_amount decimal.Decimal `json:"source_fee_amount"`
	Source_fee_curr   string          `json:"source_fee_curr"`
	Target_fee_amount decimal.Decimal `json:"target_fee_amount"`
	Target_fee_curr   string          `json:"target_fee_curr"`
	Source_name       string          `json:"source_name"`
	Source_amount_af  decimal.Decimal `json:"source_amount_af"`
	Source_curr       string          `json:"source_curr"`
	Target_name       string          `json:"target_name"`
	Target_amount_af  decimal.Decimal `json:"target_amount_af"`
	Target_curr       string          `json:"target_curr"`
	Exchange_rate     float32         `json:"exchange_rate"`
	Reference         string          `json:"reference"`
	Batch             string          `json:"batch"`
}

const date_format_wise = "2006-01-02 15:04:05"
 
func main() {
	// open file
	f, err := os.Open("transaction-history-MAY-2023.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	csvReader.Read()
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var transactionList []TransactionRecord
	for _, row := range data {
		created_on, _ := time.Parse(date_format_wise, row[3])
		finished_on, _ := time.Parse(date_format_wise, row[4])
		source_fee_amount, _ := decimal.NewFromString(row[5])
		target_fee_amount, _ := decimal.NewFromString(row[7])
		source_amount_af, _ := decimal.NewFromString(row[10])
		target_amount_af, _ := decimal.NewFromString(row[13])
		exchange_rate, _ := strconv.ParseFloat(row[15], 32)
		transactionRecord := TransactionRecord{
			ID:                row[0],
			Status:            row[1],
			Direction:         row[2],
			Created_on:        created_on,
			Finished_on:       finished_on,
			Source_fee_amount: source_fee_amount,
			Source_fee_curr:   row[6],
			Target_fee_amount: target_fee_amount,
			Target_fee_curr:   row[8],
			Source_name:       row[9],
			Source_amount_af:  source_amount_af,
			Source_curr:       row[11],
			Target_name:       row[12],
			Target_amount_af:  target_amount_af,
			Target_curr:       row[14],
			Exchange_rate:     float32(exchange_rate),
			Reference:         row[16],
			Batch:             row[17],
		}
		transactionList = append(transactionList, transactionRecord)
	}
	// just to pretty print, to see data clearly
	transactionJSON, err := json.MarshalIndent(transactionList[4], "", " ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(transactionJSON))
}
