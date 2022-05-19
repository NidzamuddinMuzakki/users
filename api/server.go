package api

import (
	"database/sql"

	"nidzamTest.com/app"
	"nidzamTest.com/controller"
	"nidzamTest.com/repository"
	"nidzamTest.com/service"
)

var (
	db             *sql.DB                   = app.Init()
	UserRepository repository.UserRepository = repository.NewUserRepository()
	UserService    service.UserService       = service.NewUserService(UserRepository, db)
	JwtService     service.JWTService        = service.NewJWTService()
	UserController controller.UserController = controller.NewUserController(UserService, JwtService)
)

func Run() {

	// customvalidator.CustomValidatorCOA(validate, transid, transen)
	defer db.Close()

	r := app.InitRouter(UserController)
	r.Start(":80")
}
