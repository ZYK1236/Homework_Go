package main

import (
	"fmt"
	"iris/src/database"
	"iris/src/model"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	mvc.New(app).Handle(new(MyController))

	// 初始化数据库
	database.Init()

	app.Run(iris.Addr(":8080"))
}

type MyController struct{}

// 自动识别你的 GET 请求
func (mc *MyController) Get() string {
	fmt.Println("get")
	return "get"
}

func (mc *MyController) GetInfo() mvc.Result {
	fmt.Print("getinfo")
	// 返回自定义类型 Object
	// json 格式
	return mvc.Response{
		Object: model.GetModel(true, map[string]interface{}{
			"name": "zyk",
			"age":  18,
		}),
	}
}
