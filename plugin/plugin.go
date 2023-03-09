package plugin

import (
	"math"

	pluginV1 "github.com/gatewayd-io/gatewayd-plugin-sdk/plugin/v1"
	goplugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type Identifier struct {
	Name      string
	Version   string
	RemoteURL string
	Checksum  string
}

type Plugin struct {
	goplugin.NetRPCUnsupportedPlugin
	pluginV1.GatewayDPluginServiceServer

	Client *goplugin.Client

	ID          Identifier
	Description string
	Authors     []string
	License     string
	ProjectURL  string
	LocalPath   string
	Args        []string
	Env         []string
	Enabled     bool
	// internal and external config options
	Config map[string]string
	// hooks it attaches to
	Hooks    []pluginV1.HookName
	Priority Priority
	// required plugins to be loaded before this one
	// Built-in plugins are always loaded first
	Requires   []Identifier
	Tags       []string
	Categories []string
}

func DefaultGRPCServer(opts []grpc.ServerOption) *grpc.Server {
	// Increase the max message size to 2GB.
	opts = append(opts,
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(math.MaxInt32)))
	return grpc.NewServer(opts...)
}
