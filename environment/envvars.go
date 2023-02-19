package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var envPath = "./environment/.env"

// Cargar del archivo llamado ".env" que esta adentro de la carpeta "./environment"
var _ = godotenv.Load(envPath)

var (
	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))
)

var JWTSecret = os.Getenv("jwt_secret")
