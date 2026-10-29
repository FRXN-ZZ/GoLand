package main

import (
	"errors"
	"fmt"
)

type Estudiante struct {
	edad int
}

func (e Estudiante) GetEdad() int {
	return e.edad
}

func (e *Estudiante) SetEdad(edad int) error {
	if edad < 0 || edad > 80 {
		return errors.New("ERROR, la edad debe estar entre 1 y 79 años")
	}
	e.edad = edad
	return nil
}

func main() {
	estudiante := Estudiante{}

	err := estudiante.SetEdad(100)

	if err != nil {
		fmt.Println("No se puede modificar", err)
	}
	fmt.Println("Hola, su edad es: ", estudiante.GetEdad())
}








package main

import (
	"fmt"  // Librería para imprimir textos y formatear salidas en la consola
	"sort" // Librería para ordenar elementos (la usaremos para ordenar los IDs)
)

// ==========================================
// 1. LA ESTRUCTURA PRODUCTO
// ==========================================
// Un 'struct' es como un formulario o plantilla que agrupa datos relacionados.
type Producto struct {
	ID       int     // Identificador único del producto (número entero)
	Nombre   string  // Nombre del producto (texto)
	Precio   float64 // Precio unitario (número decimal)
	Cantidad int     // Cantidad disponible en stock (número entero)
}

// ==========================================
// 2. CONSTRUCTOR DE PRODUCTO
// ==========================================
// Una función que te ayuda a empaquetar los datos y crear un Producto nuevo fácilmente.
func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}

// ==========================================
// 3. MÉTODOS DE LA ESTRUCTURA PRODUCTO
// ==========================================

// ActualizarPrecio: Permite cambiar el precio. 
// Usamos un puntero (*Producto) para modificar el objeto original en la memoria.
func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	if nuevoPrecio < 0 {
		fmt.Println("Error: el precio no puede ser negativo.")
		return
	}
	p.Precio = nuevoPrecio
}

// ActualizarCantidad: Permite cambiar el stock disponible de forma segura.
func (p *Producto) ActualizarCantidad(nuevaCantidad int) {
	if nuevaCantidad < 0 {
		fmt.Println("Error: la cantidad no puede ser negativa.")
		return
	}
	p.Cantidad = nuevaCantidad
}

// MostrarInfo: Imprime en pantalla los detalles del producto de manera ordenada.
func (p Producto) MostrarInfo() {
	fmt.Println("----------------------------")
	fmt.Printf("ID:        %d\n", p.ID)
	fmt.Printf("Nombre:    %s\n", p.Nombre)
	fmt.Printf("Precio:    $%.2f\n", p.Precio)
	fmt.Printf("Cantidad: %d unidades\n", p.Cantidad)
	fmt.Println("----------------------------")
}

// ==========================================
// 4. LA ESTRUCTURA INVENTARIO
// ==========================================
// El inventario almacena los productos en un mapa (map). 
// La clave (key) es el ID (int) y el valor es el struct Producto completo.
type Inventario struct {
	Productos map[int]Producto
}

// Constructor del Inventario: Inicializa el mapa para evitar errores de memoria.
func NuevoInventario() *Inventario {
	return &Inventario{
		Productos: make(map[int]Producto),
	}
}

// ==========================================
// 5. MÉTODOS DE LA ESTRUCTURA INVENTARIO
// ==========================================

// AgregarProducto: Inserta un producto nuevo al inventario si el ID no está repetido.
func (inv *Inventario) AgregarProducto(p Producto) {
	// Verificamos si ya existe un producto con ese mismo ID
	if _, existe := inv.Productos[p.ID]; existe {
		fmt.Printf("Error: ya existe un producto con ID %d.\n", p.ID)
		return
	}
	// Si no existe, lo guardamos en el mapa usando su ID como llave
	inv.Productos[p.ID] = p
	fmt.Printf("Producto '%s' agregado correctamente.\n", p.Nombre)
}

// EliminarProducto: Borra un producto del inventario buscando por su ID.
func (inv *Inventario) EliminarProducto(id int) {
	// Verificamos si el producto realmente existe antes de borrarlo
	if _, existe := inv.Productos[id]; !existe {
		fmt.Printf("Error: no existe un producto con ID %d.\n", id)
		return
	}
	delete(inv.Productos, id) // Función nativa de Go para eliminar elementos de un mapa
	fmt.Printf("Producto con ID %d eliminado.\n", id)
}

// BuscarProducto: Busca un producto por su ID. 
// Retorna el producto encontrado y un booleano (true si existe, false si no).
func (inv *Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.Productos[id]
	return p, existe
}

// ActualizarProducto: Reemplaza o actualiza los datos de un producto existente.
func (inv *Inventario) ActualizarProducto(p Producto) bool {
	if _, existe := inv.Productos[p.ID]; !existe {
		return false // Si no existe, retorna falso
	}
	inv.Productos[p.ID] = p // Sobrescribe con los nuevos datos
	return true
}

// ListarProductos: Muestra todos los productos del inventario ordenados por su ID.
func (inv *Inventario) ListarProductos() {
	if len(inv.Productos) == 0 {
		fmt.Println("El inventario está vacío.")
		return
	}

	// Como los mapas no tienen un orden fijo, sacamos los IDs a una lista temporal
	ids := make([]int, 0, len(inv.Productos))
	for id := range inv.Productos {
		ids = append(ids, id)
	}
	sort.Ints(ids) // Ordenamos los IDs de menor a mayor

	// Recorremos los IDs ya ordenados para mostrar la información en orden
	fmt.Println("\n===== INVENTARIO =====")
	for _, id := range ids {
		inv.Productos[id].MostrarInfo()
	}
}

// ==========================================
// 6. FUNCIÓN PRINCIPAL (MAIN)
// ==========================================
// Aquí es donde el programa cobra vida y se ejecuta paso a paso.
func main() {
	// 1. Creamos nuestro inventario vacío
	inventario := NuevoInventario()

	// 2. Agregamos tres productos iniciales válidos
	inventario.AgregarProducto(NuevoProducto(1, "Laptop", 850.50, 10))
	inventario.AgregarProducto(NuevoProducto(2, "Mouse", 15.99, 50))
	inventario.AgregarProducto(NuevoProducto(3, "Teclado", 45.00, 30))

	// 3. Intentamos agregar un producto con un ID repetido (el ID 1 ya lo tiene la Laptop)
	// Esto activará la validación de error y no se agregará.
	inventario.AgregarProducto(NuevoProducto(1, "Monitor", 200.00, 5))

	// 4. Mostramos en pantalla el inventario actual
	inventario.ListarProductos()

	// 5. Buscamos el producto con ID 2 (el Mouse)
	if p, ok := inventario.BuscarProducto(2); ok {
		fmt.Println("\nProducto encontrado:")
		p.MostrarInfo()

		// Modificamos su precio y cantidad usando sus métodos
		p.ActualizarPrecio(17.50)
		p.ActualizarCantidad(40)
		
		// Guardamos los cambios actualizados de vuelta en el inventario
		inventario.ActualizarProducto(p)
	}

	// 6. Intentamos buscar un producto que no existe (ID 99)
	if _, ok := inventario.BuscarProducto(99); !ok {
		fmt.Println("\nEl producto con ID 99 no existe.")
	}

	// 7. Eliminamos el producto con ID 3 (el Teclado)
	inventario.EliminarProducto(3)

	// 8. Mostramos el estado final del inventario después de todas las operaciones
	fmt.Println("\nEstado final del inventario:")
	inventario.ListarProductos()
}