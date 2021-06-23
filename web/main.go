package main

import (
	"iris/bootstrap"
	"iris/web/middleware/identity"
	"iris/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Superstardatabase", "wenqing")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
