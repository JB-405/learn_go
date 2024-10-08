package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	leftSafely := rand.Intn(5)
	eludedGuards := rand.Intn(100)
	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}
	openedVault := rand.Intn(100)
	fmt.Println("isHeistOn is currently:", isHeistOn)
	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else {
		isHeistOn = false
		fmt.Println("vault can’t be opened.")
	}

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Looks like you tripped an alarm... run?")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2:
			fmt.Println("run faster")
		case 3:
			fmt.Println("run even faster")
		default:
			fmt.Println("Start the getaway car!")
		}
		if isHeistOn {
			amtStolen := 0000 + rand.Intn(1000000)
			fmt.Println("$", amtStolen, "Not bad!")
		}

	}

}
