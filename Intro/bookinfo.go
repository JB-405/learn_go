package main

import "fmt"

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber int
	var grade float32
	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5
	fmt.Println(title, "written by", writer, "drawn by", artist)
	fmt.Println(publisher, year, pageNumber, grade)
	title = "Epic Vol. 1"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.0
	fmt.Println(title, "written by", writer, "drawn by", artist)
	fmt.Println(publisher, year, pageNumber, grade)

	// There's a mix of Println and Print!
	fmt.Println("Can", "you", "tell", "the", "difference?")
	fmt.Print("Between", "these", "two", "methods?")
	fmt.Print("Anything odd about", "the spacing? \n")
	fmt.Println("Don't worry if you can't spot it, we'll go through this together!")

	fmt.Println("Let's first see how", "the Println() method works.")
	fmt.Println("Notice that each statement adds a newline for us.")
	fmt.Println("There's also a default space", "between the string arguments.")
	// Add your code below:
	fmt.Println("Print", "is", "different")
	fmt.Print("Print", "is", "different")
	fmt.Print("See?")

	animal1 := "cat"
	animal2 := "dog"

	// Add your code below:
	fmt.Printf("Are you a %v or a %v person?", animal1, animal2)

	specialNum := 42
	fmt.Printf("This value's type is %T.", specialNum)
	// Prints: This value's type is int.

	quote := "To do or not to do"
	fmt.Printf("This value's type is %T.", quote)
	// Prints: This value's type is string.

	floatExample := 1.75
	// Edit the following Printf for the FIRST step
	fmt.Printf("Working with a %T", floatExample)

	fmt.Println("\n***") // Added for spacing

	yearsOfExp := 3
	reqYearsExp := 15
	// Edit the following Printf for the SECOND step
	fmt.Printf("I have %d years of Go experience and this job is asking for %d years", yearsOfExp, reqYearsExp)

	fmt.Println("\n***") // Added for spacing

	stockPrice := 3.50
	// Edit the following Printf for the THIRD step
	fmt.Printf("Each share of Gopher feed is $%.2f!", stockPrice)
}
