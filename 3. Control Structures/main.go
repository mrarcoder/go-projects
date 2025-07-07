// 3. Control Structures
package main

import "fmt"

func main() {
	// If-else statement
	fmt.Print("> What is your age?\n> ")
	age := 0
	fmt.Scanln(&age)
	if age < 18 && age > 0 {
		fmt.Println("> Teenager")
	} else if age >= 18 && age <= 64 {
		fmt.Println("> Adult")
	} else if age >= 65 {
		fmt.Println("> Senior")
	} else {
		fmt.Println("> Invalid Input")
	}
	// Switch Statement
	carEngine := 0
	fmt.Scanln(&carEngine)
	switch carEngine {
	case 1 :
		fmt.Println("Start")
	case 2 :
		fmt.Println("Stop")
	default :
		fmt.Println("Unknown Input")
	}
}