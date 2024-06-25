package main

// Block represents the structure of a block in the blockchain
type Block struct {
	Result struct {
		Block struct {
			LastCommit struct {
				Signatures []struct {
					ValidatorAddress string `json:"validator_address"`
				} `json:"signatures"`
			} `json:"last_commit"`
		} `json:"block"`
	} `json:"result"`
}

// Validator represents the structure of a validator
type Validator struct {
	Validator struct {
		ConsensusPubkey struct {
			Key string `json:"key"`
		} `json:"consensus_pubkey"`
	} `json:"validator"`
}

// ValidatorsResponse represents the response structure for validators
type ValidatorsResponse struct {
	Result struct {
		Validators []struct {
			PubKey struct {
				Value string `json:"value"`
			} `json:"pub_key"`
			Address string `json:"address"`
		} `json:"validators"`
	} `json:"result"`
}
