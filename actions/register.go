package actions

import (
	"cg-pkg/database"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"io/ioutil"
	"log"
	"tbc-sys/models"
)

func register(ctx iris.Context) {
	b, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Println(err)
		return
	}
	var reg models.User
	err = json.Unmarshal(b, &reg)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	db := database.Connect()
	defer db.Close()

	status := reg.Create(db)

	if status == false {
		ctx.JSON(iris.Map{ "status":"failed", "message":"Got here.", })
		return
	}

	ctx.JSON(iris.Map{
		"status":"success",
		"message":"Registration completed.",
	})
}
