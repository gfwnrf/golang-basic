package main

import "fmt"

func main() {
	var arrayOfInt = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrayOfInt)
	fmt.Println("Length of array 1 = ", len(arrayOfInt))

	var arrayOfInt2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arrayOfInt2)
	fmt.Println("Length of array 2 = ", len(arrayOfInt2))

	arrayOfInt3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrayOfInt3)
	fmt.Println("Length of array 3 = ", len(arrayOfInt3))

	arrayOfInt4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arrayOfInt4)
	fmt.Println("Length of array 4 = ", len(arrayOfInt4))

	sliceInt := []int{}
	fmt.Println(sliceInt)
	sliceInt = append(sliceInt, 20, 21)
	fmt.Println(sliceInt)

}
