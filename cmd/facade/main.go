package main

import (
	"fmt"
	"log"
	"reworked/pkg/facade"
)

func main() {
	fmt.Println("-----------")
	example := facade.NewCustomerFacade("Johnt", "passtest")
	fmt.Println("---------------------")
	err := example.AddMoneyToWallet("Johnt", "passtest", 150)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	err = example.SubtractMoneyFromWallet("Johnt", "passtest", 50)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
