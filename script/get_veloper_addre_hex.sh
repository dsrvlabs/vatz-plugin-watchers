#!/bin/bash

# Function to prompt for user input
prompt_for_input() {
    local prompt_message=$1
    local input_variable
    read -p "$prompt_message: " input_variable
    echo $input_variable
}

# Prompt for RPC address
REST_END_POINT=$(prompt_for_input "Enter the Rest Endpoint")
# Prompt for endpoint
RPC_END_POINT=$(prompt_for_input "Enter the RPC Endpoint")
# Prompt for valoper address
VALOPER_ADDR=$(prompt_for_input "Enter the valoper address")

# Display the entered values and ask for confirmation
echo " "
echo "You have entered the following details:"
echo "Rest Endpoint: $REST_END_POINT"
echo "RPC Endpoint: $RPC_END_POINT"
echo "Valoper(Validator Operator) address: $VALOPER_ADDR"

read -p "Are the entered values correct? Do you wish to proceed? (y/n): " confirmation
confirmation=$(echo "$confirmation" | tr '[:upper:]' '[:lower:]')

if [[ "$confirmation" != "y" ]]; then
    echo "Operation cancelled."
    exit 0
fi

# Fetch the consensus public key
CONS_PUB_KEY=$(curl -s "${REST_END_POINT}/cosmos/staking/v1beta1/validators/${VALOPER_ADDR}" | jq -r '.validator.consensus_pubkey.key')

# Check if the consensus public key was fetched successfully
if [ -z "$CONS_PUB_KEY" ]; then
    echo "Failed to fetch the consensus public key."
    exit 1
fi

# Fetch the validator address using the consensus public key
VALIDATOR_ADDR=$(curl -s "${RPC_END_POINT}/validators" | jq -r --arg value "$CONS_PUB_KEY" '.result.validators[] | select(.pub_key.value == $value) | .address')


# Check if the validator address was fetched successfully
if [ -z "$VALIDATOR_ADDR" ]; then
    echo "Failed to fetch the validator hex address."
    exit 1
fi

# Output the final validator address
echo "Validator hex address: $VALIDATOR_ADDR"
