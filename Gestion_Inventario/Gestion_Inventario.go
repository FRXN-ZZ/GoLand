package main

import (
	"fmt"
	"sort"
)

type Producto struct {
	id       int
	nombre   string
	precio   float64
	cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		id:       id,
		nombre:   nombre,
		precio:   precio,
		cantidad: cantidad,
	}
}

func (p Producto) ObtenerID() int {
	return p.id
}

func (p Producto) GetNombre() string {
	return p.nombre
}

func (p Producto) ObtenerPrecio() float64 {
	return p.precio
}

func (p Producto) ObtenerCantidad() int {
	return p.cantidad
}

func (p *Producto) SetPrecio(nuevoPrecio float64) {
	if nuevoPrecio < 0 {
		fmt.Println("Error: el precio no puede ser negativo.")
		return
	}
	p.precio = nuevoPrecio
}

func (p *Producto) SetCantidad(nuevaCantidad int) {
	if nuevaCantidad < 0 {
		fmt.Println("Error: la cantidad no puede ser negativa.")
		return
	}
	p.cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Println("----------------------------")
	fmt.Printf("ID:        %d\n", p.ObtenerID())
	fmt.Printf("Nombre:    %s\n", p.GetNombre())
	fmt.Printf("Precio:    $%.2f\n", p.ObtenerPrecio())
	fmt.Printf("Cantidad:  %d unidades\n", p.ObtenerCantidad())
	fmt.Println("----------------------------")
}

type Inventario struct {
	productos map[int]Producto
}

func NuevoInventario() *Inventario {
	return &Inventario{
		productos: make(map[int]Producto),
	}
}

func (inv *Inventario) AgregarProducto(p Producto) {
	if _, existe := inv.productos[p.ObtenerID()]; existe {
		fmt.Printf("Error: ya existe un producto con ID %d.\n", p.ObtenerID())
		return
	}
	inv.productos[p.ObtenerID()] = p
	fmt.Printf("Producto '%s' agregado correctamente.\n", p.GetNombre())
}

func (inv *Inventario) EliminarProducto(id int) {
	if _, existe := inv.productos[id]; !existe {
		fmt.Printf("Error: no existe un producto con ID %d.\n", id)
		return
	}
	delete(inv.productos, id)
	fmt.Printf("Producto con ID %d eliminado.\n", id)
}

func (inv *Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.productos[id]
	return p, existe
}

func (inv *Inventario) ListarProductos() {
	if len(inv.productos) == 0 {
		fmt.Println("El inventario está vacío.")
		return
	}

	ids := make([]int, 0, len(inv.productos))
	for id := range inv.productos {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	fmt.Println("\n===== INVENTARIO =====")
	for _, id := range ids {
		inv.productos[id].MostrarInfo()
	}
}

func main() {
	inventario := NuevoInventario()

	inventario.AgregarProducto(NuevoProducto(1, "Laptop", 850.50, 10))
	inventario.AgregarProducto(NuevoProducto(2, "Mouse", 15.99, 50))
	inventario.AgregarProducto(NuevoProducto(3, "Teclado", 45.00, 30))

	inventario.AgregarProducto(NuevoProducto(1, "Monitor", 200.00, 5))

	inventario.ListarProductos()

	if p, ok := inventario.BuscarProducto(2); ok {
		fmt.Println("\nProducto encontrado:")
		p.MostrarInfo()

		p.SetPrecio(17.50)
		p.SetCantidad(40)
		inventario.AgregarProducto(p)
	}

	fmt.Println()
	inventario.EliminarProducto(3)

	inventario.ListarProductos()
}
