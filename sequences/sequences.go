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
	fmt.Println(array)

}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	intsArray := [3]int{1, 2, 3}
	mapArray(addOne, intsArray)
	fmt.Println(intsArray)

	//arrays are copied when passed to a function and have a fixed length
	//slices refer to the underlying elements and have variable lengths

	newSlice := intsSlice[1:3]
	mapSlice(square, newSlice)
	// newSlice is not a completely separate copy. It is a view onto part of the same underlying array used by intSlice
	// therefore the part that is represents in intSlice also gets changed
	fmt.Println(intsSlice)
	fmt.Println(newSlice)

	intsSlice = double(intsSlice)
	fmt.Println(intsSlice)
}

/*3f:
Arrays are passed by value: when you pass an array to a function, Go makes a copy of the whole array.
If the function changes that array, the original array outside the function is unchanged unless you pass a pointer to it.
Arrays are useful when the size is fixed and meaningful.

Slices are also passed by value, but the value being copied is a small slice header containing a pointer to an underlying array,
plus its length and capacity. Because the copied slice still refers to the same underlying array, changing an element such as
slice[i] inside a function can change the original data.
Slices are more useful for most collections where the number of elements can vary.
They are commonly used when adding/removing items, passing collections around functions, or working with parts of another collection.

append() adds elements to the end of a slice. It returns a new slice because the slice's length changes. If there is enough
capacity, it may reuse the same underlying array. If there is not enough capacity, Go allocates a new underlying array and
copies the data over. Because of this, you should usually write:
slice = append(slice, value)
and if this happens inside a function, return the new slice to the caller.
*/
