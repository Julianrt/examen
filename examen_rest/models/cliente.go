package models

import (
	"errors"
	"time"
)

var clientes []Cliente

type Cliente struct {
	ClienteId          int       `json:"cliente_id"`
	NombreUsuario      string    `json:"nombre_usuario"`
	Password           string    `json:"password"`
	Nombre             string    `json:"nombre"`
	Apellidos          string    `json:"apellidos"`
	CorreoElectronico  string    `json:"correo_electronico"`
	Edad               int       `json:"edad"`
	Estatura           float32   `json:"estatura"`
	Peso               float32   `json:"peso"`
	IMC                float32   `json:"imc"`
	GEB                float32   `json:"geb"`
	ETA                float32   `json:"eta"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}

func GetClientes() []Cliente {
	return clientes
}

func GetClienteByID(id int) (*Cliente, error) {
	for _, cliente := range clientes {
		if cliente.ClienteId == id {
			return &cliente, nil
		}
	}

	return nil, errors.New("Cliente no encontrado")
}

func (c *Cliente) Guardar() error {

	for _, cliente := range clientes {
		if cliente.NombreUsuario == c.NombreUsuario {
			return errors.New("El nombre de usuario ya está registrado")
		}
		if cliente.CorreoElectronico == c.CorreoElectronico {
			return errors.New("El correo electrónico ya está registrado")
		}
	}

	c.FechaCreacion = time.Now()

	if len(clientes) < 1 {
		c.ClienteId = 1
	} else {
		c.ClienteId = clientes[len(clientes)-1].ClienteId + 1
	}

	clientes = append(clientes, *c)

	return nil
}

func (c *Cliente) Actualizar() error {

	for i := 0; i < len(clientes); i++ {

		if clientes[i].ClienteId == c.ClienteId {
			c.FechaActualizacion = time.Now()
			clientes[i] = *c
			return nil
		}

	}

	return errors.New("No se pudo actualizar el cliente")
}
