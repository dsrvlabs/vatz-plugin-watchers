package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type validatorResponse struct {
	Validator struct {
		ConsensusPubKey struct {
			Type string `json:"@type"`
			Key  string `json:"key"`
		} `json:"consensus_pubkey"`
	} `json:"validator"`
}

// GetConsensusPubkey retrieves the consensus public key for a given validator.
func GetConsensusPubkey(baseAPI string, valoper string) string {
	// Make the GET request
	resp, err := http.Get(baseAPI + "/cosmos/staking/v1beta1/validators/" + valoper)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ""
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}

	// Parse the JSON response
	var validatorResp validatorResponse
	err = json.Unmarshal(body, &validatorResp)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ""
	}

	// Extract the key value
	key := validatorResp.Validator.ConsensusPubKey.Key

	return key
}
