package main

import "fmt"

func whileloop() {
	count := 1
	for count < 5 {
		count += count
	}
	fmt.Println(count) // 8
}
