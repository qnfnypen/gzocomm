package thirdweb

import (
	"log"
	"testing"
)

var (
	tcli *Client
)

func TestMain(m *testing.M) {
	baseurl := "https://tw-engine.culturevault.com"
	token := ""
	backetWallet := ""
	tcli = NewClient(baseurl, token, backetWallet)

	m.Run()
}

func TestGetAllChainDetail(t *testing.T) {
	resp, err := tcli.GetAllChainDetail()
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range resp.Result {
		log.Println(v)
	}
}

func TestGetChainDetail(t *testing.T) {
	resp, err := tcli.GetChainDetail("1")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp.Result)
}
