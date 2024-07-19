package main

import (
	"context"
	"flag"
	"fmt"

	rpc "github.com/dsrvlabs/vatz-plugin-watcher-cosmos/rpc"
	"github.com/rs/zerolog/log"

	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	// Default values.
	defaultAddr           = "127.0.0.1"
	defaultPort           = 10001
	defaultRPCEndpoint    = "http://127.0.0.1:26657"
	pluginName            = "watcher-cosmos"
	defaultAlertCondition = 3
)

var (
	addr             string
	port             int
	validatorAddress string
	rpcep            string
	aleartCount      int
	count            int
)

func init() {
	flag.StringVar(&addr, "addr", defaultAddr, "Listening address")
	flag.IntVar(&port, "port", defaultPort, "Listening port")
	flag.StringVar(&rpcep, "rpcurl", defaultRPCEndpoint, "Cosmos RPC Endpoint")
	flag.StringVar(&validatorAddress, "validator", "", "Cosmos validator address from public key")
	flag.IntVar(&aleartCount, "condition", defaultAlertCondition, "Thresholds that trigger notifications")

	flag.Parse()
}

func main() {

	p := sdk.NewPlugin(pluginName)
	p.Register(pluginFeature)

	ctx := context.Background()
	if err := p.Start(ctx, addr, port); err != nil {
		fmt.Println("exit")
	}
}

// pluginFeature is the main function for the plugin
func pluginFeature(info, option map[string]*structpb.Value) (sdk.CallResponse, error) {

	severity := pluginpb.SEVERITY_INFO
	state := pluginpb.STATE_NONE

	var msg string

	status, err := rpc.HasValidatorSignature(rpcep, validatorAddress)

	if err == nil {
		state = pluginpb.STATE_SUCCESS
		if status == true {
			count = 0
			severity = pluginpb.SEVERITY_INFO
			msg = fmt.Sprintf("The validator is signing the block successfully.")
		} else {
			count++
			if count >= aleartCount {
				severity = pluginpb.SEVERITY_CRITICAL
				msg = fmt.Sprintf("")
			}
		}
		log.Debug().Str("module", "plugin").Msg(msg)
	} else {
		// Maybe node wil be killed. So other alert comes to you.
		severity = pluginpb.SEVERITY_CRITICAL
		state = pluginpb.STATE_FAILURE
		msg = "Failed to get validator status"
		log.Info().Str("moudle", "plugin").Msg(msg)
	}

	ret := sdk.CallResponse{
		FuncName: info["execute_method"].GetStringValue(),
		Message:  msg,
		Severity: severity,
		State:    state,
	}

	return ret, nil
}
