package models

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	DebajoPesoIdeal = -1
	PesoIdeal       = 0
	SobrePeso       = 1
)

type Persona struct {
	nombre string
	edad   int
	nss    string
	sexo   string
	peso   int
	altura float32
}

func NewPersona(nombre string, edad int, sexo string, peso int, altura float32) *Persona {
	if sexo != "M" && sexo != "H" {
		sexo = "H"
	}

	p := &Persona{
		nombre: nombre,
		edad:   edad,
		sexo:   sexo,
		peso:   peso,
		altura: altura,
	}

	p.nss = generaNSS()

	return p
}

func (p *Persona) CalcularIMC() int {

	result := float32(p.peso) / (p.altura * p.altura)

	if p.sexo == "H" {
		if result < 20 {
			return DebajoPesoIdeal
		} else if result >= 20 && result <= 25 {
			return PesoIdeal
		} else {
			return SobrePeso
		}
	} else {
		if result < 19 {
			return DebajoPesoIdeal
		} else if result >= 19 && result <= 24 {
			return PesoIdeal
		} else {
			return SobrePeso
		}
	}
}

func (p *Persona) EsMayorDeEdad() bool {
	if p.edad >= 18 {
		return true
	}
	return false
}

func (p *Persona) ToString() string {
	return fmt.Sprintf("Nombre:%s Edad:%d NSS:%s Sexo:%s Peso:%d Altura:%f", p.nombre, p.edad, p.nss, p.sexo, p.peso, p.altura)
}

func generaNSS() string {
	n := 8
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var bytes = make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

//Nombre metodo Getter de la propiedad nombre
func (p *Persona) Nombre() string {
	return p.nombre
}

func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

//Edad metodo Getter de la propiedad edad
func (p *Persona) Edad() int {
	return p.edad
}

func (p *Persona) SetEdad(edad int) {
	p.edad = edad
}

//NSS metodo Getter de la propiedad nss
func (p *Persona) NSS() string {
	return p.nss
}

//Sexo metodo Getter de la propiedad sexo
func (p *Persona) Sexo() string {
	return p.sexo
}

func (p *Persona) SetSexo(sexo string) {
	p.sexo = sexo
}

//Peso metodo Getter de la propiedad peso
func (p *Persona) Peso() int {
	return p.peso
}

func (p *Persona) SetPeso(peso int) {
	p.peso = peso
}

//Altura metodo Getter de la propiedad altura
func (p *Persona) Altura() float32 {
	return p.altura
}

func (p *Persona) SetAltura(altura float32) {
	p.altura = altura
}
