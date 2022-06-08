package main

import (
	"fmt"
	"log"
	"os"
	"time"

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
	fmt.Println()

	now := time.Now()
	timecard, err := cli.GetTimecard(now.Year(), int(now.Month()))
	handle(err)

	for _, timetableEntry := range timecard.Timetable {
		if !timetableEntry.Inconsistent {
			continue
		}
		fmt.Println("Inconsistency:", timetableEntry.Date)
		for _, timelineEntry := range timetableEntry.Timeline {
			fmt.Println("  - ", timelineEntry.DateTime)
		}
	}
}
