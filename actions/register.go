package actions

import (
	"encoding/json"
	"github.com/dkuye/database"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"tbc-sys/models"
)

func register(ctx iris.Context) {
	b, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		ctx.JSON(iris.Map{ "status":"failed", "message": err, })
		return
	}
	var reg models.User
	err = json.Unmarshal(b, &reg)
	if err != nil {
		ctx.JSON(iris.Map{ "status":"failed", "message": err, })
		return
	}

	db := database.Connect()
	defer db.Close()

	user := reg.Create(db)

	var program models.Program
	db.Last(&program)

	var existingRegistration models.ProgramRegistration
	db.Where(models.ProgramRegistration{ ProgramID:program.ID, UserID:user.ID }).First(&existingRegistration)
	if existingRegistration.ID != 0 {
		ctx.JSON(iris.Map{ "status":"failed", "message":"You are registered already.", })
		return
	}

	db.Create(&models.ProgramRegistration{ ProgramID:program.ID, UserID:user.ID })

	ctx.JSON(iris.Map{ "status":"success", "message":"Registration completed. Please make payment.",})
}
