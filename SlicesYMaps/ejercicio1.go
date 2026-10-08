package main

import "fmt"

func main() {
	// Array de 6 estudiantes y 4 materias con notas de ejemplo
	notas := [6][4]float64{
		{8.5, 9.0, 7.5, 10.0}, // Estudiante 1
		{6.0, 7.0, 6.5, 8.0},  // Estudiante 2
		{9.5, 9.0, 9.8, 10.0}, // Estudiante 3
		{5.0, 6.5, 7.0, 5.5},  // Estudiante 4
		{8.0, 8.5, 8.0, 9.0},  // Estudiante 5
		{7.0, 7.5, 8.0, 7.8},  // Estudiante 6
	}

	var sumaClase float64
	totalNotas := 6 * 4

	fmt.Println("=== ANÁLISIS DE NOTAS DE ESTUDIANTES ===")

	// Recorremos cada estudiante usando un slice para sus notas
	for i, notasEstudiante := range notas {
		// Convertimos la fila del array en un slice para procesarlo
		sliceNotas := notasEstudiante[:]

		sumaEstudiante := 0.0
		notaMasAlta := sliceNotas[0]
		notaMasBaja := sliceNotas[0]

		for _, nota := range sliceNotas {
			sumaEstudiante += nota
			sumaClase += nota

			// Verificar nota más alta
			if nota > notaMasAlta {
				notaMasAlta = nota
			}
			// Verificar nota más baja
			if nota < notaMasBaja {
				notaMasBaja = nota
			}
		}

		promedioEstudiante := sumaEstudiante / float64(len(sliceNotas))

		fmt.Printf("Estudiante %d:\n", i+1)
		fmt.Printf("  - Promedio: %.2f\n", promedioEstudiante)
		fmt.Printf("  - Nota más alta: %.2f\n", notaMasAlta)
		fmt.Printf("  - Nota más baja: %.2f\n", notaMasBaja)
		fmt.Println("---------------------------------")
	}

	// Promedio general de la clase
	promedioGeneral := sumaClase / float64(totalNotas)
	fmt.Printf("Promedio general de la clase: %.2f\n", promedioGeneral)
}