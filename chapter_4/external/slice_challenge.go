package main

import "fmt"

func modify1(s []int) {
	s[0] = 99
	s = append(s, 100)
	s[1] = 88
}

func modify2(s *[]int) {
	*s = append(*s, 200)
	(*s)[0] = 77
}

func main() {
	// Step 1: Initialization
	base := make([]int, 2, 4)
	base[0] = 10
	base[1] = 20

	// Step 2: Sub-slicing
	sub := base[1:2]

	// Step 3: Pass by value + append
	modify1(base)

	// Step 4: Print check #1
	fmt.Println("Q1 - base:", base)
	fmt.Println("Q2 - sub: ", sub)

	// Step 5: Triggering reallocation
	base = append(base, 30)
	base = append(base, 40)
	base = append(base, 50) // What happens to capacity here?

	// Step 6: Pass by pointer
	modify2(&base)

	// Step 7: Print check #2
	fmt.Println("Q3 - base:", base)

	// Step 8: Mutating old memory
	sub[0] = 555
	fmt.Println("Q4 - base:", base)
	fmt.Println("Q5 - sub: ", sub)
}
