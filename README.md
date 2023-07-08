# Cache Package

The `cache` package provides a simple interface to interact with a Redis cache. It offers methods for setting, getting, deleting, checking existence, and retrieving the time to live (TTL) of cache entries.

## Installation

To install the `cache` package, use the following command:

```shell
go get github.com/thirathawat/cache
```

## Usage

To use the `cache` package, import it into your Go code:

```go
import "github.com/thirathawat/cache"
```

### Creating a Cache

To create a new cache instance, use the `New` function. It establishes a connection to Redis and returns a `Cacher` interface, a cleanup function, and an error. The cleanup function should be called when you're done using the cache to close the underlying connection.

```go
cacher, cleanup, err := cache.New()
if err != nil {
    // Handle error
}
defer cleanup()

// Use the cacher interface to interact with the cache
...
```

### Basic Cache Operations

The `Cacher` interface provides the following methods for basic cache operations:

- `Set`: Sets the value for a given key with an optional expiration time.
- `Get`: Retrieves the value for a given key.
- `Del`: Deletes the value associated with a given key.
- `Exists`: Checks if a given key exists in the cache.
- `TTL`: Retrieves the time to live for a given key.

Example usage:

```go
// Set a value in the cache
err := cacher.Set(ctx, "key", "value", time.Minute)
if err != nil {
    // Handle error
}

// Retrieve a value from the cache
value, err := cacher.Get(ctx, "key")
if err != nil {
    // Handle error
}
fmt.Println(value)

// Delete a value from the cache
err = cacher.Del(ctx, "key")
if err != nil {
    // Handle error
}

// Check if a key exists in the cache
exists, err := cacher.Exists(ctx, "key")
if err != nil {
    // Handle error
}
fmt.Println(exists)

// Retrieve the time to live for a key
ttl, err := cacher.TTL(ctx, "key")
if err != nil {
    // Handle error
}
fmt.Println(ttl)
```

Note: The `ctx` parameter represents the context for the cache operation. You can pass a `context.Context` instance to control the operation's timeout or cancellation.

## Configuration

The `cache` package uses environment variables to configure the cache connection. The following environment variables are available:

- `CACHE_ADDRS`: A comma-separated list of cache addresses.
- `CACHE_PASSWORD`: The password for the cache.
- `CACHE_DB`: The cache database number (default: 0).
- `CACHE_POOL_SIZE`: The cache connection pool size (default: 0).

## Dependencies

The `cache` package relies on the `github.com/go-redis/redis` package to interact with Redis. You can find more information about this package and its usage in the [Go Redis repository](https://github.com/go-redis/redis).

## License

This package is released under the [MIT License](https://opensource.org/licenses/MIT). See the `LICENSE` file for more details.
