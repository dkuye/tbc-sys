package actions

import (
	"encoding/json"
	"github.com/dkuye/database"
	"github.com/dkuye/helper"
	"github.com/kataras/iris"
	"io/ioutil"
	"os"
	"tbc-sys/models"
)

func paymentPrep(ctx iris.Context)  {
	email := ctx.FormValue("email")

	db := database.Connect()
	defer db.Close()

	var user models.User
	db.Where(models.User{Email:email}).First(&user)

	if user.Email == "" {
		ctx.JSON(iris.Map{
			"status": "failed",
			"message": "Unable record with " + email,
		})
		return
	}

	//REG_AMOUNT

	ctx.JSON(iris.Map{
		"status": "success",
		"message": "Successful",
		"name" : user.FirstName + " " + user.LastName,
		"email" : user.Email ,
		"amount" : helper.StringToInt(os.Getenv("REG_AMOUNT")) * 100,
		"amountF" : "₦" + helper.FormatNumber( helper.StringToInt( os.Getenv("REG_AMOUNT") ) ),
	})

}

func paymentInfo(ctx iris.Context)  {
	email := ctx.FormValue("email")
	b, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		ctx.JSON(iris.Map{ "status":"failed", "message": err, })
		return
	}
	var payment models.Payment
	err = json.Unmarshal(b, &payment)
	if err != nil {
		ctx.JSON(iris.Map{ "status":"failed", "message": err, })
		return
	}

	db := database.Connect()
	defer db.Close()

	var user models.User
	db.Where(models.User{Email:email}).First(&user)

	var program models.Program
	db.Last(&program)

	db.Create(&models.Payment{
		Reference:   payment.Reference,
		Transaction: payment.Transaction,
		RedirectUrl: payment.RedirectUrl,
		Status:      payment.Status,
		Message:     payment.Message,
		UserID:      user.ID,
		ProgramID:   program.ID,
		Amount:      float64( helper.StringToInt(os.Getenv("REG_AMOUNT")) ),
		Uuid:        helper.Uuid(),
	})

	ctx.JSON(iris.Map{ "status": payment.Status, "message": payment.Message, })
}
