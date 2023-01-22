package main

import "fmt"

func main() {
	// string

	var str string = "Irwan"
	var strTwo = "Syah"
	var strThree string

	strFour := "Still"
	fmt.Println("Go fundamentals - by "+str, strTwo, strThree, strFour)

	// integer

	var ageOne int = 27
	var ageTwo = 27

	var ageThree int
	ageFour := 28

	fmt.Println(ageOne, ageTwo, ageThree, ageFour)

	// bits and memory

	var numOne int8 = 12
	var numTwo int16 = 8937
	var numThree int32 = 984974395
	var numFour int64 = 938934897439439384
	var decimal float32 = 743743.848

	var numberOne uint8 = 255

	fmt.Println(numFour, numOne, numThree, numTwo, numberOne)
	fmt.Println(decimal)

	// Print

	fmt.Print("Hello ")
	fmt.Print("World! \n")

	// Array

	var ages [3]int = [3]int{22, 24, 27} // strict length of array
	var agesShort = [3]int{23, 25, 28}

	fmt.Println(ages, agesShort)

	// slices

	var strArray = []string{"irwan", "hmm", "syah"}

	fmt.Println(strArray, len(strArray))

	fmt.Println(append(strArray, "nambah"), len(strArray)) // append tidak mengubah array awal, kecuali ditampung dalam array baru (panjang array tetap 3)

	appendedStr := append(strArray, "ini append") // seperti ini

	fmt.Println(appendedStr, len(appendedStr))

	// slices range

	rangeOne := appendedStr[1:4]
	// strLocation := appendedStr[4] // error

	fmt.Println(rangeOne)

	// Loop

	i := 0
	for i < 5 {
		fmt.Println("Value of i: ", i)
		i++
	}

	for x := 0; x < 5; x++ {
		fmt.Println("Value of x: ", x)
	}

	loopArrStr := []string{"Hari", "Ini", "Adalah", "Minggu"}

	for j := 0; j < len(loopArrStr); j++ {
		fmt.Println(loopArrStr[j])
	}

	for index, value := range loopArrStr {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}

	for _, value := range loopArrStr {
		fmt.Printf("The value is %v \n", value) // kalo cuma pengen value nya aja. Sama jga kalo pengen indexnya aja berarti value diganti _
	}

	for index := range loopArrStr {
		fmt.Printf("The index is %v \n", index) // kalo cuma pengen value nya aja. Sama jga kalo pengen indexnya aja berarti value diganti _ atau dihapus aja karena argumen terakhir
	}

	// Booleans and Conditional

	ageInt := 27

	fmt.Println(ageInt < 27)
	fmt.Println(ageInt == 27)
	fmt.Println(ageInt > 17)

	if ageInt < 27 {
		fmt.Println("The age is less than 27")
	} else if ageInt > 27 {
		fmt.Println("The age is more than 27")
	} else if ageInt == 27 {
		fmt.Println("The age is 27")
	}

	names := []string{"irwan", "arif", "fandi", "ata", "mahfud"}

	for index, value := range names {
		if index == 1 {
			fmt.Printf("The name at index %v is %v", index, value)
		}
	}

}
