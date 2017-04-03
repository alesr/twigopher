package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alesr/twigopher/account"
	"github.com/alesr/twigopher/client"
	"github.com/alesr/twigopher/stream"
)

func main() {

	track := flag.String("track", "#golang, gopher, go programming language, trump", "a comma-separated list of phrases to track")
	flag.Parse()

	client, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n\n", account.Stat(client))

	stream := stream.NewStream(client, os.Stdout, *track)
	defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}

	// Wait for CTRL-C
	ch := make(chan os.Signal)
	defer close(ch)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
