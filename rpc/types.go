package rpc

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
