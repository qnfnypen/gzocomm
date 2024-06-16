package thirdweb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Recipient struct {
		Address   string `json:"address"`
		SharesBps string `json:"sharesBps"`
	}
	SplitMetadata struct {
		Name              string      `json:"name"`
		Recipients        []Recipient `json:"recipients"`
		TrustedForwarders []string    `json:"trusted_forwarders"`
	}
	DeploySplitReq struct {
		Metadata SplitMetadata `json:"contractMetadata"`
	}
	DeployNFTDropReq struct {
		Metadata struct {
			Name                 string `json:"name"`
			Desc                 string `json:"description"`
			Image                string `json:"image"`
			ExternalLink         string `json:"external_link"`
			AppURL               string `json:"app_uri"`
			DefaultAdmin         string `json:"defaultAdmin"`
			SellerFeeBasisPoints int64  `json:"seller_fee_basis_points"`
			FeeRecipient         string `json:"fee_recipient"`
			Merkle               struct {
				PropertyName string `json:"property name"`
			} `json:"merkle"`
			Symbol                 string   `json:"symbol"`
			PlatformFeeBasisPoints int64    `json:"platform_fee_basis_points"`
			PlatformFeeRecipient   string   `json:"platform_fee_recipient"`
			PrimarySaleRecipient   string   `json:"primary_sale_recipient"`
			TrustedForwarders      []string `json:"trusted_forwarders"`
		} `json:"contractMetadata"`
	}

	BaseContractResp struct {
		Result struct {
			QueueId         string `json:"queueId"`
			DeployedAddress string `json:"deployedAddress"`
		} `json:"result"`
	}
)

// DeploySplit 部署分账合约，chain: id or name
func (c *Client) DeploySplit(chain string, req *DeploySplitReq) (string, error) {
	var splitResp = new(BaseContractResp)

	if req.Metadata.TrustedForwarders == nil {
		req.Metadata.TrustedForwarders = make([]string, 0)
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("json marshal request of deploy split contract error:%w", err)
	}

	url := fmt.Sprintf("%s/deploy/%s/prebuilts/split", c.baseURL, chain)

	resp, err := c.newRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("deploy split contract error:%w", err)
	}
	err = json.Unmarshal(resp, splitResp)
	if err != nil {
		return "", fmt.Errorf("deploy split contract error,json unmarshal body of response error:%w", err)
	}

	return splitResp.Result.DeployedAddress, nil
}

// DeployNFTDrop 创建721合约
func (c *Client) DeployNFTDrop(chain string, req *DeployNFTDropReq) (string, error) {
	var baseResp = new(BaseContractResp)

	if req.Metadata.TrustedForwarders == nil {
		req.Metadata.TrustedForwarders = make([]string, 0)
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("json marshal request of deploy nft-drop contract error:%w", err)
	}

	url := fmt.Sprintf("%s/deploy/%s/prebuilts/nft-drop", c.baseURL, chain)

	resp, err := c.newRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("deploy nft-drop contract error:%w", err)
	}
	err = json.Unmarshal(resp, baseResp)
	if err != nil {
		return "", fmt.Errorf("deploy nft-drop contract error,json unmarshal body of response error:%w", err)
	}

	return baseResp.Result.DeployedAddress, nil
}
