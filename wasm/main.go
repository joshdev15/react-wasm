package main

import (
	"fmt"
	"syscall/js"
	"time"
)

var (
	// Instanciamos el objeto global
	// de js (window)
	global = js.Global()

	// Definimos una lista de
	// handlers que son nuestras funciones
	// Sum y AsyncSum
	functions = []function{
		{
			path:    "sum",
			handler: sum,
		},
		{
			path:    "asyncSum",
			handler: asyncSum,
		},
	}
)

// Definimos el type function
// que posee la propiedad path y handler
type function struct {
	path    string
	handler func(js.Value, []js.Value) interface{}
}

// La funcion New añade al
// objeto Global de javascript
// las funciones que creamos
// y retorna el objeto Global
// actualizado
func newGlobal() *js.Value {
	for _, fn := range functions {
		global.Set(fn.path, js.FuncOf(fn.handler))
	}

	return &global
}

// Funcion manejadora Sum
func sum(this js.Value, args []js.Value) interface{} {
	fmt.Println("Running: Sum")
	if len(args) < 2 {
		return "There must be at least 2 arguments"
	}

	aValue := args[0].Int()
	bValue := args[1].Int()

	return aValue + bValue
}

// Funcion manejadora AsyncSum
func asyncSum(this js.Value, args []js.Value) interface{} {
	fmt.Println("Running: AsyncSum")
	if len(args) < 2 {
		return "There must be at least 2 arguments"
	}

	aValue := args[0].Int()
	bValue := args[1].Int()

	time.Sleep(3 * time.Second)

	return aValue + bValue
}

// Llamamos la funcion main
func main() {
	// Usamos el patron signal para
	// mantener el proceso corriendo
	// en todo el ciclo de vida de la
	// aplicación
	signal := make(chan struct{})

	// Agregamos las funciones
	// creadas por nosotros
	value := newGlobal()

	// Si value no es nulo imprimimos un saludo
	if value != nil {
		fmt.Println("Iniciando WASM BE")
	}

	// Mostramos el valor de signal
	fmt.Println(<-signal)
}
