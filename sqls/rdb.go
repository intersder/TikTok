package sqls

import (
	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

func OpenRedisClient(addr string, password string, db int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
}

func RDB() *redis.Client {
	return rdb
}
