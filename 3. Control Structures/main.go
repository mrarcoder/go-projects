// 3. Control Structures
package main

import "fmt"

func main() {
	// If-else statement
	// fmt.Print("> What is your age?\n> ")
	// age := 0
	// fmt.Scanln(&age)
	// if age < 18 && age > 0 {
	// 	fmt.Println("> Teenager")
	// } else if age >= 18 && age <= 64 {
	// 	fmt.Println("> Adult")
	// } else if age >= 65 {
	// 	fmt.Println("> Senior")
	// } else {
	// 	fmt.Println("> Invalid Input")
	// }

	// we can also declare variable with if-condition declaration which can be only accessible inside if body
	// if marks := 7; marks >= 50 {
	// 	fmt.Println("Marks:",marks,"Passed")
	// } else {
	// 	fmt.Println("Marks:",marks,"Failed")
	// }

	// Switch Statement
	// carEngine := 0
	// fmt.Scanln(&carEngine)
	// switch carEngine {
	// case 1 :
	// 	fmt.Println("Start")
	// case 2 :
	// 	fmt.Println("Stop")
	// default :
	// 	fmt.Println("Unknown Input")
	// }

	// Switch without expression (like chained if-else)
	// temp := 35
	// switch {
	// case temp > 40:
	// 	fmt.Println("Too Hot")
	// case temp > 30:
	// 	fmt.Println("Hot")
	// case temp > 20:
	// 	fmt.Println("Warm")
	// default:
	// 	fmt.Println("Cool")
	// }

	// Use fallthrough to force execution of the next case
	x := 1
	switch x {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
	}

	// For Loop (There are multiple ways to use For loop in GO)

	// 1st way (Traditional way like in C++ or JavaScript)
	// for i:=0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 2nd way (with Range function like in Python)
	// for i:=range 10 {
	// 	fmt.Println(i)
	// }

	// 3rd way (For is the only loop in Go so instead of While loop we can use it as)
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 4th way (infinite loop)
	// i := 0
	// for {
	// 	if i == 10 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }
}
