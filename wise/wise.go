package wise

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"model"
	"os"
	"strconv"
	"strings"
	"time"
)

type TransactionRecordWise struct {
	ID                    string    `json:"id"`
	Date                  time.Time `json:"date"`
	Amount                float32   `json:"amount"`
	Currency              string    `json:"currency"`
	Description           string    `json:"description"`
	Payment_Reference     string    `json:"payment_reference"`
	Running_Balance       float32   `json:"running_balance"`
	Exchange_From         string    `json:"exchange_from"`
	Exchange_To           string    `json:"exchange_to"`
	Exchange_Rate         float32   `json:"exchange_rate"`
	Payer_Name            string    `json:"payer_name"`
	Payee_Name            string    `json:"payee_name"`
	Payee_Account_Number  string    `json:"payee_account_number"`
	Merchant              string    `json:"merchant"`
	Card_Last_Four_Digits uint16    `json:"card_last_four_digits"`
	Card_Holder_Full_Name string    `json:"card_holder_full_name"`
	Attachment            string    `json:"attachment"`
	Note                  string    `json:"note"`
	Total_fees            float32   `json:"total_fees"`
}

const date_format_wise = "02-01-2006"

// dd-mm-yyyy

func GetData() TransactionsWise {
	// open file
	f, err := os.Open("wise/transaction-history-MAY-2023-Wise.csv")
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

	var transaction_list_wise TransactionsWise
	for _, row := range data {
		// date conversions
		date, _ := time.Parse(date_format_wise, row[1])
		// float32 conversions
		amount, _ := strconv.ParseFloat(row[2], 32)
		running_balance, _ := strconv.ParseFloat(row[6], 32)
		exchange_rate, _ := strconv.ParseFloat(row[9], 32)
		total_fees, _ := strconv.ParseFloat(row[18], 32)
		// uint16 conversions
		card_last_four_digits, _ := strconv.ParseUint(row[14], 10, 16)
		transaction_record_wise := TransactionRecordWise{
			ID:                    row[0],
			Date:                  date,
			Amount:                float32(amount),
			Currency:              row[3],
			Description:           row[4],
			Payment_Reference:     row[5],
			Running_Balance:       float32(running_balance),
			Exchange_From:         row[7],
			Exchange_To:           row[8],
			Exchange_Rate:         float32(exchange_rate),
			Payer_Name:            row[10],
			Payee_Name:            row[11],
			Payee_Account_Number:  row[12],
			Merchant:              row[13],
			Card_Last_Four_Digits: uint16(card_last_four_digits),
			Card_Holder_Full_Name: row[15],
			Attachment:            row[16],
			Note:                  row[17],
			Total_fees:            float32(total_fees),
		}
		transaction_list_wise.records = append(transaction_list_wise.records, transaction_record_wise)
	}
	return transaction_list_wise
}

type TransactionsWise struct {
	records []TransactionRecordWise
}

func (r TransactionsWise) Print() {
	fmt.Println("Wise Example")
	// random index to print
	random_index := rand.Intn(len(r.records))

	transactionJSON, err := json.MarshalIndent(r.records[random_index], "", " ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(transactionJSON))
}

func (r TransactionRecordWise) ToModel() model.Model {
	var debit, credit float32 = 0, 0
	var source, target string

	concatenated_data := fmt.Sprintf("%v", r)
	type_ := strings.Split(r.ID, "-")[0]
	if r.Amount >= 0 {
		credit = r.Amount
	} else {
		debit = r.Amount
	}

	if len(r.Payer_Name) > 0 {
		source = r.Payer_Name
	} else {
		source = "Amine Neifer"
	}

	if len(r.Payee_Name) > 0 {
		target = r.Payee_Name
	} else if len(r.Merchant) > 0 {
		target = r.Merchant
	} else {
		target = "Amine Neifer"
	}

	m := model.Model{
		ID:     model.CreateID(concatenated_data),
		Date:   r.Date,
		Type:   type_,
		Debit:  debit,
		Credit: credit,
		Desc:   r.Description,
		Note:   r.Note,
		Source: source,
		Target: target,
	}
	return m
}
