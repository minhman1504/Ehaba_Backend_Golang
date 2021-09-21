package main

import (
	"ehaba_backend_golang/db"
	"ehaba_backend_golang/handler"
	"ehaba_backend_golang/router"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {

	fmt.Println("Starting...")

	m := db.MongoDB{
		DBURL:        "mongodb://localhost:27017",
		DatabaseName: "test_database",
	}

	m.Connect()
	defer m.Disconnect()

	userhandler := handler.UserHandler{
		MongoClient: m,
	}

	e := echo.New()

	api := router.API{
		Echo:        e,
		UserHandler: userhandler,
	}
	api.SetUpRouter()

	e.Logger.Fatal(e.Start(":3000"))
}
