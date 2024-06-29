package main

import "fmt"

func main() {
	fruits := [3]string{"apple", "mple", "joker"}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	for i, val := range fruits {
		fmt.Println(i, val)
	}
}
