package main

import (
	"fmt"

	"github.com/Julianrt/examen_base/models"
)

func main() {

	var nombre string
	var edad int
	var sexo string
	var peso int
	var altura float32

	var err error

	fmt.Printf("Escribe tu nombre: ")
	_, err = fmt.Scan(&nombre)
	if err != nil {
		fmt.Println("Error al leer el nombre")
	}

	fmt.Printf("Escribe tu edad: ")
	_, err = fmt.Scan(&edad)
	if err != nil {
		fmt.Println("Error al leer la edad")
	}

	fmt.Printf("Escribe tu sexo (H si eres hombre o M si eres mujer) : ")
	_, err = fmt.Scan(&sexo)
	if err != nil {
		fmt.Println("Error al leer el sexo")
	}

	fmt.Printf("Escribe tu peso (en Kg) : ")
	_, err = fmt.Scan(&peso)
	if err != nil {
		fmt.Println("Error al leer el peso")
	}

	fmt.Printf("Escribe tu altura (en metros) : ")
	_, err = fmt.Scan(&altura)
	if err != nil {
		fmt.Println("Error al leer la altura")
	}

	persona := models.NewPersona(nombre, edad, sexo, peso, altura)

	res := persona.CalcularIMC()
	if res == models.DebajoPesoIdeal {
		fmt.Println("Estas debajo de tu peso ideal")
	} else if res == models.PesoIdeal {
		fmt.Println("Estas en tu peso ideal")
	} else if res == models.SobrePeso {
		fmt.Println("Tienes sobre peso")
	}

	if persona.EsMayorDeEdad() {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	fmt.Println(persona.ToString())
}
