package main

import "fmt"

type User struct {
	Name  string
	Score int
}

func modifyMap(m map[string]int) {
	m["Bob"] = 99
	delete(m, "Alice")
}

func main() {
	// --- STEP 1: nil map behavior ---
	var nilMap map[string]int
	val1 := nilMap["Alice"]
	_, ok := nilMap["Alice"]

	fmt.Println("Q1 - val1:", val1, "ok:", ok)

	// --- STEP 2: hmap pointer behavior ---
	scores := map[string]int{"Alice": 10, "Bob": 20}
	modifyMap(scores)

	fmt.Println("Q2 - scores[\"Alice\"]:", scores["Alice"])
	fmt.Println("Q3 - scores[\"Bob\"]:", scores["Bob"])

	// --- STEP 3: Map of Structs & Addressability ---
	users := map[string]User{
		"admin": {Name: "Alice", Score: 100},
	}

	// Will this line compile?
	// users["admin"].Score = 200
	// (Keep this in mind for Q4!)

	// We re-assign the whole struct instead:
	users["admin"] = User{Name: "Alice", Score: 200}
	fmt.Println("Q4 - admin score:", users["admin"].Score)

	// --- STEP 4: Writing to nil map ---
	var panicMap map[string]int
	// Uncommenting the line below:
	panicMap["Go"] = 100

	fmt.Println("Q5 - Will panicMap[\"Go\"] = 100 compile, panic at runtime, or work fine?")
}
