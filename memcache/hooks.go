package memcache

import (
	"context"
	"net"
)

var (
	// OnServerPicked allows hooking/extending ServerSelector with additional behaviour after a server has been picked
	OnServerPicked = defaultOnServerPicked

	// AroundEach allows hooking/extending ServerSelector with additional behaviour around each method call per server
	AroundEach = defaultAroundEach
)

func defaultOnServerPicked(_ context.Context, _ net.Addr) {}

// defaultAroundEach
func defaultAroundEach(ctx context.Context, a net.Addr, fn func(context.Context, net.Addr) error) error {
	return fn(ctx, a)
}
