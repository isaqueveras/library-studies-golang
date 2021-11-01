package main

import "github.com/gen2brain/beeep"

func main() {
	var err error
	if err = beeep.Notify("Hello, World!", "Message body", "assets/hands.png"); err != nil {
		panic(err)
	}

	if err = beeep.Alert("Schedule accepted", "Schedule accepted successfully", "assets/schedule.png"); err != nil {
		panic(err)
	}
}
