package main

import "fmt"

func mul(a1, a2 int) int {
	res := a1 * a2
	fmt.Println("Result: ", res)
	return 0
}

func show() {
	fmt.Println("Hello")
}

func add(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Result: ", res)
	return 0

}

func main() {
	mul(1, 1)

	defer func() {
		fmt.Println("Defer")
	}()

	defer mul(1, 2)

	show()

	fmt.Println("Start")

	defer func() {
		fmt.Println("Defer 2")
	}()

	defer fmt.Println("3")
	defer add(2, 2)
	defer add(1, 4)

}
