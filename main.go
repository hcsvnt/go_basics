package main

import (
	"fmt"
	"sort"
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
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.ToLower(greeting))
	fmt.Println(strings.Index(greeting, "mate"))
	fmt.Println(strings.Split(greeting, " "))
}


func sortStuff() {
	ages := []int{1, 22, 3, 4, 5, 6, 7, 8, 9}

	// transorms the original slice
	sort.Ints(ages)
	fmt.Println(ages)

	// if doesn't exist, returns length of slice + 1
	index := sort.SearchInts(ages, 77)
	fmt.Println((index))


	names := []string{"mati", "maciek", "wario"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "maciek"))
}


func loopStuff() {

	// x := 0
	// for x < 5 {
	// 	fmt.Println(x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	names := []string{"mati", "maciek", "wario"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Println(index, value)
	// }

	// omit index
	for _, value := range names {
		fmt.Println(value)
		// this will not modify the original slice, it's just a copy / reference?
		value = "new value"
	}
}



func main() {
	// print()

	// arrays()
	// slices()

	// stringStuff()

	// sortStuff()
	loopStuff()
}
