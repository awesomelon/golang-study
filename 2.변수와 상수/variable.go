package main

import "fmt"

func main() {
	var a int = 1
	a += 2
	fmt.Println(a)

	const b int = 10
	fmt.Println(b)

	const c string = "Hello"
	const d string = "World"

	fmt.Println(`` + c + ` ` + d + `!`)
}