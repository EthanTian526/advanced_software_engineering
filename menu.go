package main

import "fmt"

var command string

func main() {
	fmt.Println("Please input a number, 1 for 'helloworld':")
	for true {
		fmt.Scanln(&command)
		if command == "1" {
			fmt.Println("hello world")
		} else if command == "exit" {
			break
		} else {
			fmt.Printf("\"%s\" cannot be recongnized ,please input again.", command)
			fmt.Println()
		}
	}
}
