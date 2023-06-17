package revolut

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

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

const date_format_revolut = "1/2/2006 15:04"

// mm/dd/yyyy
func PrintRevolut() {
	// open file
	f, err := os.Open("revolut/transaction-history-MAY-2023-Rev.csv")
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
	transactionJSON, err := json.MarshalIndent(transactionListRevolut[1], "", " ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(transactionJSON))
}
