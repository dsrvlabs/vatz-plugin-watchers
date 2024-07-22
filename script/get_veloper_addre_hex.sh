#!/bin/bash

# Function to prompt for user input
prompt_for_input() {
    local prompt_message=$1
    read -p "$prompt_message: " input_variable
    echo $input_variable
}

# Prompt for necessary inputs
REST_END_POINT=$(prompt_for_input "Enter the Rest Endpoint")
RPC_END_POINT=$(prompt_for_input "Enter the RPC Endpoint")
VALOPER_ADDR=$(prompt_for_input "Enter the valoper address")

# Display entered values and ask for confirmation
echo -e "\nYou have entered the following details:"
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

if [ -z "$CONS_PUB_KEY" ]; then
    echo "Failed to fetch the consensus public key."
    exit 1
fi

# echo "Consensus public key: $CONS_PUB_KEY"

# Function to fetch validators from a specific page
fetch_validators() {
    local page=$1
    curl -s "${RPC_END_POINT}/validators?page=${page}" | jq -r '.result.validators[]'
}

# Initialize variables for pagination
PAGE=1
VALIDATOR_ADDR=""

# Loop through pages to find the validator address
while : ; do
    VALIDATORS=$(fetch_validators $PAGE)

    # If no validators are returned, break the loop
    if [ -z "$VALIDATORS" ]; then
        break
    fi

    # Check if the desired validator is in the current page
    VALIDATOR_ADDR=$(echo "$VALIDATORS" | jq -r --arg value "$CONS_PUB_KEY" 'select(.pub_key.value == $value) | .address')

    # If the validator address is found, break the loop
    if [ -n "$VALIDATOR_ADDR" ]; then
        break
    fi

    # Move to the next page
    PAGE=$((PAGE + 1))
done

if [ -z "$VALIDATOR_ADDR" ]; then
    echo "Failed to fetch the validator hex address."
    exit 1
fi

echo "Validator hex address: $VALIDATOR_ADDR"
