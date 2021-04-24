package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris/src/controller"
	"iris/src/database"
	"iris/src/redis-config"
	logMsg "iris/src/utils"
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
		a.Handle(new(controller.UploadController))
	})

	// 路由组 /find
	mvc.Configure(app.Party("/find"), func(a *mvc.Application) {
		a.Handle(new(controller.FindCourseController))
		a.Handle(new(controller.FindTeacherController))
	})

	// 路由组 /record
	mvc.Configure(app.Party("/record"), func(a *mvc.Application) {
		a.Handle(new(controller.RecordController))
	})

	// 捕获 404 错误
	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)

	// 初始化数据库
	database.Init()
	// 初始化 redis
	redisConfig.Init()

	err := app.Run(iris.Addr(":8080"))
	if err != nil {
		fmt.Println("app.Run error...")
		return
	}
}

func notFoundHandler(ctx iris.Context) {
	logMsg.LogErrorMsg(ctx.Path(), ctx.Method())
	_, err := ctx.WriteString("404 not found")
	if err != nil {
		fmt.Println("ctx.WriteString error")
	}
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
