package database

import (
	"fmt"

	"github.com/otaxhu/notes-go-application/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var NotesDB = &gorm.DB{}

func ConnectAndAutoMigrate() {
	fmt.Println(NotesDB)
	fmt.Println(DSN)
	NotesDB, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	NotesDB.AutoMigrate(&models.Note{})
	fmt.Println(NotesDB)
}
