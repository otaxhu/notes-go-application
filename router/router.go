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

	r.HandleFunc("/api", controllers.HandleHome).Methods(http.MethodGet)
	r.HandleFunc("/api/notes", controllers.GetNotes).Methods(http.MethodGet)
	r.HandleFunc("/api/notes/{id}", controllers.GetNoteByID).Methods(http.MethodGet)
	r.HandleFunc("/api/notes", controllers.CreateNote).Methods(http.MethodPost)
	r.HandleFunc("/api/notes/{id}", controllers.UpdateNoteByID).Methods(http.MethodPut)
	r.HandleFunc("/api/notes/{id}", controllers.DeleteNoteByID).Methods(http.MethodDelete)
	r.HandleFunc("/signup", controllers.SignUpUser).Methods(http.MethodPost)
	r.HandleFunc("/login", controllers.LoginUser).Methods(http.MethodPost)
	r.HandleFunc("/me", controllers.MeHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":3000", r))
}
