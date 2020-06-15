package actions

import (
	"github.com/kataras/iris/v12"
)

func Routes(app *iris.Application) {

	app.Get("/", home)
	app.Get("/db", migrateAndSeed)

	route := app.Party("/rfid").AllowMethods(iris.MethodOptions)
	{
		//route.Use(middleware)
		route.Post("/read", read)
	}

}

func home(ctx iris.Context) {
	ctx.Writef("IMPINJ API...")
}
