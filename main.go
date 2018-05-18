package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	userName := "root"
	userPassword := "rootpassword"
	dbHost := "rdb"
	dbPort := "3306"
	dbName := "firefly"
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", userName, userPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
