package learn

import (
	"fmt"
	"sync"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)


type Person struct {
	Name string `json:"name"`
	Pwd  int    `json:"pwd"`
}

func Learn() {
	app := iris.New()
	// 错误捕获 recover
	app.Use(recover.New())
	app.Use(logger.New())
	//输出html
	// 请求方式: GET
	// 访问地址: http://localhost:8080/welcome
	app.Handle("GET", "/welcome", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})
	//输出字符串
	// 类似于 app.Handle("GET", "/ping", [...])
	// 请求方式: GET
	// 请求地址: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong!!!")
	})
	//输出json
	// 请求方式: GET
	// 请求地址: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		// 获取 query
		param := ctx.URLParam("name")
		fmt.Println(param, "爬")
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	app.Post("/trypost", func(ctx iris.Context) {
		// 获取表单参数
		name := ctx.PostValue("name")
		pwd := ctx.PostValue("pwd")
		fmt.Println("name=", name, "pwd=", pwd)
		ctx.WriteString("success")
	})

	app.Post("/tryjson", func(ctx iris.Context) {
		var mutex sync.RWMutex
		// 获取 JSON 格式
		var c Person
		mutex.Lock()
		err := ctx.ReadJSON(&c)
		mutex.Unlock()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%+v", c)
		ctx.WriteString("终于搞好了")
	})

	// 多组路由
	app.PartyFunc("/users", func(users iris.Party) {
		// 处理localhost:8080/users 返回的视图
		users.Use(myAuthMiddlewareHandler, nextHandler)
		users.Get("/{id: int}/profile", handleProfile)
	})

	app.Run(iris.Addr(":8080")) //8080 监听端口
}

func myAuthMiddlewareHandler(ctx iris.Context) {
	ctx.WriteString("Authentication success")
	ctx.Next() //继续执行后续的handler
}

func nextHandler(ctx iris.Context) {
	ctx.Next()
}

func handleProfile(ctx iris.Context) {
	// 获取 param，即动态路由
	id := ctx.Params().Get("id")
	ctx.HTML(fmt.Sprintf("<h1>当前id: %s</h1>", id))
}
