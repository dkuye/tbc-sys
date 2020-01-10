package main

import (
	"log"
	"os"
	"tbc-sys/actions"

	"github.com/iris-contrib/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kataras/iris"
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
		AllowedOrigins:   []string{"http://localhost:8080", "http://tbc.deleosunmakinde.org", "http://app.tbcoutofzion.org", "http://api.tbcoutofzion.org", }, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Cache-Control", "X-File-Name", "X-Requested-With", "X-File-Name", "Content-Type", "Authorization", "Set-Cookie", "Cookie"},
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
