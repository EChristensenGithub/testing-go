package main

import "fmt"

func main() {
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}
	for _, v := range statePopulations {
		fmt.Println(v)
	}

	a := [4]int{1, 2, 3, 4} // "a" has type [4]int (array of 4 ints)
	x := a[0:]

	fmt.Println(x)

	buttons := [][]string{
		{"Some", "Test", "Buttons"},
		{"Another", "Row"},
	}

	fmt.Printf("The whole array: %v\n", buttons)
	fmt.Printf("First member of the array: %v\n", buttons[0])
	fmt.Printf("Second member of the array: %v\n", buttons[1])

	h := []int{1, 2, 3}
	fmt.Println(h)

	var s []int
	s = append(s, 0)
	fmt.Println(s)
}
