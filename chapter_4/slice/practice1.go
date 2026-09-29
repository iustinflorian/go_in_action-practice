package main

import "fmt"

func modifySlice(s *[]int) {
	(*s)[0] = 999
	*s = append(*s, 100)
	(*s)[1] = 888
}

func main() {
	orig := make([]int, 2, 4)
	orig[0] = 10
	orig[1] = 20

	modifySlice(&orig)

	fmt.Println("orig:", orig)
	fmt.Println("len:", len(orig), "cap:", cap(orig))
}
