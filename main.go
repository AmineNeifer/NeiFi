package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type TransactionRecordWise struct {
	ID                string    `json:"ID"`
	Status            string    `json:"status"`
	Direction         string    `json:"direction"`
	Created_on        time.Time `json:"created_on"`
	Finished_on       time.Time `json:"finished_on"`
	Source_fee_amount float32   `json:"source_fee_amount"`
	Source_fee_curr   string    `json:"source_fee_curr"`
	Target_fee_amount float32   `json:"target_fee_amount"`
	Target_fee_curr   string    `json:"target_fee_curr"`
	Source_name       string    `json:"source_name"`
	Source_amount_af  float32   `json:"source_amount_af"`
	Source_curr       string    `json:"source_curr"`
	Target_name       string    `json:"target_name"`
	Target_amount_af  float32   `json:"target_amount_af"`
	Target_curr       string    `json:"target_curr"`
	Exchange_rate     float32   `json:"exchange_rate"`
	Reference         string    `json:"reference"`
	Batch             string    `json:"batch"`
}

type TransactionRecordRevolut struct {
	Type           string    `json:"type"`
	Product        string    `json:"product"`
	Started_date   time.Time `json:"started_date"`
	Completed_date time.Time `json:"completed_date"`
	Description    string    `json:"description"`
	Amount         float32   `json:"amount"`
	Fee            float32   `json:"fee"`
	Currency       string    `json:"currency"`
	State          string    `json:"state"`
	Balance        float32   `json:"balance"`
}

const date_format_wise = "2006-01-02 15:04:05"

// yyyy-mm-dd
const date_format_revolut = "2/1/2006 15:04"

//dd/mm/yyyy

func main() {
	// open file
	f, err := os.Open("transaction-history-MAY-2023-Wise.csv")
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

	var transactionListWise []TransactionRecordWise
	for _, row := range data {
		created_on, _ := time.Parse(date_format_wise, row[3])
		finished_on, _ := time.Parse(date_format_wise, row[4])
		source_fee_amount, _ := strconv.ParseFloat(row[5], 32)
		target_fee_amount, _ := strconv.ParseFloat(row[7], 32)
		source_amount_af, _ := strconv.ParseFloat(row[10], 32)
		target_amount_af, _ := strconv.ParseFloat(row[13], 32)
		exchange_rate, _ := strconv.ParseFloat(row[15], 32)
		transactionRecordWise := TransactionRecordWise{
			ID:                row[0],
			Status:            row[1],
			Direction:         row[2],
			Created_on:        created_on,
			Finished_on:       finished_on,
			Source_fee_amount: float32(source_fee_amount),
			Source_fee_curr:   row[6],
			Target_fee_amount: float32(target_fee_amount),
			Target_fee_curr:   row[8],
			Source_name:       row[9],
			Source_amount_af:  float32(source_amount_af),
			Source_curr:       row[11],
			Target_name:       row[12],
			Target_amount_af:  float32(target_amount_af),
			Target_curr:       row[14],
			Exchange_rate:     float32(exchange_rate),
			Reference:         row[16],
			Batch:             row[17],
		}
		transactionListWise = append(transactionListWise, transactionRecordWise)
	}
	_ = transactionListWise
	// just to pretty print, to see data clearly
	fmt.Println("Wise Example")
	transactionJSON, err := json.MarshalIndent(transactionListWise[4], "", " ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(transactionJSON))

	// open file
	f, err = os.Open("transaction-history-MAY-2023-Rev.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader = csv.NewReader(f)
	csvReader.Read()
	data, err = csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var transactionListRevolut []TransactionRecordRevolut
	for _, row := range data {
		started_date, _ := time.Parse(date_format_revolut, row[2])
		completed_date, _ := time.Parse(date_format_revolut, row[3])
		amount, _ := strconv.ParseFloat(row[5], 32)
		fee, _ := strconv.ParseFloat(row[6], 32)
		balance, _ := strconv.ParseFloat(row[9], 32)

		transactionRecordRevolut := TransactionRecordRevolut{
			Type:           row[0],
			Product:        row[1],
			Started_date:   started_date,
			Completed_date: completed_date,
			Description:    row[4],
			Amount:         float32(amount),
			Fee:            float32(fee),
			Currency:       row[7],
			State:          row[8],
			Balance:        float32(balance),
		}
		transactionListRevolut = append(transactionListRevolut, transactionRecordRevolut)
	}
	// just to pretty print, to see data clearly
	fmt.Println("Revolut Example")
	transactionJSON, err = json.MarshalIndent(transactionListRevolut[1], "", " ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(transactionJSON))
}
