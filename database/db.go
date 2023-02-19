package database

import (
	"log"

	"github.com/otaxhu/notes-go-application/environment"
	"github.com/otaxhu/notes-go-application/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var NotesDB = &gorm.DB{}

var err error

func ConnectAndAutoMigrate() {

	// IMPORTANT! no se debe usar los dos puntos antes del "="
	// ya que se estaria creando otra instancia de &gorm.DB{}
	NotesDB, err = gorm.Open(mysql.Open(environment.DSN), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return
	}
	NotesDB.AutoMigrate(&models.Note{})
	NotesDB.AutoMigrate(&models.User{})
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := NotesDB.First(&user, "email = ?", email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func FindUserByID(id string) (*models.User, error) {
	var user models.User
	if err := NotesDB.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
