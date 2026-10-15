package main

import "fmt"

func main() {
	
	var cantidad int
	fmt.Print("¿Cuántas palabras tiene su texto? ")
	fmt.Scan(&cantidad)

	fmt.Println("Escriba su texto (en minúsculas y sin signos), separando las palabras con espacios:")

	conteo := make(map[string]int)
	for i := 0; i < cantidad; i++ {
		var palabra string
		fmt.Scan(&palabra)
		conteo[palabra]++
	}

	fmt.Println("\nConteo de palabras:")
	for palabra, veces := range conteo {
		fmt.Println(palabra, ":", veces)
	}

	masRepetida := ""
	maximo := 0
	for palabra, veces := range conteo {
		if veces > maximo {
			maximo = veces
			masRepetida = palabra
		}
	}

	fmt.Println("\nLa palabra que más se repite es:", masRepetida, "(", maximo, "veces )")
}