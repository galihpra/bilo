package main

import (
	"bilo/config"
	"bilo/helper/encrypt"
	"bilo/routes"
	"bilo/utils/database"

	uh "bilo/features/users/handler"
	ur "bilo/features/users/repository"
	us "bilo/features/users/service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var dbConfig = new(config.DatabaseMysql)
	if err := dbConfig.LoadFromEnv(); err != nil {
		panic(err)
	}

	dbConnection, err := database.MysqlInit(*dbConfig)
	if err != nil {
		panic(err)
	}

	if err := database.MysqlMigrate(dbConnection); err != nil {
		panic(err)
	}

	enc := encrypt.New()
	userRepository := ur.NewUserRepository(dbConnection)
	userService := us.New(userRepository, enc)
	userHandler := uh.NewUserHandler(userService)

	routes.InitRoute(
		e,
		userHandler,
	)

	e.Logger.Fatal(e.Start(":8000"))
}
