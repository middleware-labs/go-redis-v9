package redis

import (
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

func InstrumentTracing(rdb redis.UniversalClient, opts ...redisotel.TracingOption) error {
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		return err
	}
	return nil
}
