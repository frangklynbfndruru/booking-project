package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB        *gorm.DB
	Router    *mux.Router
	AppConfig *AppConfig
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
	AppURL  string
}

type DbConfig struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
}

// function jika argumen dalam command line kosong
func (server *Server) RunDefaultPort(address string) {
	fmt.Printf("Listening to port %s", address)

	log.Fatal(http.ListenAndServe(address, server.Router))
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DbConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

}

func (server *Server) initializeDB(dbConfig DbConfig) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/jakarta",
		dbConfig.DbHost, dbConfig.DbUser, dbConfig.DbPassword, dbConfig.DbName, dbConfig.DbPort)

	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}

func initCommands() {

}
