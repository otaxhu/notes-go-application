package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var envPath = "./database/.env"

// Cargar del archivo llamado ".env" que esta adentro de la carpeta "./database"
var _ = godotenv.Load(envPath)

var (
	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))
)
