package main

import (
	"github.com/jffhtchr/record_project/database"

	"github.com/jffhtchr/record_project/cmd/server"
	_ "github.com/lib/pq"
)

func main() {

	// Connect Database
	database.ConnectDatabase()

	// Start Server
	server.StartServer(database.DB.Db)
}
