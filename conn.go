package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/kelseyhightower/envconfig"
)

// client is a wrapper around redis.UniversalClient.
type client struct {
	redis.UniversalClient
}

// config holds the configuration for the cache connection.
type config struct {
	// Addrs is a list of cache addresses.
	Addrs []string `envconfig:"ADDRS"`

	// Password is the cache password.
	Password string `envconfig:"PASSWORD"`

	// Database to be selected after connecting to the server.
	DB int `envconfig:"DB" default:"0"`

	// PoolSize is the cache pool size.
	PoolSize int `envconfig:"POOL_SIZE" default:"0"`
}

// readConfig reads the configuration from the environment.
func readConfig() *config {
	cfg := new(config)
	envconfig.MustProcess("CACHE", cfg)
	return cfg
}

// connect connects to the cache and returns a client.
func connect() (*client, error) {
	cfg := readConfig()

	c := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    cfg.Addrs,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &client{c}, nil
}
