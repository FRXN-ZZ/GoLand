package main

import "fmt"


func obtenerGanador(votos map[string]int) string {
	masVotada := ""
	maxVotos := -1

	for actividad, cantidad := range votos {
		if cantidad > maxVotos {
			maxVotos = cantidad
			masVotada = actividad
		}
	}
	return masVotada
}

func main() {
	
	votos := map[string]int{
		"deportes":   0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	fmt.Println("Opciones: deportes, videojuegos, cine, musica")


	for i := 1; i <= 5; i++ {
		var voto string
		fmt.Printf("Ingrese el voto %d: ", i)
		fmt.Scanln(&voto)
		
		
		votos[voto]++ 
	}

	
	fmt.Println("\n--- RESULTADOS ---")
	for actividad, cantidad := range votos {
		fmt.Printf("%s: %d votos\n", actividad, cantidad)
	}

	
	ganador := obtenerGanador(votos)
	fmt.Printf("\nLa actividad con mayor número de votos es: %s\n", ganador)
}