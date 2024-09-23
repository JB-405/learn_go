package main

import "fmt"

func main() {
	var numOfFlavors int
	// Assign a value for numOfFlavors below:
	numOfFlavors = 57
	// Declare flavorScale below:
	fmt.Println(numOfFlavors)

	// Create three variables
	// emptyInt an int8
	var emptyInt int8
	// emptyFloat a float32
	var emptyFloat float32
	// and emptyString a string
	var emptyString string
	// Finally, print them all out
	fmt.Println(emptyInt, emptyFloat, emptyString)

	// Define daysOnVacation using := below:
	daysOnVacation := 7

	// Define hoursInDay using var and = below:
	var hoursInDay = 24

	fmt.Println("You have spent", daysOnVacation*hoursInDay, "hours on vacation.")

	coolSneakers := 65.99
	niceNecklace := 45.50

	// Define taxCalculation here
	var taxCalculation float64
	// Add coolSneakers to taxCalculation
	taxCalculation += coolSneakers
	// Add niceNecklace to taxCalculation
	taxCalculation += niceNecklace
	// Compute the NYC sales tax
	// 8.875% of the purchase here:
	taxCalculation *= .08875
	// Uncomment this line for a receipt!
	fmt.Println("Purchase of", coolSneakers+niceNecklace, "with 8.875% sales tax", taxCalculation, "equal to", coolSneakers+niceNecklace+taxCalculation)

	// Define magicNum and powerLevel below:
	var magicNum, powerLevel int32
	magicNum = 2048
	powerLevel = 9001
	fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)

	// Define amount and unit below:
	amount, unit := 10, "doll hairs"

	fmt.Println(amount, unit, ", that's expensive...")

}
