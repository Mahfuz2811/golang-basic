package main

import (
	"fmt"
	"math"
)

func sayGreetings(name string) {
	fmt.Println("Good morning", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}

func countryName(names []string, funcName func(string)) {
	for _, name := range names {
		funcName(name)
	}
}

func calculateCircleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreetings("Mahfuz")
	// sayGreetings("Shishir")
	// sayGoodbye("Farazi")

	countryName([]string{"bangladesh", "india", "pakishtan"}, sayGreetings)
	countryName([]string{"bangladesh", "india", "pakishtan"}, sayGoodbye)

	area1 := calculateCircleArea(10.5)
	area2 := calculateCircleArea(20.5)

	fmt.Println(area1, area2)
	fmt.Printf("The area1 is %0.3f and the area2 is %0.3f \n", area1, area2)
	
}