package main

import (
	"errors"
	"log"
)

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(a, b string) {
		log.Println(formatter(a, b))
	}
}

// don't touch below this line

func main() {
	concatLogger := getLogger(concat)
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	for _, err := range dbErrors {
		concatLogger("Error on database server: ", err.Error())
	}

	commaDelimitLogger := getLogger(commaDelimit)
	mailErrors := []error{
		errors.New(" email too large"),
		errors.New(" non alphanumeric symbols found"),
	}
	for _, err := range mailErrors {
		commaDelimitLogger("Error sending mail", err.Error())
	}
}

// concat takes 2 strings and returns a new combined string
func concat(first, second string) string {
	return first + "" + second
}

// commaDelimit takes 2 strings and returns a new string
// where the inputs are joined by a ","
func commaDelimit(first, second string) string {
	return first + "," + second
}
