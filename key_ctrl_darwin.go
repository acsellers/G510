package main

import "fmt"

func DoKeyDown(k Key) {
	fmt.Println("KeyDown:", k)
}

func DoKeyUp(k Key) {
	fmt.Println("KeyUp:", k)
}
