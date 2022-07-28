package main

import (
	"fmt"
	"strings"
)

func print() {

	// Println -> print in own line
	fmt.Println("Hello mati!")


	name := "mati"
	var verb string = "to"
	var mati = "ciul"

	fmt.Println(name, verb, mati)

	num1 := 3
	num2 := 17

	fmt.Println(num1 * num2)

	// var flo1 = 2.3
	// fmt.Println(flo1 * num2)



	// Print => no automatic new line
	fmt.Print("hello mate, \n")
	fmt.Print("hello pal")



	// PrintF

	// default formatting
	fmt.Printf("my age is %v and my name is %v \n", num2, name)
	// formatting w/ quotes around stings
	fmt.Printf("my age is %q and my name is %q \n", num2, name)
	// %T -> print type
	fmt.Printf("my age is %T and my name is %T \n", num2, name)
	// print float with 1 decimal %0.1f
	fmt.Printf("this is going to be a floaty thing %0.1f \n", 255.55555)



	// Sprintf (save formatted strings)
	string := fmt.Sprintf("this shall be saved, but you not quite")
	fmt.Println(string)
}

func arrays () {

	// arrays have predefined length and data type, also note the curly braces, bruh
	// standard
	// var ages[3]int = [3]int{1, 2, 3}
	// inferred types
	var ages = [3]int{1, 2, 3}

	fmt.Println(ages)
	fmt.Println(len(ages))
}


func slices() {
	// slices are more like js arrays - don't require a predefined length
	var scores = []int{11, 21, 37}

	// set specific slice entry
	scores[2] = 99

	// append items - js push
	scores = append(scores, 81)

	fmt.Println(scores, len(scores))


	// slice ranges
	rangeOne := scores[1:3]
	rangeTwo := scores[3:]
	rangeThree := scores[:2]

	fmt.Println((rangeOne))
	fmt.Println((rangeTwo))
	fmt.Println((rangeThree))

	rangeOne = append(rangeOne, 77)

	fmt.Println(rangeOne)
}


func stringStuff() {
	greeting := "hello there, mate"

	fmt.Println("does this string contain 'hello'?", strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
}



func main() {
	// print()

	// arrays()
	// slices()

	stringStuff()
}
