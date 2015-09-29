package keyboard

import "fmt"

func LookupKey(k Key) (int, bool) {
	return 0, false
}

func RawDoKeyDown(k Key) {
	fmt.Println("KeyDown:", k)
}

func RawDoKeyUp(k Key) {
	fmt.Println("KeyUp:", k)
}
