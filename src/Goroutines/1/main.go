package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	fmt.Println("before go f(0)")
	go f(0)
	fmt.Println("after go f(0)")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Dragana")
}
