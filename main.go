package main

import "fmt"


func main(){
	fmt.Println("Quiz game")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, welcome to the quiz game! \n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	//Conditions



	//Conditionals
	//if / elif / else

	if age >= 10 {
		fmt.Println("You can play!")
	} else {
		fmt.Println("You can't play")
		return
	}

	score := 0
	num_questions := 2

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer + " " + answer2 == "RTX 3090"{
		
		fmt.Println("Correct!")
		score += 1
	
		} else if answer + " " + answer2 == "rtx 3090"{
			 
		fmt.Println("Correct!")
		score += 1
	}else if answer + " " + answer2 == "Rtx 3090"{
			 
		fmt.Println("Correct!")
		score += 1

	}else if answer + " " + answer2 == "Rtx3090"{
			 
		fmt.Println("Correct!")
		score += 1
		

	}else if answer + " " + answer2 == "rtx3090"{
			 
		fmt.Println("Correct!")
		score += 1

	}else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the RYZEN 9 3990x have? ")
	var cores int
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score += 1
	}else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100

	fmt.Printf("You scored: %v%%", percent)


}