package act

import (
	"context"
	"errors"
	"github.com/hashicorp/go-hclog"
	"github.com/redis/go-redis/v9"
)

type MessageHandlerFn func(ctx context.Context, payload []byte) error

type IConsumer interface {
	Subscribe(ctx context.Context, handler MessageHandlerFn) error
}

var _ IConsumer = (*Consumer)(nil)

type Consumer struct {
	logger       hclog.Logger
	rdb          *redis.Client
	workersCount int
	channel      string
	redisPubSub  *redis.PubSub
	redisChan    <-chan *redis.Message
}

func NewConsumer(
	logger hclog.Logger, redisClient *redis.Client, workersCount int, channelName string,
) (*Consumer, error) {
	if logger == nil {
		return nil, errors.New("logger is required")
	}
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		logger.Error("failed to connect to redis", "error", err)
		return nil, err
	}
	return &Consumer{
		logger:       logger,
		rdb:          redisClient,
		channel:      channelName,
		workersCount: workersCount,
	}, nil
}

func (c *Consumer) Subscribe(ctx context.Context, handler MessageHandlerFn) error {
	c.redisPubSub = c.rdb.Subscribe(ctx, c.channel)
	if err := c.redisPubSub.Ping(context.Background()); err != nil {
		c.logger.Error("failed to subscribe to redis channel", "error", err, "channel", c.channel)
		return err
	}
	c.redisChan = c.redisPubSub.Channel()
	for i := 0; i < c.workersCount; i++ {
		go c.run(ctx, handler)
	}
	return nil
}

func (c *Consumer) run(ctx context.Context, handler MessageHandlerFn) {
	for {
		select {
		case <-ctx.Done():
			c.logger.Info("context done, stopping redis consumer")
			if err := c.redisPubSub.Close(); err != nil {
				c.logger.Error("failed to close redis pubsub", "error", err)
			}
			return
		case msg, ok := <-c.redisChan:
			if !ok {
				c.logger.Warn("redis channel closed")
				return
			}
			if err := handler(ctx, []byte(msg.Payload)); err != nil {
				c.logger.Error("failed to process task", "error", err, "payload", msg.Payload)
			} else {
				c.logger.Debug("async redis task processed successfully", "payload", msg.Payload)
			}
		}
	}
}
