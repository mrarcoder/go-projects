// 2, Variables and Data Types

package main

import "fmt"

func main() {

	// variables

	var chapterNumber int = 2 // Explicit type declaration
	
	var chapterName = "Variables and Data types" // Type inference (Go figures out the type automatically)

	topicName := "Variables" // Short declaration (only inside functions)

	fmt.Println("Chapter:",chapterNumber,chapterName,"| Topic:",topicName)

	// Constants

	const PI = 22.0 / 7.0 // Must use .0 if need and floating output otherwise it gives result just integer number

	fmt.Println(PI)

	// Data Types
	var age byte = 25 // byte (Short Integer)
	var weight float32 = 70.4 // floating numbers (32Bits Decimal Numbers)
	var height float64 = 5.7 // floating numbers (64Bits Decimal Numbers)
	var isMale bool = true // boolean (0 or 1)
	var name string = "Ghost" // string (Sentences) 
	var phone int = 123456789 // Integer (Numeric values)
	
	fmt.Println("Age:",age)
	fmt.Println("Weight:",weight)
	fmt.Println("Height:",height)
	fmt.Println("Male:",isMale)
	fmt.Println("Name:",name)
	fmt.Println("Phone:",phone)
}