package main

import "fmt"

func main() {
	lockCombo := "2-35-19"
	robAttempt := "1-1-1"
	// Add your code below:
	if lockCombo == robAttempt {
		fmt.Print("The vault is now opened.")

	} else {
		fmt.Print("that must be the wrong one.")
	}

}
