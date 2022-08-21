package main

/*
ejercicio5(3, 0, []string{"a", "b", "c", "d", "e", "f", "g", "h"})
Retorna: [d, e, f, g, h, a, b, c]

ejercicio5(3, 2, []string{"a", "b", "c", "d", "e", "f", "g", "h"})
Retorna:  [b c d e f g h a]

*/
import "fmt"

func acomodaDerEj5(lista *[]string, pivote int, dato string) []string {
	if (*lista)[len(*lista)-1] == (*lista)[pivote] { //si se llega a la ultima posicion, termina.
		(*lista)[0] = dato
		return *lista
	}
	if (*lista)[pivote] == dato {
		dato = (*lista)[pivote+1]
		(*lista)[pivote+1] = (*lista)[pivote]
		return acomodaDerEj5(lista, pivote+1, dato)
	}
	(*lista)[0] = (*lista)[pivote+1] //se usa esa posicion para guardar datos
	(*lista)[pivote+1] = dato
	return acomodaDerEj5(lista, pivote+1, (*lista)[0])
}
func acomodaIzqEj5(lista *[]string, pivote int, dato string) []string {
	if pivote == 0 { //cuando el pivote llegue a la primera posicion de la lista
		//lista[len(lista)-1] = lista[0]
		return *lista
	}
	if (*lista)[pivote] == dato {
		dato = (*lista)[pivote-1]
		(*lista)[pivote-1] = (*lista)[pivote]
		return acomodaIzqEj5(lista, pivote-1, dato)
	}
	(*lista)[len(*lista)-1] = (*lista)[pivote-1] //espacio que no se ocupa por el momento en la lista
	(*lista)[pivote-1] = dato
	return acomodaIzqEj5(lista, pivote-1, (*lista)[len(*lista)-1])
}
func ejercicio5(izq int, der int, lista *[]string) []string {
	if izq == 0 && der == 0 {
		return *lista
	}
	if izq != 0 {
		izq -= 1
		acomodaIzqEj5(lista, len(*lista)-1, (*lista)[len(*lista)-1]) //ultima posicion de la lista como pivote
	}
	if der != 0 {
		der -= 1
		acomodaDerEj5(lista, 0, (*lista)[0]) //primera posicion para el pivote
	}
	return ejercicio5(izq, der, lista)
}

func main() {
	lista := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	a := &lista
	fmt.Println("Retorna: ", ejercicio5(3, 2, a))
}
