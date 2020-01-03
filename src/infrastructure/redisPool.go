package infrastructure

import (
	"github.com/gomodule/redigo/redis"
)

// RedisPool Redisの接続プール
type RedisPool interface {
	NewRedisPool() redis.Conn
	Close()
}

type redisPool struct {
	redisConn redis.Conn
}

func NewRedis() RedisPool {
	return &redisPool{}
}

func (r *redisPool) NewRedisPool() redis.Conn {
	conn, err := redis.Dial("tcp", "172.17.0.2:6379")
	r.redisConn = conn
	if err != nil {
		panic(err)
	}
	return conn
}

func (r *redisPool) Close() {
	r.redisConn.Close()
}
