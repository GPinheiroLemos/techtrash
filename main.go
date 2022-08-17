package main

import (
	"log"
	"net/http"
	"techTrash/connection"
	"techTrash/handlers"
	"techTrash/user"

	"github.com/gorilla/mux"
)

func main() {
	connection.MysqlConnect()
	router := mux.NewRouter()
	mux.CORSMethodMiddleware(router)
	router.HandleFunc("/lixeira", handlers.GetLixeira).Methods("GET")
	router.HandleFunc("/loglixeira", handlers.GetLog).Methods("GET")
	router.HandleFunc("/lixeira", handlers.PostLixeira).Methods("POST")
	router.HandleFunc("/loglixeira", handlers.PostLog).Methods("POST")
	router.HandleFunc("/cadastrarusuario", user.NewUser).Methods("POST")
	router.HandleFunc("/autenticarusuario", user.AuthUser).Methods("POST")
	log.Print("Running at port :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
