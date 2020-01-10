package actions

import (
	"github.com/kataras/iris"
)

func Routes(app *iris.Application) {

	app.Get("/", home)
	app.Get("/db", migrateAndSeed)

	route := app.Party("/api").AllowMethods(iris.MethodOptions)
	{
		route.Use(middleware)
		route.Post("/register", register)
	}

}

func home(ctx iris.Context) {
	ctx.Writef("TBC API...")
}




