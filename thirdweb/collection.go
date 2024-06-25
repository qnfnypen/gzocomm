package thirdweb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// Recipient 接收者
	Recipient struct {
		Address   string `json:"address"`
		SharesBps string `json:"sharesBps"`
	}
	// DeploySplitReq 创建分账合约请求
	DeploySplitReq struct {
		Metadata struct {
			Name              string      `json:"name"`               // required
			Recipients        []Recipient `json:"recipients"`         // required
			TrustedForwarders []string    `json:"trusted_forwarders"` // default []
		} `json:"contractMetadata"`
	}

	// DeployNFTDropReq 创建721合约请求
	DeployNFTDropReq struct {
		Metadata struct {
			Name                 string `json:"name"` // required
			Desc                 string `json:"description"`
			Image                string `json:"image"`
			ExternalLink         string `json:"external_link"`
			AppURL               string `json:"app_uri"`
			DefaultAdmin         string `json:"defaultAdmin"`
			SellerFeeBasisPoints int64  `json:"seller_fee_basis_points"` // required
			FeeRecipient         string `json:"fee_recipient"`           // required
			Merkle               struct {
				PropertyName string `json:"property name"`
			} `json:"merkle"`
			Symbol                 string   `json:"symbol"`                    // required
			PlatformFeeBasisPoints int64    `json:"platform_fee_basis_points"` // required
			PlatformFeeRecipient   string   `json:"platform_fee_recipient"`    // default 0
			PrimarySaleRecipient   string   `json:"primary_sale_recipient"`    // required
			TrustedForwarders      []string `json:"trusted_forwarders"`        // default []
		} `json:"contractMetadata"`
	}

	// DeployEditionDropReq 创建1155合约请求
	DeployEditionDropReq struct {
		Metadata struct {
			Name                 string `json:"name"` // required
			Desc                 string `json:"description"`
			Image                string `json:"image"`
			ExternalLink         string `json:"external_link"`
			AppURL               string `json:"app_uri"`
			DefaultAdmin         string `json:"defaultAdmin"`
			SellerFeeBasisPoints int64  `json:"seller_fee_basis_points"` // required
			FeeRecipient         string `json:"fee_recipient"`           // required
			Merkle               struct {
				PropertyName string `json:"property name"`
			} `json:"merkle"`
			Symbol                 string   `json:"symbol"`                    // required
			PlatformFeeBasisPoints int64    `json:"platform_fee_basis_points"` // required
			PlatformFeeRecipient   string   `json:"platform_fee_recipient"`    // default 0
			PrimarySaleRecipient   string   `json:"primary_sale_recipient"`    // required
			TrustedForwarders      []string `json:"trusted_forwarders"`        // default []
		} `json:"contractMetadata"`
	}

	// SetRoyaltyDetailReq 设置合约版税
	SetRoyaltyDetailReq struct {
		SellerFeeBasisPoints int64  `json:"seller_fee_basis_points"` // required
		FeeRecipient         string `json:"fee_recipient"`           // required
	}

	// Snapshot 快照
	Snapshot struct {
		Price           string `json:"price,omitempty"`
		CurrencyAddress string `json:"currencyAddress,omitempty"`
		Address         string `json:"address,omitempty"`
		MaxClaimable    string `json:"maxClaimable,omitempty"`
	}
	// ClaimCondition 条件
	ClaimCondition struct {
		MaxClaimableSupply    interface{} `json:"maxClaimableSupply,omitempty"`    // string or number
		StartTimestamp        interface{} `json:"startTime,omitempty"`             // string or number
		Price                 interface{} `json:"price,omitempty"`                 // string or number
		CurrencyAddress       string      `json:"currencyAddress,omitempty"`       // string or number
		MaxClaimablePerWallet interface{} `json:"maxClaimablePerWallet,omitempty"` // string or number
		WaitInSeconds         interface{} `json:"waitInSeconds,omitempty"`         // string or number
		MerkleRootHash        interface{} `json:"merkleRootHash,omitempty"`        // string or array of numbers
		Metadata              struct {
			Name string `json:"name"`
		} `json:"metadata"`
	}
	// SetClaimCondtitionFor721Req 设置721合约的drop条件
	SetClaimCondtitionFor721Req struct {
		ClaimConditionInputs ClaimCondition `json:"claimConditionInputs"`
	}
	// SetClaimCondtitionFor1155Req 设置1155合约的drop条件
	SetClaimCondtitionFor1155Req struct {
		TokenID              interface{}    `json:"tokenId"` // string or number
		ClaimConditionInputs ClaimCondition `json:"claimConditionInputs"`
	}

	// BaseContractResp 合约基础响应
	BaseContractResp struct {
		Result struct {
			QueueID         string `json:"queueId"`
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

	// [post] https://cors.redoc.ly/deploy/{chain}/prebuilts/split
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

	// [post] https://cors.redoc.ly/deploy/{chain}/prebuilts/nft-drop
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

// DeployEditionDrop 创建1155合约
func (c *Client) DeployEditionDrop(chain string, req *DeployEditionDropReq) (string, error) {
	var baseResp = new(BaseContractResp)

	if req.Metadata.TrustedForwarders == nil {
		req.Metadata.TrustedForwarders = make([]string, 0)
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("json marshal request of deploy edition-drop contract error:%w", err)
	}

	// [post] https://cors.redoc.ly/deploy/{chain}/prebuilts/edition-drop
	url := fmt.Sprintf("%s/deploy/%s/prebuilts/edition-drop", c.baseURL, chain)

	resp, err := c.newRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("deploy edition-drop contract error:%w", err)
	}
	err = json.Unmarshal(resp, baseResp)
	if err != nil {
		return "", fmt.Errorf("deploy edition-drop contract error,json unmarshal body of response error:%w", err)
	}

	return baseResp.Result.DeployedAddress, nil
}

// SetRoyaltyDetail 设置合约版税
func (c *Client) SetRoyaltyDetail(chain, contractAddr string, req *SetRoyaltyDetailReq) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json marshal request of set royalty detail error:%w", err)
	}

	// [post] https://cors.redoc.ly/contract/{chain}/{contractAddress}/royalties/set-default-royalty-info
	url := fmt.Sprintf("%s/contract/%s/%s/royalties/set-default-royalty-info", c.baseURL, chain, contractAddr)
	_, err = c.newRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("set royalty detail error:%w", err)
	}

	return nil
}

// OverwriteConditionsFor721 设置721合约的购买条件
func (c *Client) OverwriteConditionsFor721(chain, contractAddr string, req *SetClaimCondtitionFor721Req) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json marshal request of set royalty detail error:%w", err)
	}

	// https://cors.redoc.ly/contract/{chain}/{contractAddress}/erc721/claim-conditions/set
	url := fmt.Sprintf("%s/contract/%s/%s/erc721/claim-conditions/set", c.baseURL, chain, contractAddr)
	_, err = c.newRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("set royalty detail error:%w", err)
	}

	return nil
}

// OverwriteConditionsFor1155 设置1155合约的购买条件
func (c *Client) OverwriteConditionsFor1155(chain, contractAddr string, req *SetClaimCondtitionFor1155Req) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json marshal request of set royalty detail error:%w", err)
	}

	// https://cors.redoc.ly/contract/{chain}/{contractAddress}/erc1155/claim-conditions/set
	url := fmt.Sprintf("%s/contract/%s/%s/erc1155/claim-conditions/set", c.baseURL, chain, contractAddr)
	_, err = c.newRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("set royalty detail error:%w", err)
	}

	return nil
}
