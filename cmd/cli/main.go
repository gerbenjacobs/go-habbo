package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gerbenjacobs/go-habbo"
	"github.com/gerbenjacobs/go-habbo/habbo"
)

func main() {
	// set up Habbo Go API client
	parser := client.NewParser(http.DefaultClient)
	api := client.NewHabboAPI(parser)

	// check arguments
	if len(os.Args) < 3 {
		log.Fatal("Usage: go-habbo <hotel> <habboName>")
	}
	hotel := os.Args[1]
	identifier := os.Args[2]

	if !habbo.IsValidHotel(hotel) {
		log.Fatalf("Invalid hotel provided: %v", hotel)
	}

	// fetch Habbo information
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	h, err := api.GetHabboByName(ctx, hotel, identifier)
	if err != nil {
		log.Fatalf("failed to lookup %s from %s: %v", identifier, hotel, err)
	}

	b, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal Habbo: %v", err)
	}

	// remove all flags, allows this to be piped
	l := log.New(os.Stdout, "", 0).Writer()
	_, err = l.Write(b)
	if err != nil {
		log.Fatalf("failed to write out result: %v", err)
	}
}
