package main

import "fmt"

func main(){

	array := [10] int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println(array)

	slice_1 := array[3:8]
	fmt.Println("slice_1 = ",slice_1)

	fmt.Println()
	fmt.Println("The length of slice_1 = ",len(slice_1))
	fmt.Println("The capacity of slice_1 = ",cap(slice_1))

	fmt.Println()
	slice_1_1 := slice_1[2:4]
	fmt.Println("slice_1_1 = ",slice_1_1)
	fmt.Println("The length of slice_1 = ",len(slice_1_1))
	fmt.Println("The capacity of slice_1 = ",cap(slice_1_1))
}