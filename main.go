package main

import (
	"impinj-server/actions"
	"log"
	"os"

	"github.com/iris-contrib/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Start iris app
	app := iris.New()

	// Config CORS
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	app.UseGlobal(crs)

	// Engage routers
	actions.Routes(app)

	// Run Server
	_ = app.Run(
		iris.Addr(os.Getenv("RUN_ON")),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
