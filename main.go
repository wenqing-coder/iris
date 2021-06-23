package main

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12/mvc"

	"github.com/kataras/iris/v12/core/host"

	"github.com/kataras/iris/v12"
)

func Configurator(app *iris.Application) {
	counter := 0
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			counter++
		}
		app.ConfigureHost(func(su *host.Supervisor) {
			su.RegisterOnShutdown(func() {
				ticker.Stop()
				fmt.Println("server terminated")
			})
		})
	}()
	app.Get("/counter", func(ctx iris.Context) {
		ctx.Writef("counter value == %d", counter)
	})
}

func main() {
	app := iris.New()
	mvc.Configure(app.Party("/root"), myMVC)
	app.Run(iris.Addr(":8080"))
}

func myMVC(app *mvc.Application) {
	// app.Register(...)
	// app.Router.Use/UseGlobal/Done(...)
	app.Handle(new(MyController))
}

type MyController struct{}

func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done // 和你已知的任何标准 API  调用

	// 1-> 方法
	// 2-> 路径
	// 3-> 控制器函数的名称将被解析未一个处理程序 [ handler ]
	// 4-> 任何应该在 MyCustomHandler 之前运行的处理程序[ handlers ]
	b.Handle("GET", "/something/{id:long}", "MyCustomHandler")
}

// GET: http://localhost:8080/root
func (m *MyController) Get() string { return "Hey" }

// GET: http://localhost:8080/root/something/{id:long}
func (m *MyController) MyCustomHandler(id int64) string { return "MyCustomHandler says Hey" }
