package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	countries := []string{"bangladesh", "india", "pakistan", "maldivs", "bhutan"}

	for index, value := range countries {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}

		if index > 3 {
			fmt.Println("breaking at pos", index)
			break
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
}