package example

import (
	"go-rds-kit/cache"
	"go-rds-kit/config"
	"testing"
)

func NewCache(conf *config.Memory) cache.Cache {

	return cache.NewMemory(conf)

}
func Test_Redis(t *testing.T) {

	//1. 内存

	memory := NewCache(&config.Memory{
		Type: config.CacheMemoryType,
	})
	memory.Set("key", "value", -1)

	//2. 单点 redis
	redis := NewCache(&config.Memory{
		Type: config.CacheRedisType,
		Redis: &config.Redis{
			Addr:     []string{"127.0.0.1:6379"},
			DB:       0,
			Password: "",
		},
	})
	redis.Set("key", "value", -1)

	// 3. 哨兵 redis

	redis = NewCache(&config.Memory{
		Type: config.CacheRedisType,
		Redis: &config.Redis{
			MasterName:       "mymaster",
			SentinelPassword: "",
			Addr:             []string{"127.0.0.1:26379"},
			Password:         "",
			DB:               0,
		},
	})
	redis.Set("key", "value", -1)

}
