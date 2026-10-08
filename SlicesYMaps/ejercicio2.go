package main

import "fmt"

// Función para encontrar la actividad con más votos
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
	// 1. Inicializar el map con las 4 actividades en 0 votos
	votos := map[string]int{
		"deportes":   0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	fmt.Println("=== VOTACIÓN DE ACTIVIDADES ===")
	fmt.Println("Opciones: deportes, videojuegos, cine, musica\n")

	// 2. Solicitar 5 votos por teclado
	for i := 1; i <= 5; i++ {
		var opcion string
		fmt.Printf("Ingresa el voto %d: ", i)
		fmt.Scanln(&opcion)

		// Verificar si la opción ingresada existe en el map
		_, existe := votos[opcion]
		if existe {
			votos[opcion]++ // Aumentar voto
		} else {
			fmt.Println("⚠️ Opción inválida. Intenta con: deportes, videojuegos, cine o musica.")
			i-- // Repetir el intento si se equivocó
		}
	}

	// 3. Mostrar resultados del map
	fmt.Println("\n=== RESULTADOS DE LA VOTACIÓN ===")
	for actividad, cantidad := range votos {
		fmt.Printf("- %s: %d votos\n", actividad, cantidad)
	}

	// 4. Determinar el ganador
	ganadora := actividadMasVotada(votos)
	fmt.Printf("\n🏆 La actividad ganadora es: %s\n", ganadora)
}