/**
 ******************************************************************************
 * File Name          : 查询当前网站访问人数
 * Author             : 张宇恺
 * Description        : 查 redis key = record
 ******************************************************************************
 */

package controller

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/kataras/iris/v12"
	"iris/src/redis-config"
)

type RecordController struct{}

func (rc *RecordController) Get() interface{} {
	isKey, err := redis.Bool(redisConfig.Redis.Do("EXISTS", "record"))
	if err != nil {
		panic(err)
	}

	// 如果没有 record 值，建立它
	if isKey != true {
		_, err := redisConfig.Redis.Do("set", "record", "1")
		if err != nil {
			fmt.Println("redis set error")
		}
		return "1"
	} else {
		var err error
		var record interface{}

		_, err = redisConfig.Redis.Do("incr", "record")
		if err != nil {
			fmt.Println("redis set error")
		}
		record, err = redisConfig.Redis.Do("get", "record")
		if err != nil {
			fmt.Println("redis get error")
		}

		return record
	}
}

func (rc *RecordController) Options(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Authorization")
	ctx.Header("Access-Control-Max-Age", "1")
}
