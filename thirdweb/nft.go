package thirdweb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// NFTMetadata nft元数据
	NFTMetadata struct {
		Name            string      `json:"name,omitempty"`
		Description     string      `json:"description,omitempty"`
		Image           string      `json:"image,omitempty"`
		ExternalURL     string      `json:"external_url,omitempty"`
		AnimationURL    string      `json:"animation_url,omitempty"`
		Properties      interface{} `json:"properties,omitempty"`
		Attributes      interface{} `json:"attributes,omitempty"`
		BackgroundColor string      `json:"background_color,omitempty"`
	}
	// NFTMetadataWithSupply nft带数量的元数据
	NFTMetadataWithSupply struct {
		Metadata NFTMetadata `json:"metadata"`
		Supply   string      `json:"supply"` // 数量
	}
	// NFTTxOverrides 交易覆盖
	NFTTxOverrides struct {
		Gas                  string `json:"gas,omitempty"`                  // Gas limit for the transaction
		MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`         // Maximum fee per gas
		MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"` // Maximum priority fee per gas
		Value                string `json:"value,omitempty"`                // Amount of native currency to send
	}

	// BatchMint721Req 批量铸造721请求
	BatchMint721Req struct {
		Receiver    string         `json:"receiver"`
		Metadatas   []NFTMetadata  `json:"metadatas"`
		TxOverrides NFTTxOverrides `json:"txOverrides,omitempty"`
	}
	// BatchMint1155Req 批量铸造1155请求
	BatchMint1155Req struct {
		Receiver           string                  `json:"receiver"`
		MetadataWithSupply []NFTMetadataWithSupply `json:"metadataWithSupply"`
		TxOverrides        NFTTxOverrides          `json:"txOverrides,omitempty"`
	}

	// LazyMintReq 延迟铸造请求
	LazyMintReq struct {
		Metadatas   []NFTMetadata  `json:"metadatas"`
		TxOverrides NFTTxOverrides `json:"txOverrides,omitempty"`
	}
)

// BatchMint721 批量铸造721 token
func (c *Client) BatchMint721(chain, contractAddr string, req *BatchMint721Req) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("request of batch mint 721 error:%w", err)
	}

	// https://cors.redoc.ly/contract/{chain}/{contractAddress}/erc721/mint-batch-to
	url := fmt.Sprintf("%s/contract/%s/%s/erc721/mint-batch-to", c.baseURL, chain, contractAddr)

	_, err = c.newRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("batch mint 721 error:%w", err)
	}

	return nil
}

// BatchMint1155 批量铸造1155 token
func (c *Client) BatchMint1155(chain, contractAddr string, req *BatchMint1155Req) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("request of batch mint 1155 error:%w", err)
	}

	// https://cors.redoc.ly/contract/{chain}/{contractAddress}/erc1155/mint-batch-to
	url := fmt.Sprintf("%s/contract/%s/%s/erc1155/mint-batch-to", c.baseURL, chain, contractAddr)

	_, err = c.newRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("batch mint 1155 error:%w", err)
	}

	return nil
}

// LazyMint721 延迟铸造721 tokens
func (c *Client) LazyMint721(chain, contractAddr string, req *LazyMintReq) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("request of lazy mint 721 error:%w", err)
	}

	// https://cors.redoc.ly/contract/{chain}/{contractAddress}/erc721/lazy-mint
	url := fmt.Sprintf("%s/contract/%s/%s/erc721/lazy-mint", c.baseURL, chain, contractAddr)

	_, err = c.newRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("lazy mint 721 error:%w", err)
	}

	return nil
}

// LazyMint1155 延迟铸造1155 tokens
func (c *Client) LazyMint1155(chain, contractAddr string, req *LazyMintReq) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("request of lazy mint 1155 error:%w", err)
	}

	// https://cors.redoc.ly/contract/{chain}/{contractAddress}/erc1155/lazy-mint
	url := fmt.Sprintf("%s/contract/%s/%s/erc1155/lazy-mint", c.baseURL, chain, contractAddr)

	_, err = c.newRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("lazy mint 1155 error:%w", err)
	}

	return nil
}
