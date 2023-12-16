package main

import (
	"fmt"
	"strings"
)   

func main() {

	height := 20
	 trunk := "|||"
	 message:= " Merry Christmas!"
	max_width := find_max_tree_width(height)

	for width := 1; width <= max_width; width += 2 {
		spaces_row := max_width - width
		spaces_on_each_side := spaces_row / 2
		fmt.Printf("%s%s%s\n",
			strings.Repeat(" ", spaces_on_each_side),
			strings.Repeat("*", width),
			strings.Repeat(" ", spaces_on_each_side),
		)
	}

	spaces_row := max_width - 3
	spaces_on_each_side := spaces_row / 2
	fmt.Printf("%s%s%s\n",
		strings.Repeat(" ", spaces_on_each_side),
		trunk,
		strings.Repeat(" ", spaces_on_each_side),
	)
    
	fmt.Println()

	spaces_row = max_width - len(message)
	spaces_on_each_side = spaces_row / 2
	fmt.Printf("%s%s%s\n",
		strings.Repeat(" ", spaces_on_each_side),
		message,
		strings.Repeat(" ", spaces_on_each_side),
	)

	fmt.Scanln()
}
func find_max_tree_width(height int) int {

	width := 1
	for i := 1; i < height; i++ {
		width += 2
	}

	return width
}