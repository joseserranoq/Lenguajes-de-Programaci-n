package main

import (
	"fmt"
)

type calzado struct {
	marca  string
	precio int
	talla  uint8
}

func agregarCalzado(lista *[]calzado, nom string, precio int, talla uint8) {
	if talla >= 34 && talla <= 44 {
		*lista = append(*lista, calzado{nom, precio, talla})
	} else {
		fmt.Println("No se acepta la talla: ", talla)
	}
}
func vendeZapatos(lista []calzado, nom string, precio int, talla uint8) {
	for i, s := range lista {
		if s.marca == nom && precio == s.precio && talla == s.talla {
			lista = append(lista[:i], lista[i+1:]...)
			fmt.Println("Zapatos", s.marca, "comprados exitosamente")
			return
		}

	}
	fmt.Println("No hay en stock")
	return
}
func datosQuemados() []calzado {
	lCalzado := make([]calzado, 0)
	lista := &lCalzado
	agregarCalzado(lista, "Nike", 10000, 42)
	agregarCalzado(lista, "Adidas", 20000, 43)
	agregarCalzado(lista, "New Balance", 40000, 45)
	agregarCalzado(lista, "Broncos", 20000, 41)
	agregarCalzado(lista, "Nike", 40000, 44)
	agregarCalzado(lista, "Adidas", 30000, 42)
	agregarCalzado(lista, "Speedo", 45000, 45)
	agregarCalzado(lista, "Timberland", 70000, 44)
	agregarCalzado(lista, "Puma", 45000, 39)
	agregarCalzado(lista, "Timberland", 70000, 41)
	agregarCalzado(lista, "Puma", 45000, 39)
	agregarCalzado(lista, "Puma", 45000, 39)
	agregarCalzado(lista, "Puma", 45000, 39)
	agregarCalzado(lista, "Timberland", 70000, 41)
	return lCalzado
}
func main() {
	var lista []calzado = datosQuemados()
	vendeZapatos(lista, "Puma", 45000, 39)
	vendeZapatos(lista, "Puma", 45000, 39)
	vendeZapatos(lista, "Puma", 45000, 39)
	vendeZapatos(lista, "Puma", 45000, 39)
	vendeZapatos(lista, "Puma", 45000, 39)
	vendeZapatos(lista, "Puma", 45000, 39)
}
