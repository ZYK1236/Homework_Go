package main

import (
	"fmt"
	"iris/src/controller"
	"iris/src/database"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()

	// 跨域
	app.Use(Cors)

	// 路由组 /student
	mvc.Configure(app.Party("/student"), func(a *mvc.Application) {
		a.Handle(new(controller.StudentController))
		a.Handle(new(controller.CourseController))
		a.Handle(new(controller.TeacherController))
	})

	// 捕获 404 错误
	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)

	// 初始化数据库
	database.Init()

	app.Run(iris.Addr(":8080"))
}

func notFoundHandler(ctx iris.Context) {
	path := ctx.Path()
	fmt.Println("path:", path, "----> 404 not found ❌")
	ctx.WriteString("404 not found")
}

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}

// type MyController struct{}

// // 自动识别你的 GET 请求
// func (mc *MyController) Get() string {
// 	fmt.Println("get")
// 	return "get"
// }

// func (mc *MyController) GetInfo() mvc.Result {
// 	fmt.Print("getinfo")
// 	// 返回自定义类型 Object
// 	// json 格式
// 	return mvc.Response{
// 		Object: model.GetModel(true, map[string]interface{}{
// 			"name": "zyk",
// 			"age":  18,
// 		}),
// 	}
// }
