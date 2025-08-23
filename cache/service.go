package cache

import (
	"go-rds-kit/config"
	"time"
)

type Cache interface {
	Set(key string, value interface{}, expireTime time.Duration)
	Get(key string, result interface{}) bool
	Delete(key string)
	Increment(key string, count int64) error
	Exists(key string) bool
}

var cacheOperate = CacheOperate{}

type CacheOperate struct { // 基础实现类
}

func (c *CacheOperate) Set(key string, value interface{}, expireTime time.Duration) {

}
func (c *CacheOperate) Get(key string, result interface{}) bool {
	return false
}
func (c *CacheOperate) Delete(key string) {

}

func (c *CacheOperate) Increment(key string, count int64) error {
	return nil
}
func (c *CacheOperate) Exists(key string) bool {
	return false
}

func NewMemory(conf *config.Memory) Cache {
	if conf == nil || conf.Type == config.CacheMemoryType {
		return NewMemoryCache()
	}
	return NewRedisCache(conf.Redis)

}
