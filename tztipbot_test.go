package tztipbot

import (
	"encoding/json"
	"fmt"
	goTezos "github.com/DefinitelyNotAGoat/go-tezos"
	"os"
	"testing"
)

var Gt *goTezos.GoTezos

func TestMain(m *testing.M) {
	// set up shared info for tests, then run them
	var err error
	Gt, err = ConnectService()
	if err != nil {
		fmt.Printf("could not connect to network: %v\n", err)
		os.Exit(2)
	}

	code := m.Run()
	os.Exit(code)
}

func TestHello(t *testing.T) {
	block, err := Gt.Block.GetHead()
	if err != nil {
		t.Errorf("count not get head: %v", err)
	}

	t.Log(PrettyReport(block.Header))
}

//Takes an interface v and returns a pretty json string.
func PrettyReport(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		return string(b)
	}
	return ""
}
