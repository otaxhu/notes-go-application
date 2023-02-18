package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/otaxhu/notes-go-application/database"
	"github.com/otaxhu/notes-go-application/models"
	"gorm.io/gorm"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	var notes []models.Note
	database.NotesDB.Find(&notes)
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {

	var note models.Note
	params := mux.Vars(r)
	id := params["id"]

	if err := database.NotesDB.First(&note, id).Error; err != nil {
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
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Crear la nota en la base de datos
	if err := database.NotesDB.Create(&note).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Codificar la nota como JSON y enviarla como respuesta
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func UpdateNoteByID(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la nota a actualizar
	params := mux.Vars(r)
	id := params["id"]

	// Obtener la nota de la base de datos
	var note models.Note
	if err := database.NotesDB.First(&note, "id = ?", id).Error; err != nil {
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
	if err := database.NotesDB.Model(&note).Updates(updatedNote).Error; err != nil {
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
	if err := database.NotesDB.First(&deletedNote, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := database.NotesDB.Delete(&deletedNote).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
