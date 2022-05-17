package main

import (
	"log"
	"os"

	"github.com/mniak/adpexpert"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cli := adpexpert.Client{
		Debug: true,
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	handle(cli.Login(username, password))
	handle(cli.PunchIn())
}
