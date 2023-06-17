package main

import (
	"errors"
	"fmt"
	"syscall/js"
	"time"
)

var (
	// Instanciamos el objeto global
	// de js (window)
	global    = js.Global()
	jsErr     = js.Global().Get("Error")
	jsPromise = js.Global().Get("Promise")

	// Definimos una lista de
	// handlers que son nuestras funciones
	// Sum y AsyncSum
	functions = []function{
		{
			path:    "sum",
			handler: js.FuncOf(sum),
		},
		{
			path:    "asyncSum",
			handler: promise(asyncSum),
		},
	}
)

// Definimos el type function
// que posee la propiedad path y handler
type function struct {
	path    string
	handler js.Func
}

type asyncFunction func(this js.Value, args []js.Value) (any, error)

// La funcion New añade al
// objeto Global de javascript
// las funciones que creamos
// y retorna el objeto Global
// actualizado
func newGlobal() *js.Value {
	for _, fn := range functions {
		global.Set(fn.path, fn.handler)
	}

	return &global
}

// Promise - Encapsulador de promesa
// La siguiente funcion tiene la tarea de comportarse como lo
// haría una promesa en JavaScript, lo que emulara el
// comportamiento retornando una promesa
func promise(callback asyncFunction) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		handler := js.FuncOf(func(_ js.Value, promFn []js.Value) any {
			resolve, reject := promFn[0], promFn[1]

			go func() {
				defer func() {
					if r := recover(); r != nil {
						reject.Invoke(jsErr.New(fmt.Sprint("panic:", r)))
					}
				}()

				res, err := callback(this, args)
				if err != nil {
					reject.Invoke(jsErr.New(err.Error()))
				} else {
					resolve.Invoke(res)
				}
			}()

			return nil
		})

		return jsPromise.New(handler)
	})
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
// A diferencia de Sum, AsyncSum retorna dos valores
// emulando una respuesta correcta y un error, que es
// lo que esperaríamos normalmente en una promesa de JS.
// Esta funcion sera llamada dentro de nuestro
// encapsulador
func asyncSum(this js.Value, args []js.Value) (interface{}, error) {
	fmt.Println("Running: AsyncSum")
	if len(args) < 2 {
		return nil, errors.New("There must be at least 2 arguments")
	}

	aValue := args[0].Int()
	bValue := args[1].Int()

	time.Sleep(3 * time.Second)

	return aValue + bValue, nil
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
