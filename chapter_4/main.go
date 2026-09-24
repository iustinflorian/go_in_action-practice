package main

import "fmt"

func show(array []int) {
	for _, val := range array {
		fmt.Println(val)
	}
}

func update(array *[]int, idx int, val int) {
	(*array)[idx] = val
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice := array[1:5]

	update(&array, 1, 5)
	show(slice)

	//array1 := [5]*int{0: new(int), 1: new(int)}
	//*array1[0] = 10
	//*array1[1] = 20
	//
	//array2 := array1
	//val := 33
	//array2[1] = &val
	//array2[0] = new(int)
	//*array2[0] = 1
	//
	//for _, value := range array1 {
	//	if value != nil {
	//		fmt.Println(*value)
	//	} else {
	//		fmt.Println("nil")
	//	}
	//}
	//
	//for _, value := range array2 {
	//	if value != nil {
	//		fmt.Println(*value)
	//	} else {
	//		fmt.Println("nil")
	//	}
	//}
}
