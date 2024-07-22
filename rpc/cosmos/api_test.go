package cosmos

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock response for the /block endpoint
const mockBlockResponse = `
{
	"result": {
		"block": {
			"last_commit": {
				"signatures": [
					{
						"validator_address": "test-validator-address",
						"signature": "test-signature"
					}
				]
			}
		}
	}
}
`

// Mock response with error
const mockErrorResponse = `{
	"error": "some error"
}`

// TestHasValidatorSignature tests the HasValidatorSignature function
func TestHasValidatorSignature(t *testing.T) {
	// Create a mock server that returns the mock response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockBlockResponse))
	}))
	defer server.Close()

	tests := []struct {
		name             string
		validatorAddress string
		expectedResult   bool
		expectedError    bool
	}{
		{
			name:             "Valid signature",
			validatorAddress: "test-validator-address",
			expectedResult:   true,
			expectedError:    false,
		},
		{
			name:             "Invalid signature",
			validatorAddress: "invalid-validator-address",
			expectedResult:   false,
			expectedError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := HasValidatorSignature(server.URL, tt.validatorAddress)
			if result != tt.expectedResult {
				t.Errorf("expected %v, got %v", tt.expectedResult, result)
			}
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}

// TestHasValidatorSignature_Error tests error handling in HasValidatorSignature function
func TestHasValidatorSignature_Error(t *testing.T) {
	// Create a mock server that returns a response with invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	_, err := HasValidatorSignature(server.URL, "any-validator-address")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

// TestHasValidatorSignature_NetworkError tests network error handling in HasValidatorSignature function
func TestHasValidatorSignature_NetworkError(t *testing.T) {
	// Create a mock server that will be closed to simulate a network error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	server.Close()

	_, err := HasValidatorSignature(server.URL, "any-validator-address")
	if err == nil {
		t.Errorf("expected network error, got nil")
	}
}
