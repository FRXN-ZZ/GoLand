package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Función para calcular el promedio (requerida por la Opción 1)
func averageGrade(grades []float64) float64 {
	sum := 0.0
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

// Opción 1: Notas de estudiantes
func ejecutarOpcion1() {
	var cantidad int
	fmt.Print("Ingrese la cantidad de estudiantes: ")
	fmt.Scan(&cantidad)

	var notas []float64
	for i := 1; i <= cantidad; i++ {
		var nota float64
		fmt.Printf("Ingrese la nota del estudiante %d (0-100): ", i)
		fmt.Scan(&nota)
		notas = append(notas, nota)
	}

	promedio := averageGrade(notas)
	fmt.Printf("\nEl promedio del curso es: %.2f\n", promedio)

	// If para determinar si aprueba o reprueba
	if promedio >= 70 {
		fmt.Println("Estado: Aprobado")
	} else {
		fmt.Println("Estado: Reprobado")
	}

	// Switch para el rango del promedio
	switch {
	case promedio >= 90:
		fmt.Println("Rendimiento: Excellent performance")
	case promedio >= 80:
		fmt.Println("Rendimiento: Good performance")
	case promedio >= 70:
		fmt.Println("Rendimiento: Satisfactory performance")
	default:
		fmt.Println("Rendimiento: Needs improvement")
	}
}

// Opción 2: Suma de números del 1 al n
func ejecutarOpcion2() {
	var n int
	fmt.Print("Ingrese el número n: ")
	fmt.Scan(&n)

	suma := 0
	for i := 1; i <= n; i++ {
		suma += i
	}
	fmt.Printf("La suma del 1 al %d es: %d\n", n, suma)
}

// Opción 3: Celsius a Fahrenheit
func ejecutarOpcion3() {
	var celsius float64
	fmt.Print("Ingrese la temperatura en grados Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Printf("%.2f °C equivalen a %.2f °F\n", celsius, fahrenheit)
}

// Opción 4: Fahrenheit a Celsius
func ejecutarOpcion4() {
	var fahrenheit float64
	fmt.Print("Ingrese la temperatura en grados Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f °F equivalen a %.2f °C\n", fahrenheit, celsius)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Bucle para que el menú se repita hasta que escriban "salir" o "0"
	for {
		fmt.Println("\n--- MENÚ PRINCIPAL ---")
		fmt.Println("1. Calcular promedio de notas")
		fmt.Println("2. Suma de números del 1 al n")
		fmt.Println("3. Convertir Celsius a Fahrenheit")
		fmt.Println("4. Convertir Fahrenheit a Celsius")
		fmt.Println("0. Salir (o escribe 'salir')")
		fmt.Print("Elige una opción: ")

		scanner.Scan()
		opcion := strings.TrimSpace(scanner.Text())

		if opcion == "salir" || opcion == "0" {
			fmt.Println("¡Hasta luego!")
			break
		}

		switch opcion {
		case "1":
			ejecutarOpcion1()
		case "2":
			ejecutarOpcion2()
		case "3":
			ejecutarOpcion3()
		case "4":
			ejecutarOpcion4()
		default:
			fmt.Println("Opción no válida, intenta de nuevo.")
		}
	}
}