package appcache

import (
	"github.com/lokesh-go/go-api-microservice/internal/config"
	appCachePkg "github.com/lokesh-go/go-api-microservice/pkg/cache/appcache"
)

// AppCache conn
type appCache struct {
	conn *appCachePkg.AppCache
}

// New app cache
func New(config config.Methods) *appCache {
	// Forms appCache config
	ac := &appCachePkg.Config{
		DefaultExpirationInSeconds: config.Get().DAL.Cache.AppCache.DefaultExpirationInSeconds,
		CleanupIntervalInMinutes:   config.Get().DAL.Cache.AppCache.CleanupIntervalInMinutes,
	}

	// Get the appCache connection
	return &appCache{
		conn: appCachePkg.New(ac),
	}
}
