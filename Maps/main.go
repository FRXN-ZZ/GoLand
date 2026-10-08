package main

import "fmt"

// Estructura map[tipodeclave]tipovalor

func main() {
	visitas := make(map[string]int)
	visitas["Inicio"] = 10
	visitas["Noticias"] = 20
	visitas["Deportes"] = 25
	visitas["Hogar"] = 80
	fmt.Println("Las Noticias de Hogar son:", visitas["Hogar"])

	visitas["Hogar"] = 20
	fmt.Println("Actualizado Hogar....\nLas Noticias de Hogar ahora son:", visitas["Hogar"])

	delete(visitas, "Hogar")
	fmt.Println(visitas)
	fmt.Println("El conjunto de visitas se clasifica en:")
	for key, val := range visitas {
		fmt.Println(key, ":", val)
	}

	visitas["Galeria"] = 20
	fmt.Println("Las Noticias de Galeria son:", visitas["Galeria"]

	visitas["Inicio0"] = 200
	fmt.Println("Ahora las Noticias de Inicio son:", visitas["Inicio"]
	fmt.Println("El Total de vistas es:, suma(visitas))

}

Func suma(visitas map[sting]int)int{
  sum := 0
  for _, val := range visitas {
     sum += val
  }
	return sum
}