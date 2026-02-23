package main

import (
	// "di"
	"fmt"
	// "log"
	// "net/http"
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

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, b+a
		return a
	}
}

func main() {
	// di.Greet(os.Stdout, Hello("Aeron",  "English"))
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(di.MyGreeterHandler)))
	a := fib()
	fmt.Println(a(), a(), a(), a())
}