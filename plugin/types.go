package plugin

import (
	"context"

	v1 "github.com/gatewayd-io/gatewayd-plugin-sdk/plugin/v1"
	"google.golang.org/grpc"
)

type (
	// Priority is the priority of a hook.
	// Smaller values are executed first (higher priority).
	Priority uint
	Method   func(
		context.Context, *v1.Struct, ...grpc.CallOption) (*v1.Struct, error)
)
