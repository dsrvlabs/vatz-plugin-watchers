package rpc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HasValidatorSignature checks if a validator has signed the block
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
