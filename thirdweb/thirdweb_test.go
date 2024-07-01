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
	token := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiIweDgwQThjRTgxMjc3ZmJFRDlFMjQ1OTdBZTVmREY3MTRBMkUwNkE3M0MiLCJzdWIiOiIweDMzOGZlNERFZTYzMDMxZTYxYWM5MENDMjFjQzY2MDRmMmMzMDZiMjMiLCJhdWQiOiJ0aGlyZHdlYi5jb20iLCJleHAiOjQ4NzI3ODM1NzgsIm5iZiI6MTcxOTE4MzU3OCwiaWF0IjoxNzE5MTgzNTc4LCJqdGkiOiIyMmEyZjMyYS1iMmY5LTRmMjQtYTkyMS0xZmI5N2NjZTM5MTciLCJjdHgiOnsicGVybWlzc2lvbnMiOiJBRE1JTiJ9fQ.MHgyYWUyZmFiYjg3MjRmZDg1MTFiYTljMzc4NWQyMTNjYWI2N2M0YjhmMzY3MGExYmZkMTc5MTU0NWFlZjExNmU5NTRhMTVjZjY5NTExODY5MWRhOTYzNGQzYmM5MmU5N2Q4OThlYjAzYzJhZjhmM2Y4YmZiMWY3NDUwNzQzYWI4MTFi"
	backetWallet := "0x80a0254A7c07562d0fF9E144ED00ff6c43818971"
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
