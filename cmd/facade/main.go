package main

import (
	"fmt"
	"log"

	"github.com/saharanomad-wb/learning/pkg/facade"
)

func main() {
	fmt.Println("Wallet service started")

	f := facade.CreateWalletSystem()

	fmt.Scanln()

	for i := 0; i < 5; i++ {
		var b int
		var name string

		fmt.Printf("Enter you name: ")
		fmt.Scan(&name)
		fmt.Printf("amount of money: ")
		fmt.Scan(&b)

		fmt.Println("You entered ", name, b)

		_, err := f.CreateNewWallet(&b, name)
		if err != nil {
			log.Fatal(err)
		}

		f.ShowFacadeSustem()
	}

}
