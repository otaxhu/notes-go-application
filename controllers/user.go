package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/otaxhu/notes-go-application/database"
	"github.com/otaxhu/notes-go-application/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud en un nuevo objeto de usuario
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newUser.Email == "" || newUser.Password == "" {
		http.Error(w, "el email y el password son campos obligatorios", http.StatusBadRequest)
		return
	}

	// Verificar si el email ya está registrado
	user, err := database.FindUserByEmail(newUser.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user != nil {
		http.Error(w, "already registered", http.StatusBadRequest)
		return
	}

	// Hash la contraseña del usuario
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newUser.Password = string(hashedPassword)

	id := uuid.New().String()

	newUser.ID = id

	// Crear el usuario en la base de datos
	if err := database.NotesDB.Create(&newUser).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Devolver el nuevo usuario como respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
