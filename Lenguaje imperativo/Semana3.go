package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	lProductos = append(lProductos, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cant int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 && cant > 0 {
		if (*l)[prod].cantidad >= cant {
			(*l)[prod].cantidad = (*l)[prod].cantidad - cant
		} else {
			fmt.Println("No se puede mayor cantidad de productos que los que hay en existencia")
		}

		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista"
	}
}
func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	var lista listaProductos
	for i := 0; i < len(lProductos); i++ { //se recorre la lista de productos para encontrar el minimo
		p := (*l)[i]
		if p.cantidad < existenciaMinima { //se evalua en base a la cantidad del producto
			lista = append(lista, p) //se agrega a una lista nueva la struct
		}
	}
	return lista
}
func (l *listaProductos) aumentarInventarioDeMinimos() listaProductos {
	for i := 0; i < len(*l); i++ {
		if (*l)[i].cantidad < existenciaMinima {
			(*l)[i].cantidad = existenciaMinima
		}
	}
	fmt.Println("Minimos aumentados ")
	return *l
}
func (l *listaProductos) sortLista(opcion string) {
	if opcion == "nombre" {
		sort.SliceStable(lProductos, func(i, j int) bool {
			return lProductos[i].nombre < lProductos[j].nombre
		})
	} else if opcion == "cantidad" {
		sort.SliceStable(lProductos, func(i, j int) bool {
			return lProductos[i].cantidad < lProductos[j].cantidad
		})
	} else {
		sort.SliceStable(lProductos, func(i, j int) bool {
			return lProductos[i].precio < lProductos[j].precio
		})
	}
}
func main() {
	llenarDatos()
	//fmt.Println(lProductos)
	//lProductos.venderProducto("arroz", 4)
	//fmt.Println(lProductos)
	//lProductos.agregarProducto("azucar", 20, 1500)
	//fmt.Println(lProductos)
	//lProductos.venderProducto("frijoles", 4)
	//fmt.Println(lProductos)
	//lProductos.venderProducto("leche", 10)
	/*
	   Lista de productos minimos:  [{frijoles 4 2000} {leche 8 1200}]
	   Minimos aumentados
	   [{frijoles 10 2000} {leche 10 1200}]
	   Este es ordenado por nombre:  [{arroz 15 2500} {café 12 4500} {frijoles 4 2000}
	   {leche 8 1200}]
	   Este es ordenado por cantidad:  [{frijoles 4 2000} {leche 8 1200} {café 12 4500}
	    {arroz 15 2500}]
	   Este es ordenado por precio:  [{leche 8 1200} {frijoles 4 2000} {arroz 15 2500}
	   {café 12 4500}]

	*/
	lista := lProductos.listarProductosMínimos()
	fmt.Println("Lista de productos minimos: ", lista)
	lista.aumentarInventarioDeMinimos()
	fmt.Println(lista)
	lProductos.sortLista("nombre")
	fmt.Println("Este es ordenado por nombre: ", lProductos)
	lProductos.sortLista("cantidad")
	fmt.Println("Este es ordenado por cantidad: ", lProductos)
	lProductos.sortLista("precio")
	fmt.Println("Este es ordenado por precio: ", lProductos)
}

// si ha sobrado tiempo antes del receso, el ejercicio se podría ampliar para que los productos se almacenen en archivos de texto
// que al inicio se carguen del archivo a la lista y que al final se actualice el archivo con el contenido de la lista
