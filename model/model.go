package model

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Model struct {
	ID     string    `json:"id"`
	Date   time.Time `json:"date"`
	Type   string    `json:"type"`
	Debit  float32   `json:"debit"`
	Credit float32   `json:"credit"`
	Desc   string    `json:"desc"`
	Source string    `json:"source"`
}

func createID(concatenated_data string) string {
	return encrypt(concatenated_data)
}

func encrypt(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	id := fmt.Sprintf("%x", h.Sum(nil))
	return id
}
