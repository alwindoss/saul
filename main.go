package main

import (
	"fmt"
	"log"

	"github.com/alwindoss/saul/internal/server"
)

func main() {
	fmt.Println("Welcome to the SAUL")
	s := server.New(server.WithAddr("", "8080"))
	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Exiting SAUL")
}
