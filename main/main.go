package main

import (
	"fmt"
	//goTezos "github.com/DefinitelyNotAGoat/go-tezos"
)

func main() {
	fmt.Println("hello")

	

	pkh := "tz1Nhj1wHs7nzHSwdybxrYjpEQCTaEpWwu6w"
	pk := "edpkudUXnqaKAFTdzDkiWareV7B7QVPZU9R3wrjqgp55FVStgTj7DH"
	sks := "edsk4VqLBCE5ETiFNWBb27tMVQc2fwjYLEDADuBqCQYbTyh417tye8"

	myWallet, err := gt.Account.ImportWallet(pkh, pk, sks)

	if err != nil {
		fmt.Printf("failed to import wallet for %v\n", pkh)
	}
}
