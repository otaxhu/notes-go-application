package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/otaxhu/notes-go-application/database"
	"github.com/otaxhu/notes-go-application/models"
	"gorm.io/gorm"
)

// /////////////////////////////////////////////////////////////////////////
// TODAS LA FUNCIONES LAS HICE ANTES DE CREAR EL SISTEMA DE AUTENTICACION
// ASI QUE PUEDE QUE NO TENGAN MUCHO SENTIDO
// ////////////////////////////////////////////////////////////////////////
func GetNotes(w http.ResponseWriter, r *http.Request) {
	var notes []models.Note
	database.DB.Find(&notes)
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
	// TODO: el user pueda crear una nota y luego en la funcion de GetNotes
	//       reciba las notas que coincida el UserID de la nota con el ID del usuario

	// 1- Suponiendo que el usuario esta autenticado con su JWT en el Authorization Header
	//    procede a enviar un request el cual va a contener una nota con Title y Description

	// 2- Se procede a crear una instancia de models.Note con el Title, Description.
	//    el UserID de models.Note va a ser el id del usuario, y el ID de la nota va a ser
	//    generado con uuid

	// 3- Se guarda en la base de datos

	// 4- Se Codifica en la respuesta el ID de la nota, el Title y la Description
	//    para saber que funciono
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
