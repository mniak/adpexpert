package main

import (
	"fmt"
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
	cli := adpexpert.Client{}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	handle(cli.Login(username, password))
	// handle(cli.PunchIn())

	punches, err := cli.GetLastPunches()
	handle(err)

	fmt.Println("Last Punches:")
	for _, punch := range punches.LastPunches {
		fmt.Println(punch.PunchDateTime)
	}
}
