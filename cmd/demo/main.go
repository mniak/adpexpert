package main

import (
	"os"

	"github.com/mniak/adpexpert"
)

func main() {
	cli := adpexpert.Client{}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	cli.Login(username, password)
}
