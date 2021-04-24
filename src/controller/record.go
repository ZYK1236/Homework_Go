/**
 ******************************************************************************
 * File Name          : 查询学生对应老师 Controller
 * Author             : 张宇恺
 * Description        : 根据传入的 stuno 去查对应的老师
 ******************************************************************************
 */

package controller

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
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
