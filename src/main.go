package main

import (
	"fmt"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris/src/controller"
	"iris/src/database"
	"iris/src/redis-config"
	logMsg "iris/src/utils"
)

var mySecret = []byte("secret")

var jwtConfig = jwt.New(jwt.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	},

	Expiration: true,

	SigningMethod: jwt.SigningMethodHS256,
})

func main() {
	app := iris.New()

	// 跨域
	app.Use(Cors)

	// jwt 鉴权
	app.Use(Jwt)

	// 捕获 404 错误
	app.OnErrorCode(iris.StatusNotFound, NotFoundHandler)

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

	// 路由组 /login
	mvc.Configure(app.Party("/login"), func(a *mvc.Application) {
		a.Handle(new(controller.LoginController))
	})

	// 初始化数据库
	database.Init()

	// 初始化 redis
	redisConfig.Init()

	// 服务启动
	err := app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.Configuration{}))
	if err != nil {
		fmt.Println("app.Run error...")
		return
	}
}

func NotFoundHandler(ctx iris.Context) {
	fmt.Println("not found 404", ctx.Method())
	logMsg.LogErrorMsg(ctx.Path(), ctx.Method())
	_, err := ctx.WriteString("404 not found")
	if err != nil {
		fmt.Println("ctx.WriteString error")
	}
}

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Next()
}

func Jwt(ctx iris.Context) {
	if ctx.Path() == "/login" {
		ctx.Next()
		return
	}

	if err := jwtConfig.CheckJWT(ctx); err != nil {
		jwtConfig.Config.ErrorHandler(ctx, err)
		return
	}
	// 验证 token 通过，放行
	ctx.Next()
}
