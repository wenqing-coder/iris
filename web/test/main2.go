package main

import (
	"iris/bootstrap"

	"iris/web/middleware/identity"

	"iris/web/routes"
)

func main() {
	app := bootstrap.New("Superstar database", "一凡Sir")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	app.Listen(":8080")
}
