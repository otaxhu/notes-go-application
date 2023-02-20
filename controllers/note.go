package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/otaxhu/notes-go-application/database"
	"github.com/otaxhu/notes-go-application/environment"
	"github.com/otaxhu/notes-go-application/models"
	"gorm.io/gorm"
)

type NoteResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(environment.JWTSecret), nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		var notes []models.Note
		var response []NoteResponse
		pageSize := 2
		pageNumber, _ := strconv.Atoi(strings.TrimSpace(r.URL.Query().Get("page")))
		if err := database.DB.Offset((pageNumber-1)*pageSize).Limit(pageSize).Find(&notes, "user_id = ?", claims.UserID).Order("created_at DESC").Scan(&response).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}
}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {

	var note models.Note
	params := mux.Vars(r)
	id := params["id"]

	if err := database.DB.First(&note, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(environment.JWTSecret), nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		var note models.Note
		if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		dbUser, err := database.FindUserByID(claims.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if dbUser == nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		id := uuid.New().String()
		note.ID = id
		note.UserID = claims.UserID
		if err := database.DB.Create(&note).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := NoteResponse{
			ID:          note.ID,
			Title:       note.Title,
			Description: note.Description,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}
}

func UpdateNoteByID(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la nota a actualizar
	params := mux.Vars(r)
	id := params["id"]

	// Obtener la nota de la base de datos
	var note models.Note
	if err := database.DB.First(&note, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Decodificar la nota actualizada de la solicitud
	var updatedNote models.Note
	if err := json.NewDecoder(r.Body).Decode(&updatedNote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Actualizar los campos de la nota
	if err := database.DB.Model(&note).Updates(updatedNote).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Codificar la nota actualizada como JSON y enviarla como respuesta
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func DeleteNoteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var deletedNote models.Note
	if err := database.DB.First(&deletedNote, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := database.DB.Delete(&deletedNote).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
