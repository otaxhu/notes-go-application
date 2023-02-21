package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/otaxhu/notes-go-application/authorization"
	"github.com/otaxhu/notes-go-application/database"
	"github.com/otaxhu/notes-go-application/models"
	"gorm.io/gorm"
)

type NoteResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	claims, err := authorization.GetClaims(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
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
}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	claims, err := authorization.GetClaims(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	params := mux.Vars(r)
	dbNote, err := database.FindNoteByID(claims.UserID, params["id"])
	if err == gorm.ErrRecordNotFound {
		http.Error(w, "note not found, or you have no authorization to see this note", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := NoteResponse{
		ID:          dbNote.ID,
		Title:       dbNote.Title,
		Description: dbNote.Description,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	claims, err := authorization.GetClaims(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
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
}

func UpdateNoteByID(w http.ResponseWriter, r *http.Request) {
	claims, err := authorization.GetClaims(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	params := mux.Vars(r)
	dbNote, err := database.FindNoteByID(claims.UserID, params["id"])
	if err == gorm.ErrRecordNotFound {
		http.Error(w, "note not found, or you have no authorization to update this note", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&dbNote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Save(&dbNote).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := NoteResponse{
		ID:          dbNote.ID,
		Title:       dbNote.Title,
		Description: dbNote.Description,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteNoteByID(w http.ResponseWriter, r *http.Request) {
	claims, err := authorization.GetClaims(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	params := mux.Vars(r)
	dbNote, err := database.FindNoteByID(claims.UserID, params["id"])
	if err == gorm.ErrRecordNotFound {
		http.Error(w, "note not found, or you have no the authorization to delete this note", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := database.DB.Delete(&dbNote).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
