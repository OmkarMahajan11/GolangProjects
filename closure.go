package main

import "fmt"

func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {
	makeHarryPotter := concatter()
	makeHarryPotter("Mr.")
	makeHarryPotter("and")
	makeHarryPotter("Mrs.")
	makeHarryPotter("Dursley")
	makeHarryPotter("of")
	makeHarryPotter("number")
	makeHarryPotter("four,")
	makeHarryPotter("Privet")
	fmt.Println(makeHarryPotter("Drive"))
	// prints Mr. and Mrs. Dursley of number four, Privet Drive
}
