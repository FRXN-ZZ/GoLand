package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func valorTotal(p Producto) float64 {
	return p.Precio * float64(p.Cantidad)
}

func tieneStock(p Producto) bool {
	return p.Cantidad > 0
}

func mostrarProducto(p Producto) {
	estado := "sin stock"
	if tieneStock(p) {
		estado = "disponible"
	}

	fmt.Println("Nombre:      ", p.Nombre)
	fmt.Printf("Precio:       $%.2f\n", p.Precio)
	fmt.Println("Cantidad:    ", p.Cantidad)
	fmt.Printf("Valor total:  $%.2f\n", valorTotal(p))
	fmt.Println("Stock:       ", estado)
	fmt.Println("-------------------------")
}

func main() {
	p1 := Producto{Nombre: "Laptop", Precio: 850.50, Cantidad: 5}
	p2 := Producto{Nombre: "Mouse", Precio: 12.99, Cantidad: 30}
	p3 := Producto{Nombre: "Teclado", Precio: 25.00, Cantidad: 0}

	productos := []Producto{p1, p2, p3}

	fmt.Println("=== LISTA DE PRODUCTOS ===")
	for _, p := range productos {
		mostrarProducto(p)
	}
}