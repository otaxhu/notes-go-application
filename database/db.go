package database

import (
	"log"

	"github.com/otaxhu/notes-go-application/environment"
	"github.com/otaxhu/notes-go-application/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = &gorm.DB{}

var err error

func ConnectAndAutoMigrate() {

	// IMPORTANT! no se debe usar los dos puntos antes del "="
	// ya que se estaria creando otra instancia de &gorm.DB{}
	DB, err = gorm.Open(mysql.Open(environment.DSN), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return
	}
	if err := DB.Migrator().DropTable(&models.User{}, &models.Note{}); err != nil {
		log.Println(err.Error())
		return
	}
	if err := DB.AutoMigrate(&models.User{}, &models.Note{}); err != nil {
		log.Println(err.Error())
		return
	}
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := DB.First(&user, "email = ?", email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func FindUserByID(id string) (*models.User, error) {
	var user models.User
	if err := DB.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
