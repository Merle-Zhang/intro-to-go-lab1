package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	return append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = f(slice[i])
	}
}

func mapArray(f func(a int) int, array [5]int) [5]int {
	for i := 0; i < len(array); i++ {
		array[i] = f(array[i])
	}
	return array
}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)
	intsArray := [5]int{1, 2, 3, 4, 5}
	intsArray = mapArray(addOne, intsArray)
	fmt.Println(intsArray)

	fmt.Println()

	newSlice := intsSlice[1:3]
	fmt.Println(intsSlice)
	fmt.Println(newSlice)
	mapSlice(square, newSlice)
	fmt.Println(intsSlice)
	fmt.Println(newSlice)

	fmt.Println()

	fmt.Println(intsSlice)
	intsSlice = double(intsSlice)
	fmt.Println(intsSlice)
}

// Explain how arrays and slices are passed to functions.
// arrays: pass by value
// slices: pass by reference
// Explain how append() works.
// return a new reference address, not modify the original one
// Give use cases for arrays and slices.
// arrays: fixed length
// slices: dynamic length
