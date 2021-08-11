package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Julianrt/examen_rest/models"
	"github.com/gorilla/mux"
)

func GetClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var resp models.Respuesta
	clientes := models.GetClientes()

	if len(clientes) < 1 {
		resp.CveError = -404
		resp.CveMensaje = "No se encontró clientes"
		resp.Data = []models.Cliente{}

		json.NewEncoder(w).Encode(&resp)
		return
	}

	resp.CveError = 0
	resp.CveMensaje = ""
	resp.Data = clientes

	json.NewEncoder(w).Encode(&resp)
}

func GetCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var resp models.Respuesta
	var err error

	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		resp.CveError = -404
		resp.CveMensaje = "No se encontró el cliente"

		json.NewEncoder(w).Encode(&resp)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.CveError = -400
		resp.CveMensaje = "Ingresa un id de cliente correcto"

		json.NewEncoder(w).Encode(&resp)
		return
	}
	if id < 1 {
		resp.CveError = -400
		resp.CveMensaje = "Ingresa un id de cliente correcto"

		json.NewEncoder(w).Encode(&resp)
		return
	}

	cliente, err := models.GetClienteByID(id)
	if err != nil {
		resp.CveError = -404
		resp.CveMensaje = "No se encontró el cliente"

		json.NewEncoder(w).Encode(&resp)
		return
	}

	resp.CveError = 0
	resp.CveMensaje = ""
	resp.Data = cliente

	json.NewEncoder(w).Encode(&resp)
}

func CrearCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var err error
	var cliente models.Cliente
	var resp models.Respuesta

	if err = json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		resp.CveError = -400
		resp.CveMensaje = "No se pudo leer el JSON"

		json.NewEncoder(w).Encode(&resp)
		return
	}

	if err = cliente.Guardar(); err != nil {
		resp.CveError = -400
		resp.CveMensaje = err.Error()

		json.NewEncoder(w).Encode(&resp)
		return
	}

	resp.CveError = 0
	resp.CveMensaje = ""
	resp.Data = cliente
	json.NewEncoder(w).Encode(&resp)
}

func ActualizarCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var err error
	var clienteNuevo models.Cliente
	var resp models.Respuesta

	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		resp.CveError = -400
		resp.CveMensaje = "Parámetro vacío"

		json.NewEncoder(w).Encode(&resp)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.CveError = -400
		resp.CveMensaje = "Ingresa un id de cliente correcto"

		json.NewEncoder(w).Encode(&resp)
		return
	}
	if id < 1 {
		resp.CveError = -400
		resp.CveMensaje = "Ingresa un id de cliente correcto"

		json.NewEncoder(w).Encode(&resp)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&clienteNuevo); err != nil {
		resp.CveError = -400
		resp.CveMensaje = "No se pudo leer el JSON"

		json.NewEncoder(w).Encode(&resp)
		return
	}

	cliente, err := models.GetClienteByID(id)
	if err != nil {
		resp.CveError = -404
		resp.CveMensaje = "No se encontró el cliente"

		json.NewEncoder(w).Encode(&resp)
		return
	}

	if clienteNuevo.Password != "" {
		cliente.Password = clienteNuevo.Password
	}

	if clienteNuevo.Nombre != "" {
		cliente.Nombre = clienteNuevo.Nombre
	}

	if clienteNuevo.Apellidos != "" {
		cliente.Apellidos = clienteNuevo.Apellidos
	}

	if clienteNuevo.Edad != 0 {
		cliente.Edad = clienteNuevo.Edad
	}

	if clienteNuevo.Estatura > 0.0 {
		cliente.Estatura = clienteNuevo.Estatura
	}

	if clienteNuevo.Peso > 0.0 {
		cliente.Peso = clienteNuevo.Peso
	}

	cliente.IMC = clienteNuevo.IMC
	cliente.GEB = clienteNuevo.GEB
	cliente.ETA = clienteNuevo.ETA

	if err = cliente.Actualizar(); err != nil {
		resp.CveError = -500
		resp.CveMensaje = err.Error()

		json.NewEncoder(w).Encode(&resp)
		return
	}

	resp.CveError = 0
	resp.CveMensaje = ""
	resp.Data = cliente
	json.NewEncoder(w).Encode(&resp)
}
