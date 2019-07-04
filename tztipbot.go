package tztipbot

import (
	"fmt"
	goTezos "github.com/DefinitelyNotAGoat/go-tezos"
	"os"
)

func ConnectService() (*goTezos.GoTezos, error) {
	rpcAddr := os.Getenv("RPCADDR")
	if rpcAddr == "" {
		return nil, fmt.Errorf("RPCADDR must be set")
	}

	gt, err := goTezos.NewGoTezos(rpcAddr)
	if err != nil {
		return nil, fmt.Errorf("could not connect to network: %v", rpcAddr)
	}

	return gt, nil
}
