package main

import "fmt"

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
	fmt.Println(Hello("Aeron",  "English"))
}