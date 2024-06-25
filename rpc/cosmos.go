package rpc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Validator represents a single validator in the response
type Validator struct {
	Address string `json:"address"`
	PubKey  struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"pub_key"`
	VotingPower      string `json:"voting_power"`
	ProposerPriority string `json:"proposer_priority"`
}

// ValidatorsResponse represents the structure of the entire response
type ValidatorsResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockHeight string      `json:"block_height"`
		Validators  []Validator `json:"validators"`
		Count       string      `json:"count"`
		Total       string      `json:"total"`
	} `json:"result"`
}

// BlockResponse represents the structure of a block in the cosmos
type BlockResponse struct {
	Result struct {
		Block struct {
			LastCommit struct {
				Signatures []struct {
					ValidatorAddress string `json:"validator_address"`
					Signature        string `json:"signature"`
				} `json:"signatures"`
			} `json:"last_commit"`
		} `json:"block"`
	} `json:"result"`
}

// GetValidatorAddressByPubKey finds the validator address by its pub_key value
func GetValidatorAddressByPubKey(basRPC string, pubKey string) (string, error) {
	resp, err := http.Get(basRPC + "/validators?page=2&per_page=180")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var validatorsResponse ValidatorsResponse
	if err := json.Unmarshal(body, &validatorsResponse); err != nil {
		return "", err
	}

	for _, validator := range validatorsResponse.Result.Validators {
		if validator.PubKey.Value == pubKey {
			return validator.Address, nil
		}
	}

	return "", fmt.Errorf("validator with pubKey %s not found", pubKey)
}

func HasValidatorSignature(basRPC string, validatorAddress string) (bool, error) {
	// HTTP GET request
	resp, err := http.Get(basRPC + "/block")
	if err != nil {
		return false, fmt.Errorf("error fetching the block: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("error reading response body: %w", err)
	}

	// Unmarshal JSON response
	var blockResponse BlockResponse
	if err := json.Unmarshal(body, &blockResponse); err != nil {
		return false, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// Check for signature
	for _, signature := range blockResponse.Result.Block.LastCommit.Signatures {
		if signature.ValidatorAddress == validatorAddress {
			return true, nil
		}
	}

	return false, nil
}
