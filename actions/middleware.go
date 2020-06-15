package actions

import (
	"fmt"
	"github.com/dkuye/database"
	"github.com/kataras/iris/v12"
	"impinj-server/models"
	"strings"
)

// Middleware is
func middleware(ctx iris.Context) {
	authorization := ctx.GetHeader("authorization")
	fmt.Println(authorization)
	if authorization == "" {
		ctx.JSON(iris.Map{"status": "failed", "message": "No Authorization header was found"})
		return
	}
	if strings.Contains(authorization, "Bearer ") == false {
		ctx.JSON(iris.Map{"status": "failed", "message": "Invalid Authorization in your request header."})
		return
	}
	key := authorization[7:len(authorization)]
	db := database.Connect()
	defer db.Close()
	apiBearer := models.APIBearer{}
	where := map[string]interface{}{
		"key":    key,
		"status": 1,
	}
	db.Where(where).First(&apiBearer)
	if apiBearer.Key == "" {
		ctx.JSON(iris.Map{"status": "failed", "message": "Invalid token."})
		return
	}
	ctx.Next()
}
