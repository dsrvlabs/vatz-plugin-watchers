package main

import (
	"context"
	"flag"
	"fmt"
	rpc "github.com/dsrvlabs/vatz-plugin-watchers/rpc/cosmos"

	"github.com/rs/zerolog/log"

	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	pluginName = "watcher-cosmos"
	// Default values.
	defaultAddr                 = "127.0.0.1"
	defaultPort                 = 10001
	defaultRPCEndpoint          = "http://127.0.0.1:26657"
	defaultWarnCountCondition   = 3
	defaultCriticCountCondition = 5
)

var (
	addr                        string
	port                        int
	validatorOperatorAddressHex string
	rpcEndPoint                 string
	warnConditionCount          int
	criticConditionCount        int
	count                       int
)

func init() {
	flag.StringVar(&addr, "addr", defaultAddr, "Listening address")
	flag.IntVar(&port, "port", defaultPort, "Listening port")
	flag.StringVar(&rpcEndPoint, "rpcURI", defaultRPCEndpoint, "Cosmos protocols' RPC Endpoint")
	flag.StringVar(&validatorOperatorAddressHex, "valoperAddr", "", "Cosmos validator address from public key(HEX)")
	flag.IntVar(&warnConditionCount, "warning", defaultWarnCountCondition, "warning thresholds condition.")
	flag.IntVar(&criticConditionCount, "critical", defaultCriticCountCondition, "Critical thresholds condition.")
	flag.Parse()
}

func main() {

	p := sdk.NewPlugin(pluginName)
	if err := p.Register(pluginFeature); err != nil {
		log.Fatal().Err(err).Msg("Failed to register plugin feature")
	}

	ctx := context.Background()
	if err := p.Start(ctx, addr, port); err != nil {
		fmt.Println("exit")
	}
}

// pluginFeature is the main function for the plugin
func pluginFeature(info, option map[string]*structpb.Value) (sdk.CallResponse, error) {

	var (
		severity = pluginpb.SEVERITY_INFO
		state    pluginpb.STATE
	)
	var msg string
	status, err := rpc.HasValidatorSignature(rpcEndPoint, validatorOperatorAddressHex)
	if err != nil {
		// Maybe node wil be killed. So other alert comes to you.
		severity = pluginpb.SEVERITY_CRITICAL
		state = pluginpb.STATE_FAILURE
		msg = "Failed to get validator status"
		log.Info().Str("module", "plugin > watcher_cosmos").Msg(msg)
	} else {
		state = pluginpb.STATE_SUCCESS
		if status {
			count = 0
			severity = pluginpb.SEVERITY_INFO
			msg = "The validator is signing on the block successfully."
		} else {
			count++
			log.Debug().Str("module", "plugin").Msgf("The validator is not signing the block. count: %d", count)
			if count >= criticConditionCount {
				severity = pluginpb.SEVERITY_CRITICAL
				msg = fmt.Sprintf("The validator has failed to sign the block more than %d times.", count)
			} else if count >= warnConditionCount {
				severity = pluginpb.SEVERITY_WARNING
				msg = fmt.Sprintf("The validator has failed to sign the block more than %d times.", count)
			}
		}
		log.Debug().Str("module", "plugin > watcher_cosmos").Msg(msg)
	}

	ret := sdk.CallResponse{
		FuncName: info["execute_method"].GetStringValue(),
		Message:  msg,
		Severity: severity,
		State:    state,
	}

	return ret, nil
}
