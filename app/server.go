package app

import (
	"flag"
	"log"
	"os"

	"github.com/frangklynbfndruru/booking-project/app/configuration"

	"github.com/joho/godotenv"
)

func Run() {
	var server = configuration.Server{}
	var appConfig = configuration.AppConfig{}
	var dbConfig = configuration.DbConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error OS loading .env file !", err)
	}

	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppEnv = os.Getenv("APP_ENV")
	appConfig.AppPort = os.Getenv("APP_PORT")

	dbConfig.DbHost = os.Getenv("DB_HOST")
	dbConfig.DbUser = os.Getenv("DB_USER")
	dbConfig.DbPassword = os.Getenv("DB_PASSWORD")
	dbConfig.DbName = os.Getenv("DB_NAME")
	dbConfig.DbPort = os.Getenv("DbPort")

	flag.Parse()       //untuk menerima command go run dari terminal
	// arg := flag.Arg(0) //mengambil argumen pertamadari command line. contoh `go run db:migrate`

	// if arg != "" {
	// 	server.InitCommands(appConfig, dbConfig)

	// } else {
		server.Initialize(appConfig, dbConfig)
		server.RunDefaultPort(":" + appConfig.AppPort)
	// }
}
