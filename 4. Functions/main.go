package main

import (
	"fmt"
	"strings"
)

// Simple Function
func greet() {
	fmt.Println("Hello! from simple function")
}

// Function with parameter and no return type
func welcomeUser(name string) {
	fmt.Println("Welcome,", name, "Ready to learn Go?")
}

// Function with parameter and return type
func square(n float32) float32 {
	return n * n
}

// Function with multiple return values
func getDrink(d string, a int) (string, int) {
	switch d {
	case "Coca-Cola":
		return "Coca-Cola", a - 80
	case "7up":
		return "7up", a - 70
	case "Sprite":
		return "Sprite", a - 60
	default:
		return "No Drink Available", a
	}
}

// Function with named returned values
func calculateBill(price, quantity int) (total int, note string) {
	total = price * quantity
	if total > 500 {
		note = "Expensive"
	} else {
		note = "Reasonable"
	}
	return
}

// Variadic Function
func concatWords(words ...string) string {
	sentence := ""
	for i, w := range words {
		sentence += w
		if i < len(w)-1 {
			sentence += " "
		}
	}
	return sentence
}

// Anonymous Functions ( Used in Inside any function (main, init, etc.) Inside loops, if-else blocks, goroutines, etc. and As arguments to higher-order functions )
// func init() {
// 	// Syntax 1
// 	sayBye := func(name string) {
// 		fmt.Println("Goodbye", name)
// 	}
// 	sayBye("Ghost")

// 	// Syntax 2
// 	func(name string) {
// 		fmt.Println("Goodbye", name)
// 	}("Ali")
// }

// Functions as First-Class Citizens

func greetUser(name string, formatter func(string) string) string {
	return formatter(name)
}

func main() {
	// greet()

	// welcomeUser("Ghost")

	// fmt.Println(square(13))

	// fmt.Println(getDrink("Coca-Cola",100))
	// d,_ := getDrink("Pepsi",100)
	// fmt.Println(d)
	// _,a := getDrink("Pepsi",100)
	// fmt.Println(a)

	// fmt.Println(calculateBill(100,6))
	// fmt.Println(calculateBill(100,4))

	// sentence := concatWords("Hello","This","is","Variadic","Function")
	// fmt.Println(sentence)

	fmt.Println(greetUser("Ghost", func(name string) string {
		return "Hello, " + strings.ToUpper(name)
	}))
}
