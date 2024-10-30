package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a JSON file as an argument.")
        os.Exit(1)
    }

    filename := os.Args[1]

    content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	result := Deserializer(content)
	if result {
		fmt.Printf("The file is a valid\n")
	} else {
		fmt.Printf("The file is invalid\n")
	}
}
