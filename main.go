package main

import (
	"crud-api-rest-go/commons"
	"crud-api-rest-go/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()
	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutando en el puerto 9000")
	log.Println(server.ListenAndServe())

}
