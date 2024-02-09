package main

import (
	"bilo/config"
	"bilo/helper/encrypt"
	"bilo/routes"
	"bilo/utils/database"

	uh "bilo/features/users/handler"
	ur "bilo/features/users/repository"
	us "bilo/features/users/service"

	ph "bilo/features/products/handler"
	pr "bilo/features/products/repository"
	ps "bilo/features/products/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
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

	var jwtConfig = new(config.JWT)
	if err := jwtConfig.LoadFromEnv(); err != nil {
		panic(err)
	}

	enc := encrypt.New()
	userRepository := ur.NewUserRepository(dbConnection)
	userService := us.New(userRepository, enc)
	userHandler := uh.NewUserHandler(userService, *jwtConfig)

	productRepository := pr.NewProductRepository(dbConnection)
	productService := ps.NewProductService(productRepository)
	productHandler := ph.NewProductHandler(productService, *jwtConfig)

	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	route := routes.Routes{
		JWTKey:         jwtConfig.Secret,
		Server:         app,
		UserHandler:    userHandler,
		ProductHandler: productHandler,
	}

	route.InitRouter()

	app.Logger.Fatal(app.Start(":8000"))
}
