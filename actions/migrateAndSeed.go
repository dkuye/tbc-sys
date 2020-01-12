package actions

import (
	"fmt"
	"github.com/dkuye/database"
	"github.com/dkuye/random"
	"github.com/kataras/iris"
	"os"
	"tbc-sys/models"
)

func migrateAndSeed(ctx iris.Context) {
	env := os.Getenv("APP_ENV")
	if env != "development" {
		data := map[string]string{"status": "failed", "message": "Not right environment...",}
		ctx.JSON(data)
		return
	}

	db := database.Connect()
	defer db.Close()

	tables := []interface{}{
		&models.APIBearer{},
		&models.PasswordReset{},
		&models.Payment{},
		&models.Program{},
		&models.ProgramRegistration{},
		&models.Role{},
		&models.User{},
	}

	for _, table := range tables {
		db.DropTableIfExists(table) // Drop table if exists
		db.AutoMigrate(table)       // Create table
	}
	fmt.Println("Migration done.")

	// Seeder
	models.APIBearerSeeder()
	models.ProgramSeeder()
	models.RoleSeeder()
	models.UserSeeder()


	_, _ = ctx.Writef("Status: Success \n")
	_, _ = ctx.Writef("Message: Migration and seeding done!\n")
	_, _ = ctx.Writef("Reference: " + random.String{Upper:true, Number:true}.Gen(26) + "\n")
}






// OLD one
/*func MigrateAndSeed(ctx iris.Context) {
	env := os.Getenv("APP_ENV")
	if env != "development" {
		data := map[string]string{
			"status":  "failed",
			"message": "Not right environment...",
		}
		ctx.JSON(data)
		return
	}

	db := database.Connect()
	defer db.Close()

	// Drop tables
	db.DropTableIfExists(
		&models.APIBearer{},
		&models.Company{},
		&models.Issue{},
		&models.LogAuth{},
		&models.LogFailAuth{},
		&models.PasswordReset{},
		&models.Question{},
		&models.QuestionResponse{},
		&models.Role{},
		&models.Service{},
		&models.State{},
		&models.TestCenter{},
		&models.Ticket{},
		&models.TicketAttachment{},
		&models.TicketCategory{},
		&models.TicketReassignment{},
		&models.TicketStatus{},
		&models.TicketStatusChange{},
		&models.TicketThread{},
		&models.User{},
		&models.ValidateEmail{},

		//Relation tables
		&models.TestCenterUser{},
		&models.ServiceUser{},
	)
	fmt.Println("Delete all tables")

	// Migrate tables
	db.AutoMigrate(
		&models.APIBearer{},
		&models.Company{},
		&models.Issue{},
		&models.LogAuth{},
		&models.LogFailAuth{},
		&models.PasswordReset{},
		&models.Question{},
		&models.QuestionResponse{},
		&models.Role{},
		&models.Service{},
		&models.State{},
		&models.TestCenter{},
		&models.Ticket{},
		&models.TicketAttachment{},
		&models.TicketCategory{},
		&models.TicketReassignment{},
		&models.TicketStatus{},
		&models.TicketStatusChange{},
		&models.TicketThread{},
		&models.User{},
		&models.ValidateEmail{},
	)
	fmt.Println("Install all tables")

	// Seeders
	// Important first
	seeder.RolesTableSeeder()
	fmt.Println("Call RolesTableSeeder() ")
	seeder.TicketStatusTableSeeder()
	fmt.Println("Call TicketStatusTableSeeder()")
	seeder.StatesTableSeeder()
	fmt.Println("Call StatesTableSeeder() ")
	seeder.APIBearerTableSeeder()
	fmt.Println("Call APIBearerTableSeeder() ")
	seeder.CompaniesTableSeeder()
	fmt.Println("Call CompaniesTableSeeder() ")

	//seeder.QuestionRespondsTableSeeder() 	;fmt.Println("Call QuestionRespondsTableSeeder() ")
	seeder.QuestionsTableSeeder()
	fmt.Println("Call QuestionsTableSeeder() ")
	//seeder.TicketAttachmentsTableSeeder() ;fmt.Println("Call TicketAttachmentsTableSeeder()")
	seeder.TicketsTableSeeder()
	fmt.Println("Call TicketsTableSeeder()")
	seeder.TicketThreadsTableSeeder()
	fmt.Println("Call TicketThreadsTableSeeder()")

	seeder.UsersTableSeeder()
	fmt.Println("Call UsersTableSeeder()")
	// Dependent on users table
	seeder.ServicesTableSeeder()
	fmt.Println("Call ServicesTableSeeder() ")
	seeder.TestCentersTableSeeder()
	fmt.Println("Call TestCentersTableSeeder()")

	data := map[string]string{
		"status": "success", //"message" : "DB Migrated And Seeded",
	}

	ctx.JSON(data)

}*/
