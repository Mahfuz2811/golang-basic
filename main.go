package main

import (
	"fmt"
	"sort"
)

func main() {
	// variables
	/*var nameOne string = "Mahfuz"
	var nameTwo = "Shishir"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Hello"
	nameThree = "Go"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Hello Go"

	fmt.Println(nameFour) */

	// formated string
	/*var ageOne int = 10
	var ageTwo = 20
	ageThree := 30

	fmt.Println(ageOne, ageTwo, ageThree)

	var numerOne int8 = 12
	var name string = "Mahfuz"

	fmt.Printf("My name is %v and my age is %v", name, numerOne)
	fmt.Printf("\n")

	var str = fmt.Sprintf("My name is %v and my age is %v", name, numerOne)

	fmt.Println("The saved string is:", str)*/

	// arrays
	/*var ages [3]int = [3]int{10, 20, 30}
	fmt.Println(ages, len(ages))
	ages[1] = 25
	fmt.Println(ages, len(ages))

	var names = [3]string{"mahfuz", "shishir", "farazi"}
	fmt.Println(names, len(names))

	numbers := [4]int{1, 2, 3, 4}
	fmt.Println(numbers, len(numbers))*/

	// slices (use arrays under the hood)
	/*var scores = []int{10, 20, 30}
	fmt.Println(scores, len(scores))
	scores[2] = 40;
	fmt.Println(scores, len(scores))

	scores = append(scores, 50, 60)
	fmt.Println(scores, len(scores))*/

	// slice ranges
	/*var names = [5]string{"mahfuz", "shishir", "farazi", "zinat", "anni"}
	fmt.Println(names, len(names))

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)*/

	/*var gretting string = "hello there friends!"

	fmt.Println(strings.Contains(gretting, "hello"))
	fmt.Println(strings.ReplaceAll(gretting, "hello", "hi"))
	fmt.Println(strings.ToUpper(gretting))
	fmt.Println(strings.Index(gretting, "th"))

	var counts = []string{}
	counts = strings.Split(gretting, " ")
	fmt.Println(counts, len(counts))*/

	ages := []int{23, 45, 12, 55, 21, 65, 34, 33}
	sort.Ints(ages)

	fmt.Println(ages)

	index := sort.SearchInts(ages, 2443)
	fmt.Println(index)
}
