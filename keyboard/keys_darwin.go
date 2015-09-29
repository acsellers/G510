package keyboard

import "fmt"

func LookupKey(k Key) (int, bool) {
	return 0, false
}

func RawDoKeyDown(k int) {
	fmt.Println("KeyDown:", k)
}

func RawDoKeyUp(k int) {
	fmt.Println("KeyUp:", k)
}
