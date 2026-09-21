package main

import "fmt"


func saludar() {
	fmt.Println("Hola esta es mi primera función")
}

func bienvenida(nombre string) {
	fmt.Println("Bienvenido@", nombre)
}

func suma(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func main () {
	var usr string

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scan(&usr)

	saludar()
	bienvenida(usr)
}



	func suma_resta (num1, num2 int) (int,int) {
		if num2<num1 {
			ResSuma := num1+num2
			ResResta := num2-num1
			return ResSuma, ResResta
		}else{
            ResSuma :=num1+num2
			ResResta := 0
			return ResSuma, ResResta
		}
}

func sumaresta(a int, b int)(int, int){
	if a<b {
		return b - a, a + b
	}
	return a + b,0
}

func mostrarNum(numeros ...int){
   fmt.Println("Los numeros ingresados son: ", numeros)
}
func main (){
	var usr string

	fmt.Println("Ingresa tu nombre")
	fmt.Println(&usr)
	saludar()
	bienvenida(usr)

	fmt.Println("El resultado de la suma es: ", suma(4, 5))
	res1, rest2 := suma_resta(5,8)

	fmt.Println("La suma es:", res1, "La resta es: ", rest2)
	mostrarNum(5, 10, 15, 20, 25)
	fmt.Println("La sumatoria es:", sumatoria(1, 2, 3, 4, 5, 6, 7, 8, 9, ))
}
func sumatoria(numeros ...int) int{
	total :=0
	for _,numero:= range numeros{
		total+= numero 
	}
	return total
}