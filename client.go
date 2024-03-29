package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

const Nil = redis.Nil

type (
	Cmd             = redis.Cmd
	Client          = redis.Cmdable
	Cmdable         = redis.Cmdable
	XAddArgs        = redis.XAddArgs
	XReadArgs       = redis.XReadArgs
	XReadGroupArgs  = redis.XReadGroupArgs
	XPendingExtArgs = redis.XPendingExtArgs
	XClaimArgs      = redis.XClaimArgs
	XAutoClaimArgs  = redis.XAutoClaimArgs
	XStream         = redis.XStream
	XMessage        = redis.XMessage
)

func NewCmd(ctx context.Context, args ...interface{}) *redis.Cmd {
	return redis.NewCmd(ctx)
}

var _ Client = (*redis.Client)(nil)
var _ Client = (*redis.ClusterClient)(nil)

func NewClient(o *Options) Client {
	// 没节点直接抛异常
	if len(o.NodeOptions) == 0 {
		panic("require one node at lease")
	}
	// 单节点返回普通的Client
	if len(o.NodeOptions) == 1 {
		return redis.NewClient(o.NodeOptions[0])
	}
	// 多节点返回集群Client
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        o.getAddrs(),
		NewClient:    o.getNewClient(),
		ClusterSlots: o.getClusterSlots(),
		ReadOnly:     false,
	})
}
