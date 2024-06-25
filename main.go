package main

import (
	"fmt"
	"vatz-plugin-watcher-cosmos/api"
	"vatz-plugin-watcher-cosmos/rpc"

	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	// Default values.
	defaultAddr        = "127.0.0.1"
	defaultPort        = 9091
	defaultAPIEndpoint = "http://127.0.0.1:1317"
	defaultRPCEndpoint = "http://127.0.0.1:26657"
	pluginName         = "vatz-plugin-watcher-cosmos"
)

var (
	addr  string
	port  int
	apiep string
	rpcep string
)

func init() {
	// flag.StringVar(&addr, "addr", defaultAddr, "IP Address(e.g. 0.0.0.0, 127.0.0.1)")
	// flag.IntVar(&port, "port", defaultPort, "Port number, default 9091")
	// flag.StringVar(&apiep, "endpoint", defaultAPIEndpoint, "API Endpoint(e.g. http://127.0.0.1:1317)")
	// flag.StringVar(&rpcep, "endpoint", defaultRPCEndpoint, "API Endpoint(e.g. http://127.0.0.1:26657)")
	//
	// flag.Parse()
}

func main() {
	pubKey := api.GetConsensusPubkey("https://cosmos.blockpi.network/lcd/v1/public", "cosmosvaloper1wlagucxdxvsmvj6330864x8q3vxz4x02rmvmsu")
	address, err := rpc.GetValidatorAddressByPubKey("https://cosmos.blockpi.network/rpc/v1/public", pubKey)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("The address for pub_key %s is %s\n", pubKey, address)
	status, err := rpc.HasValidatorSignature("https://cosmos.blockpi.network/rpc/v1/public", address)
	fmt.Printf("The status for validator %s is %v\n", address, status)

	// p := sdk.NewPlugin(pluginName)
	// p.Register(pluginFeature)
	//
	// ctx := context.Background()
	// if err := p.Start(ctx, addr, port); err != nil {
	// 	fmt.Println("exit")
	// }
}

func pluginFeature(info, option map[string]*structpb.Value) (sdk.CallResponse, error) {
	// TODO: Fill here.
	ret := sdk.CallResponse{
		FuncName: "YOUR_FUNCTION_NAME",
		Message:  "YOUR_MESSAGE_CONTENTS",
		Severity: pluginpb.SEVERITY_UNKNOWN,
		State:    pluginpb.STATE_NONE,
	}

	return ret, nil
}
