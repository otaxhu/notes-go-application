package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/otaxhu/notes-go-application/database"
	"github.com/otaxhu/notes-go-application/environment"
	"github.com/otaxhu/notes-go-application/models"

	"golang.org/x/crypto/bcrypt"
)

type SignUpResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var loginUser models.User
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := database.FindUserByEmail(loginUser.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "email or password are incorrect", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err != nil {
		http.Error(w, "email or password are incorrect", http.StatusUnauthorized)
		return
	}

	claims := models.AppClaims{
		UserID: loginUser.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(environment.JWTSecret))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := LoginResponse{
		Token: tokenString,
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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
		http.Error(w, "email already registered", http.StatusBadRequest)
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

	response := SignUpResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}

	// Devolver el nuevo usuario como respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
