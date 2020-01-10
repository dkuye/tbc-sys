package main

import "github.com/kataras/iris"

func routes(app *iris.Application) {

	app.Get("/", home)
	app.Get("/db", migrateAndSeed)

	route := app.Party("/api").AllowMethods(iris.MethodOptions)
	{
		//route.Use(middleware)
		route.Get("/hello", hello)
	}

}

func home(ctx iris.Context) {
	ctx.Writef("TBC API...")
}

func hello(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":"success",
		"message":"Hello world.",
	})
}
