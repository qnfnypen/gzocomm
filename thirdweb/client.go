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
	req.Header.Set("Authorization", c.token)
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
			if msg, ok := (body.(map[string]interface{}))["message"].(string); ok {
				return nil, errors.New(msg)
			}
		}
	}

	return rb, nil
}
