package main

import (
	"fmt"
	"math"
)

func greetingName(n string) {
	fmt.Printf("Welcome home %v \n", n)
}

func circleAre(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	name := "Irwansyah"
	greetingName(name)

	area := circleAre(25.5)
	areaTwo := circleAre(10.5)
	fmt.Println(area, areaTwo)
	fmt.Printf("%0.3f and %0.3f", area, areaTwo)
}
