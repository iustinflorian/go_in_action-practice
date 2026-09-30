package main

import "fmt"

type Player struct {
	Score int
	Level int
}

func main() {
	lobby := make(map[string]*Player)
	lobby["alex"] = &Player{Score: 10, Level: 1}

	// ❌ COMPILE ERROR: cannot assign to struct field lobby["alex"].Score in map
	lobby["alex"].Score += 5

	//lobby["alex"] = Player{Score: lobby["alex"].Score + 5, Level: 1}

	fmt.Println(lobby["alex"])
}