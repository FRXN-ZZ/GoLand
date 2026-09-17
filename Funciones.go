package main

import "fmt"

/*

func <nombre>(param1, param2, ....param n)<valores de retorno>{
  ------------------------
  ------------------------
  ------------------------
  //return en el caso de que nuestra funcion retorne valores
}

*/
func saludar() {
	fmt.Println("Hola esta es mi primera función")
}

func bienvenida(nombre string) {
	fmt.Println("Bienvenido@", nombre)
}

func suma(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func main() {
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
/* Variadica Functions ======= Funciones Variadicas*/

func sumatoria(numeros ...int) int{
	total :=0
	for _,numero:= range numeros{
		total+= numero 
	}
	return total
}