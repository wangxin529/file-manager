package svc

import (
	"file-manager/config"
	cache2 "file-manager/db/cache"
	"file-manager/db/cache/memory"
	"file-manager/db/cache/redis"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache2.Cache
}

func NewServiceContext(conf config.Config) *ServiceContext {
	var cache cache2.Cache
	if conf.Redis != nil {
		cache = redis.NewRedisCache(conf.Redis)
	}
	cache = memory.NewMemoryCache()
	return &ServiceContext{
		Config: conf,

		Cache: cache,
	}
}
