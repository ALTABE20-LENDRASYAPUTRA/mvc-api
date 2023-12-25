package main

import (
	"fmt"
	"mpc-api/config"
	"mpc-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("running")
	config.InitDB()
	config.InitialMigration()

	e := echo.New()
	routes.InitRoute(e)

	//start server and port
	e.Logger.Fatal(e.Start(":8000"))
}
