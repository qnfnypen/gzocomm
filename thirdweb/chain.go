package thirdweb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// ChainDetail 链详情
	ChainDetail struct {
		Name           string   `json:"name"`
		Chain          string   `json:"chain"`
		RPC            []string `json:"rpc"`
		NativeCurrency struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Decimals int64  `json:"decimals"`
		}
		ShortName string `json:"shortName"`
		ChainID   int64  `json:"chainId"`
		TestNet   bool   `json:"testnet"`
		Slug      string `json:"slug"`
	}
	// AllChainDetailResp 获取所有链的响应
	AllChainDetailResp struct {
		Result []ChainDetail `json:"result"`
	}
	// ChainDetailResp 获取链详情
	ChainDetailResp struct {
		Result ChainDetail `json:"result"`
	}
)

// GetAllChainDetail 获取所有链的详细信息
func (c *Client) GetAllChainDetail() (*AllChainDetailResp, error) {
	var details = new(AllChainDetailResp)

	// https://cors.redoc.ly/chain/get-all
	url := fmt.Sprintf("%s/chain/get-all", c.baseURL)

	resp, err := c.newRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("get detail of all-chain error:%w", err)
	}
	err = json.Unmarshal(resp, details)
	if err != nil {
		return nil, fmt.Errorf("get detail of all-chain error,json unmarshal body of response error:%w", err)
	}

	return details, nil
}

// GetChainDetail 通过链ID或者链name获取链详情
func (c *Client) GetChainDetail(chain string) (*ChainDetailResp, error) {
	var detail = new(ChainDetailResp)

	// https://cors.redoc.ly/chain/get
	url := fmt.Sprintf("%s/chain/get?chain=%s", c.baseURL, chain)

	resp, err := c.newRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("get detail of chain(%s) error:%w", chain, err)
	}
	err = json.Unmarshal(resp, detail)
	if err != nil {
		return nil, fmt.Errorf("get detail of chain(%s) error,json unmarshal body of response error:%w", chain, err)
	}

	return detail, nil
}
