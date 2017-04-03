package main

import (
	"fmt"
	"log"

	"github.com/alesr/twigopher/account"
	"github.com/alesr/twigopher/client"
)

func main() {

	client, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n\n", account.Stat(client))
}
