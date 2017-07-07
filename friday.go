package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	today := now.Weekday().String()
	days_left := "0"
	if today == "Friday" {
		fmt.Println("It's Friday!!!")
	} else {
		switch today {
		case "Monday":
			days_left = "4. Damn"
		case "Tuesday":
			days_left = "3."
		case "Wednesday":
			days_left = "2. Wednesday is a petite Friday!"
		case "Thursday":
			days_left = "1. Almost there!"
		default:
			days_left = "It's weekend. Take a rest"
		}
		fmt.Println(days_left)
	}
}
