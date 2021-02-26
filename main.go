package main

import (
	"github.com/joho/godotenv"
	"github.com/vipul-08/cassandra-api/app"
	"github.com/vipul-08/cassandra-api/database"
	"os"
)

func main() {
	if os.Getenv("ENVIRONMENT") != "prod" {
		err := godotenv.Load()
		if err != nil {
			panic("Unable to load .env file")
		}
	}
	database.ConnectDatabase()
	app.StartRoutes()
}
