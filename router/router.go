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
	r.Use(middlewares.CheckAuth)

	r.HandleFunc("/api", controllers.HandleHome).Methods("GET")
	r.HandleFunc("/api/notes", controllers.GetNotes).Methods("GET")
	r.HandleFunc("/api/notes/{id}", controllers.GetNoteByID).Methods("GET")
	r.HandleFunc("/api/notes", controllers.CreateNote).Methods("POST")
	r.HandleFunc("/api/notes/{id}", controllers.UpdateNoteByID).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", controllers.DeleteNoteByID).Methods("DELETE")
	r.HandleFunc("/signup", controllers.SignUpUser).Methods("POST")
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	r.HandleFunc("/me", controllers.MeHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}
