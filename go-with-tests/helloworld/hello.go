package main

import (
	"di"
	"log"
	"net/http"
	// "os"
)

func Hello(name, language string) string {
	if (name == "") {
		name = "World"
	}

	prefix := "Hello "
	switch language {
		case "Filipino":
			prefix = "Kamusta "
	}

	return prefix + name
}

func main() {
	// di.Greet(os.Stdout, Hello("Aeron",  "English"))
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(di.MyGreeterHandler)))
}