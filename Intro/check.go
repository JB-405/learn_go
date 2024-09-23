package main

import "fmt"

func main() {
	grade := "100"
	compliment := "Great job!"
	teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)

	fmt.Print(teacherSays)

	step1 := "Breathe in..."
	step2 := "Breathe out..."

	// Add your code below:
	meditation := fmt.Sprintln(step1, step2)

	fmt.Print(meditation)

}
