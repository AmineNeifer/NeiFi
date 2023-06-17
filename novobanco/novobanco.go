package novobanco

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type TransactionRecordNovoBanco struct {
	Date_of_transaction time.Time `json:"date_of_transaction"`
	Value_date          time.Time `json:"value_date"`
	Type                string    `json:"type"`
	Description         string    `json:"description"`
	Debit               float32   `json:"debit"`
	Credit              float32   `json:"credit"`
	Balance             float32   `json:"balance"`
}

const date_format_novobanco = "02-01-2006"
// dd-mm-yyyy

func PrintNovoBanco() {

	// open file
	f, err := os.Open("novobanco/transaction-history-MAY-2023-Novo.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	// csvReader.Read()     we don't use this with Novobanco because of the different csv
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var transactionListNovoBanco []TransactionRecordNovoBanco
	for i, row := range data {
		if i < 10 {
			continue
		}
		date_of_transaction, _ := time.Parse(date_format_novobanco, row[0])
		value_date, _ := time.Parse(date_format_novobanco, row[1])
		debit, _ := strconv.ParseFloat(row[4], 32)
		credit, _ := strconv.ParseFloat(row[5], 32)
		balance, _ := strconv.ParseFloat(row[6], 32)

		transactionRecordNovoBanco := TransactionRecordNovoBanco{
			Date_of_transaction: date_of_transaction,
			Value_date:          value_date,
			Type:                row[2],
			Description:         row[3],
			Debit:               float32(debit),
			Credit:              float32(credit),
			Balance:             float32(balance),
		}
		transactionListNovoBanco = append(transactionListNovoBanco, transactionRecordNovoBanco)
	}
	// just to pretty print, to see data clearly
	fmt.Println("NovoBanco Example")
	transactionJSON, err := json.MarshalIndent(transactionListNovoBanco[3], "", " ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(transactionJSON))
}
