package main

import (
	"github.com/NeiFi/novobanco"
	"github.com/NeiFi/revolut"
	"github.com/NeiFi/wise"
)

// Naming for variables: name_of_variable
// Naming for types: nameOfType

type TransactionRecord interface {
	Print()
}

func main() {

	w_data := wise.GetData()
	r_data := revolut.GetData()
	n_data := novobanco.GetData()

	w_data.Print()
	r_data.Print()
	n_data.Print()

}
