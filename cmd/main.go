package main

import (
	"log"
	_transactionHandler "main/handlers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {

	// dbHost := viper.GetString(`database.host`)
	// dbPort := viper.GetString(`database.port`)
	// dbUser := viper.GetString(`database.user`)
	// dbPass := viper.GetString(`database.pass`)
	// dbName := viper.GetString(`database.name`)
	// dbconnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	// val := url.Values{}
	// val.Add("parseTime", "true")
	// val.Add("loc", "Asia/Bangkok")
	// dsn := fmt.Sprintf("%s?%s", dbconnection, val.Encode())
	// dbConn, err := sql.Open(`mysql`, dsn)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = dbConn.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("connect to database successfully")

	// defer func() {
	// 	err := dbConn.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	e := echo.New()
	_transactionHandler.NewTransactionHandler(e)
	log.Fatal(e.Start("0.0.0.0:8090"))
}
