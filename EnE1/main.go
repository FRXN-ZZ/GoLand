package main

import "fmt"

type Persona struct {
	nombre string
}

func (p Persona) GetNombre() string {
	return p.nombre
}

func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

func main() {
	persona := Persona{}

	persona.SetNombre("FRXN")
	fmt.Println("Su nombre es: ", persona.GetNombre())

	persona.SetNombre("Pauli")
	fmt.Println("Ahora, su nombre es: ", persona.GetNombre())
}