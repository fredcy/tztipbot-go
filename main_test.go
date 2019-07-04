package tztipbot

import (
	//"fmt"
	"encoding/json"
	"fmt"
	goTezos "github.com/DefinitelyNotAGoat/go-tezos"
	"os"
	"testing"
)

var RpcAddr string

func TestMain(m *testing.M) {
	// set up shared info for tests, then run them
	RpcAddr = os.Getenv("RPCADDR")
	if RpcAddr == "" {
		fmt.Printf("Error: RPCADDR must be set before running tests\n")
		os.Exit(1)
	}
	code := m.Run()
	os.Exit(code)
}

func TestHello(t *testing.T) {
	gt, err := goTezos.NewGoTezos(RpcAddr)
	if err != nil {
		t.Errorf("could not connect to network: %v", err)
	}

	block, err := gt.Block.GetHead()
	if err != nil {
		t.Errorf("count not get head: %v", err)
	}
	t.Log(PrettyReport(block))
}

//Takes an interface v and returns a pretty json string.
func PrettyReport(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		return string(b)
	}
	return ""
}
