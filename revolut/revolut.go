package revolut

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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
func GetData() TransactionsRevolut {
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

	var transactionListRevolut TransactionsRevolut
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
		transactionListRevolut.records = append(transactionListRevolut.records, transactionRecordRevolut)
	}
	return transactionListRevolut
}

type TransactionsRevolut struct {
	records []TransactionRecordRevolut
}

func (r TransactionsRevolut) Print() {
	fmt.Println("Revolut Example")
	// random index to print
	random_index := rand.Intn(len(r.records))

	transactionJSON, err := json.MarshalIndent(r.records[random_index], "", " ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(transactionJSON))
}
