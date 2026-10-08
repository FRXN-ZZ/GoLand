package main

import "fmt"

func actividadMasVotada(votos map[string]int) string {
	ganador := ""
	maxVotos := -1

	for actividad, cantidad := range votos {
		if cantidad > maxVotos {
			maxVotos = cantidad
			ganador = actividad
		}
	}
	return ganador
}

func main() {

	votos := map[string]int{
		"deportes":   0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	fmt.Println("=== VOTACIÓN DE ACTIVIDADES ===")
	fmt.Println("Opciones: deportes, videojuegos, cine, musica\n")

	for i := 1; i <= 5; i++ {
		var opcion string
		fmt.Printf("Ingresa el voto %d: ", i)
		fmt.Scanln(&opcion)

	
		_, existe := votos[opcion]
		if existe {
			votos[opcion]++ 
		} else {
			fmt.Println("⚠️ Opción inválida. Intenta con: deportes, videojuegos, cine o musica.")
			i-- 
		}
	}

	
	fmt.Println("\n=== RESULTADOS DE LA VOTACIÓN ===")
	for actividad, cantidad := range votos {
		fmt.Printf("- %s: %d votos\n", actividad, cantidad)
	}

	
	ganadora := actividadMasVotada(votos)
	fmt.Printf("\n🏆 La actividad ganadora es: %s\n", ganadora)
}