package main

import (
	"fmt"
	"go_in_action-practice/chapter_4/challenge1/processor"
)

func main() {
	var ids []processor.ID

	ids = append(ids, processor.ID(101), processor.ID(21), processor.ID(11), processor.ID(2))

	fmt.Println(processor.DisplayBatches(ids, 1))
}
