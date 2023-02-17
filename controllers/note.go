package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/otaxhu/notes-go-application/database"
	"github.com/otaxhu/notes-go-application/models"
)

/*
func GetNotes(w http.ResponseWriter, r *http.Request) {
	var notes []models.Note
	database.NotesDB.Find(&notes)
	err := json.NewEncoder(w).Encode(notes)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
}
*/
/*
func GetNoteByID(w http.ResponseWriter, r *http.Request) {
*/
/*
	idClean, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if idClean <= 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
*/
/*
	var note models.Note
	params := mux.Vars(r)
	err := database.NotesDB.First(&note, params["id"]).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
}
*/

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
