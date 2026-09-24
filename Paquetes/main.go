package main

import (
	"fmt"
	"practica/practica1/saludo"
)

func main() {
	fmt.Println("***Bienvenido a la clase de paquetes 🤩😎***")
	mensaje := saludo.Saludar("Juan")
	fmt.Println(mensaje)
}
