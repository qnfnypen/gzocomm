package thirdweb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Client thirdweb客户端
type Client struct {
	baseURL      string
	token        string
	backedWallet string
}

// API thirdweb-api接口
type API interface {
	// chain
	GetChainDetail(chain string) (*ChainDetailResp, error)
	GetAllChainDetail() (*AllChainDetailResp, error)

	// collection
	DeploySplit(chain string, req *DeploySplitReq) (string, error)
	DeployNFTDrop(chain string, req *DeployNFTDropReq) (string, error)
	DeployEditionDrop(chain string, req *DeployEditionDropReq) (string, error)
	SetRoyaltyDetail(chain, contractAddr string, req *SetRoyaltyDetailReq) error
	OverwriteConditionsFor721(chain, contractAddr string, req *SetClaimCondtitionFor721Req) error
	OverwriteConditionsFor1155(chain, contractAddr string, req *SetClaimCondtitionFor1155Req) error
	ReadeFromContract(chain, contractAddr, funcName string, args ...string) (interface{}, error)

	// nft
	BatchMint721(chain, contractAddr string, req *BatchMint721Req) error
	BatchMint1155(chain, contractAddr string, req *BatchMint1155Req) error
	LazyMint721(chain, contractAddr string, req *LazyMintReq) error
	LazyMint1155(chain, contractAddr string, req *LazyMintReq) error
}

// NewClient 创建 thirdweb-api client
func NewClient(burl, token, bwallet string) *Client {
	return &Client{
		baseURL:      strings.TrimSuffix(strings.TrimSpace(burl), "/"),
		token:        strings.TrimSpace(token),
		backedWallet: strings.TrimSpace(bwallet),
	}
}

// newRequest 新建请求
func (c *Client) newRequest(method, url string, body io.Reader) ([]byte, error) {
	var respBody = make(map[string]interface{})
	cli := &http.Client{}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("new request error:%w", err)
	}
	req.Header.Set("Authorization", "bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-backend-wallet-address", c.backedWallet)

	resp, err := cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("exec do request error:%w", err)
	}
	defer resp.Body.Close()

	rb, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response error:%w", err)
	}
	err = json.Unmarshal(rb, &respBody)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal of response error:%w", err)
	}

	if resp.StatusCode != http.StatusOK {
		if body, ok := respBody["error"]; ok {
			if obj, ok := body.(map[string]interface{}); ok {
				if msg, ok := obj["message"].(string); ok {
					return nil, errors.New(msg)
				}
			}
		}

		return nil, fmt.Errorf("code:%d error:%s", resp.StatusCode, string(rb))
	}

	return rb, nil
}
