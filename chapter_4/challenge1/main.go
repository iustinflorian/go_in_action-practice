package main

import (
	"fmt"
	"go_in_action-practice/chapter_4/challenge1/processor"
)

func main() {
	var ids []processor.ID

	ids = append(ids,
		processor.ID(101),
		processor.ID(21),
		processor.ID(11),
		processor.ID(2),
		processor.ID(231),
		processor.ID(111),
		processor.ID(22),
	)

	fmt.Println(processor.DisplayBatches(ids, 3))
}
