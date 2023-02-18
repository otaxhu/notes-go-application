package database

import (
	"fmt"

	"github.com/otaxhu/notes-go-application/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var NotesDB = &gorm.DB{}

var err error

func ConnectAndAutoMigrate() {

	// IMPORTANT! no se debe usar los dos puntos antes del "="
	// ya que se estaria creando otra instancia de &gorm.DB{}
	NotesDB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	NotesDB.AutoMigrate(&models.Note{})
}
