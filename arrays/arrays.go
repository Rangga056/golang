package main

import "fmt"

func main() {
	//NOTE Arrays in go can either use var or :=
	var array1 = [3]int{1, 2, 3} // [length]type[data values] length is defined
	var cars = [3]string{"toyota", "nissan", "honda"}
	array2 := [...]int{4, 5, 6, 7, 8} // [...]type[data values] length is inferred

	fmt.Print(array1, "\n", array2, "\n", cars, "\n")
	fmt.Println(cars[0]) // toyota

	//change the value in the array
	cars[1] = "datsun"
	fmt.Println(cars[1])

	//NOTE array initialization
	// default value for int is 0
	numbers1 := [5]int{}
	numbers2 := [5]int{1, 2}

	fmt.Print(numbers1, "\n", numbers2, "\n")

	// you can also initializes specific elements in the array
	numbers3 := [5]int{1: 10, 2: 29}
	fmt.Print(numbers3, "\n") // [0, 10, 29, 0, 0]

	//len() to check the length of an array
	fmt.Println(len(numbers1))

	//NOTE slices
	// myslise = []datatype{}

	// arr1 := [6]int{12, 13, 14, 15, 16, 17}
	// myslice := arr1[2:4] // create a slice from arr1 from the end to 4th elements
	//
	// fmt.Printf("myslice = %v\n", myslice)
	// fmt.Printf("length = %d\n", len(myslice))
	// fmt.Printf("capacity = %d\n", cap(myslice))

	// you can also use make() to create a slice
	//NOTE slice_name := make([]type, length, capacity)

	// myslice1 := make([]int, 5, 10)
	// fmt.Printf("myslice1 = %v\n", myslice1)
	// fmt.Printf("length = %d\n", len(myslice1))
	// fmt.Printf("capacity = %d\n", cap(myslice1))

	// NOTE you can modify slices like you modify an array
	// or you can use append()
	//slice_name = append(slice_name, element1, element2, ...)
	// myslice1 = append(myslice1, 20, 21)
	// fmt.Printf("myslice1 = %v\n", myslice1)
	// fmt.Printf("length = %d\n", len(myslice1))
	// fmt.Printf("capacity = %d\n", cap(myslice1))

	// Append one slice to another slice
	// slice3 = append(slice1, slice2...) ... is necessary
	// myslice1 := []int{1, 2, 3}
	// myslice2 := []int{4, 5, 6}
	// myslice3 := append(myslice1, myslice2...)
	//
	// fmt.Printf("myslice3=%v\n", myslice3)
	// fmt.Printf("length=%d\n", len(myslice3))
	// fmt.Printf("capacity=%d\n", cap(myslice3))

	//NOTE unlike array you can change the length of a slcie
	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice1 := arr1[1:5]                 // Slice array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	//NOTE when using slices go loads all the underlying array into the memmory
	// is better to use copy() if the array is large it will reduce memory used
	// copy(dest, src) copy the source to destination

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
  fmt.Print("\n")
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

}
