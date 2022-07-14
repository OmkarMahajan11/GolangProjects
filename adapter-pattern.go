package main

import "fmt"

type IProcess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}

type Adaptee struct {
	adapterType int
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

func main() {
	var processor IProcess = Adapter{}
	processor.process()
}
