package main

import (
	"strconv"
	"strings"
	"syscall/js"
)

func add(this js.Value, i []js.Value) interface{} {
	println(js.ValueOf(i[0].Int() + i[1].Int()).String())
	return js.ValueOf(i[0].Int() + i[1].Int())
}

func subtract(this js.Value, i []js.Value) interface{} {
	println(js.ValueOf(i[0].Int() - i[1].Int()).String())
	return js.ValueOf(i[0].Int() - i[1].Int())
}

func addArray(this js.Value, i []js.Value) interface{} {

	sum := 0

	slice := strings.Split(i[0].String(), ",")

	for _, value := range slice {
		vi, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		sum += vi
	}

	return js.ValueOf(sum)
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
	js.Global().Set("addArray", js.FuncOf(addArray))
}

func main() {
	c := make(chan struct{}, 0)

	println("Wasm Go Initializer")

	registerCallbacks()
	<-c
}
