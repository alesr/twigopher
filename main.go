package main

import (
	"log"

	"github.com/alesr/twigopher/client"
)

func main() {

	_, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}
}
