package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/otaxhu/notes-go-application/controllers"
	"github.com/otaxhu/notes-go-application/middlewares"
)

func InitializeRouter() {
	r := mux.NewRouter()

	r.Use(middlewares.Logging)

	r.HandleFunc("/api", controllers.HandleHome).Methods("GET")
	r.HandleFunc("/api/notes", controllers.GetNotes).Methods("GET")
	r.HandleFunc("/api/notes/{id}", controllers.GetNoteByID).Methods("GET")
	r.HandleFunc("/api/notes/create", controllers.CreateNote).Methods("POST")
	r.HandleFunc("/api/notes/{id}/update", controllers.UpdateNoteByID).Methods("PUT")
	r.HandleFunc("/api/notes/{id}/delete", controllers.DeleteNoteByID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
