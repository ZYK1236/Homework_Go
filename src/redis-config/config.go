package redisConfig

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var Redis redis.Conn

func Init() {
	var err error
	Redis, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis start error")
		panic(err)
	}
	fmt.Println("redis 连接成功 ✅ ----- >", Redis)
}
