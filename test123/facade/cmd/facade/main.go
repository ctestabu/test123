/*cmd/
All the programs this project owns belongs inside the cmd/ folder.
The folders under cmd/ are always named for each program that will be built.
Use the letter d at the end of a program folder to denote it as a daemon.
Each folder has a matching source code file that contains the main package.
*/
package main

import (
	"fmt"
	"log"
	"facade"
	"history"
	"notification"
	"password"
	"product"
	"username"
	"wallet"
)

func main() {
	fmt.Println("-----------")
	example := newCustomerFacade("Johnt", "passtest")
	fmt.Println("---------------------")
	err := example.addMoneyToWallet("Johnt", "passtest", 150)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	err = example.subtractMoneyFromWallet("Johnt", "passtest", 50)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
