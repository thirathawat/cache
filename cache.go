// Package cache provides a simple interface to Redis.
package cache

import (
	"context"
	"time"
)

// Cacher is the interface that wraps the basic cache methods.
type Cacher interface {
	// Set sets the value for the given key.
	Set(ctx context.Context, key, value string, exp time.Duration) error

	// Get returns the value for the given key.
	Get(ctx context.Context, key string) (string, error)

	// Del deletes the value for the given key.
	Del(ctx context.Context, key string) error

	// Exists checks if the given key exists.
	Exists(ctx context.Context, key string) (bool, error)

	// TTL returns the time to live for the given key.
	TTL(ctx context.Context, key string) (time.Duration, error)
}

// cache is a wrapper around client.
type cache struct {
	client *client
}

// New returns a new cache.
func New() (cacher Cacher, cleanup func(), err error) {
	c, err := connect()
	if err != nil {
		return nil, nil, err
	}

	return &cache{c}, func() {
		c.Close()
	}, nil
}

// Set sets the value for the given key.
func (c *cache) Set(ctx context.Context, key, value string, exp time.Duration) error {
	return c.client.Set(ctx, key, value, exp).Err()
}

// Get returns the value for the given key.
func (c *cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

// Del deletes the value for the given key.
func (c *cache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

// Exists checks if the given key exists.
func (c *cache) Exists(ctx context.Context, key string) (bool, error) {
	val, err := c.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return val == 1, nil
}

// TTL returns the time to live for the given key.
func (c *cache) TTL(ctx context.Context, key string) (time.Duration, error) {
	return c.client.TTL(ctx, key).Result()
}
