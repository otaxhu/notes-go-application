package main

import (
	"github.com/otaxhu/notes-go-application/database"
	"github.com/otaxhu/notes-go-application/router"
)

func main() {
	database.ConnectAndAutoMigrate()
	router.InitializeRouter()
}
