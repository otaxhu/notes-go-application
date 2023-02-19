package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/otaxhu/notes-go-application/database"
	"github.com/otaxhu/notes-go-application/environment"
	"github.com/otaxhu/notes-go-application/models"

	"golang.org/x/crypto/bcrypt"
)

var tokenDuration = 48 * time.Hour

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

	if loginUser.Email == "" || loginUser.Password == "" {
		http.Error(w, "el email y la contraseña son campos obligatorios", http.StatusBadRequest)
		return
	}

	dbUser, err := database.FindUserByEmail(loginUser.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if dbUser == nil {
		http.Error(w, "email or password are incorrect", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password)); err != nil {
		http.Error(w, "email or password are incorrect", http.StatusUnauthorized)
		return
	}

	claims := models.AppClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenDuration).Unix(),
			Id:        dbUser.ID,
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

func SignUpUser(w http.ResponseWriter, r *http.Request) {
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

func MeHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(environment.JWTSecret), nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		user, err := database.FindUserByID(claims.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if user == nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
