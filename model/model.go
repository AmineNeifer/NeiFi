package model

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Transactions interface {
	Print()
}

type Record interface {
	ToModel() Model
}

type Model struct {
	ID     string    `json:"id"`
	Date   time.Time `json:"date"`
	Type   string    `json:"type"`
	Debit  float32   `json:"debit"`
	Credit float32   `json:"credit"`
	Desc   string    `json:"desc"`
	Note   string    `json:"note"`
	Source string    `json:"source"`
	Target string    `json:"target"`
}

func CreateID(r Record) string {
	concatenated_data := fmt.Sprintf("%v", r)
	return encrypt(concatenated_data)
}

func encrypt(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	id := fmt.Sprintf("%x", h.Sum(nil))
	return id
}
