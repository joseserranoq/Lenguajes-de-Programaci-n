package main

import "fmt"

/*
ejercicio1("Hola me llamo jose y estoy aprendiendo a programar en go.\n")
Retorna
	numero de caracteres = 48
	numero de palabras = 10
	numL int = 1
*/

func ejercicio1(cadena string) {
	var (
		numC int = 0
		numP int = 0
		numL int = 0
	)
	for i := 0; i < len(cadena); i++ {
		//cuenta los caracteres
		c := string(cadena[i])
		if c == "." {
			numC++
			numP++
		} else if c != " " {
			numC++
		} else {
			numP++
		}
		if c == "\n" {
			numL++
		}
		//fmt.Println(string(cadena[i]))

	}
	fmt.Println("Caracteres: ", numC, "\nPalabras: ", numP, "\nLineas: ", numL)
}
func main() {
	ejercicio1("Hola me llamo jose y estoy aprendiendo a programar en go.\n")
}
