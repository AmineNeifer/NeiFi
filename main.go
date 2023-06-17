package main

import (
	"github.com/NeiFi/novobanco"
	"github.com/NeiFi/revolut"
	"github.com/NeiFi/wise"
)

func main() {
	revolut.PrintRevolut()
	wise.PrintWise()
	novobanco.PrintNovoBanco()
}
