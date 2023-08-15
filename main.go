package main

import "fmt"

func main() {
	/*x := 0
	for x < 5 {
		fmt.Println("The value of x is:", x)
		x++
	}*/

	/*for i := 0; i < 5; i++ {
		fmt.Println("The value of i is:", i)
	}*/

	/*names := []string{"bangladesh", "india", "pakistan", "nepal", "bhutan"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}*/

	// for in
	/*names := []string{"bangladesh", "india", "pakistan", "nepal", "bhutan"}

	for index, value := range names {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}*/

	// if we want only value then
	names := []string{"bangladesh", "india", "pakistan", "nepal", "bhutan"}

	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
	}

	for index, _ := range names {
		fmt.Printf("The index is %v \n", index)
	}

}