package main

import (
	"fmt"
	//goTezos "github.com/DefinitelyNotAGoat/go-tezos"
	"github.com/fredcy/tztipbot"
	"os"
)

func main() {
	gt, err := tztipbot.ConnectService()
	if err != nil {
		fmt.Printf("cannot connect to RPC service")
		os.Exit(1)
	}

	// test account on alphanet
	pkh := "tz1Nhj1wHs7nzHSwdybxrYjpEQCTaEpWwu6w"
	pk := "edpkudUXnqaKAFTdzDkiWareV7B7QVPZU9R3wrjqgp55FVStgTj7DH"
	sks := "edsk4VqLBCE5ETiFNWBb27tMVQc2fwjYLEDADuBqCQYbTyh417tye8"

	myWallet, err := gt.Account.ImportWallet(pkh, pk, sks)

	if err != nil {
		fmt.Printf("failed to import wallet for %v\n", pkh)
	}

	fmt.Printf("myWallet = %v\n", myWallet.Address)

	balance, err := gt.Account.GetBalance(myWallet.Address)
	if err != nil {
		fmt.Printf("failed to get account balance: %v", err)
	}
	fmt.Printf("balance = %v\n", balance)
}
