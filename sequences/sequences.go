package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	for i, value := range slice {
		slice[i] = f(value)
	}
}

func mapArray(f func(a int) int, array [3]int) {
	for i, value := range array {
		array[i] = f(value)
		//changes only the copied array. the original ints in main stay unchanged
		//this is different to the slice version, because slices refer to underlying data
	}

}

func main() {
	intsSlice := []int{1, 2, 3}
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	intsArray := [3]int{1, 2, 3}
	mapArray(addOne, intsArray)
	fmt.Println(intsArray)

	//arrays are copied when passed to a function and have a fixed length
	//slices refer to the underlying elements and have variable lengths
}
