package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/otaxhu/notes-go-application/controllers"
)

func InitializeRouter() {
	r := mux.NewRouter()

	//TODO: no funciona ningun tipo de metodo http, solo funciona el metodo GET en /api

	r.HandleFunc("/api", controllers.HandleHome).Methods("GET")
	//r.HandleFunc("/api/notes", controllers.GetNotes).Methods("GET")
	//r.HandleFunc("/api/notes/{id}", controllers.GetNoteByID).Methods("GET")
	//r.HandleFunc("/api/notes/create", controllers.CreateNote).Methods("POST")
	http.ListenAndServe(":3000", r)
}
