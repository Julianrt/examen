package main

import (
	"fmt"
	"net/http"

	"github.com/Julianrt/examen_rest/handlers"
	"github.com/gorilla/mux"
)

var portNum string = ":8080"

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/NutriNET/Cliente", handlers.GetClientes).Methods("GET")
	mux.HandleFunc("/NutriNET/Cliente", handlers.CrearCliente).Methods("POST")
	mux.HandleFunc("/NutriNET/Cliente/{id}", handlers.GetCliente).Methods("GET")
	mux.HandleFunc("/NutriNET/Cliente/{id}", handlers.ActualizarCliente).Methods("PUT")

	fmt.Println("Servidor corriendo por el puerto", portNum)
	_ = http.ListenAndServe(portNum, mux)
}
