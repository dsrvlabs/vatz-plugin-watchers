package api

import (
	"testing"
)

func TestGetConsensusPubkey(t *testing.T) {
	// Call the getConsensusPubkey function
	pubkey := GetConsensusPubkey("https://cosmos.blockpi.network/lcd/v1/public/cosmos/staking/v1beta1/validators/", "cosmosvaloper1wlagucxdxvsmvj6330864x8q3vxz4x02rmvmsu")

	// Verify the returned pubkey
	expectedPubkey := "efOai5jzck+C46Zt8ruUcD1w2E7wnDnL9u2ATsODIPg="
	if pubkey != expectedPubkey {
		t.Errorf("Expected pubkey %s, got %s", expectedPubkey, pubkey)
	}
}
