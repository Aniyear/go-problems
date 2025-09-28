package main

import (
	"fmt"
	"piscine/piscine"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(piscine.CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(piscine.CanJump(input2))

	input3 := []uint{0}
	fmt.Println(piscine.CanJump(input3))

	input4 := []uint{1,1,1,1,0}
	fmt.Println(piscine.CanJump(input4))
}
