package main

import (
	"image/color"
	"log"
	"os"
	"os/signal"

	"github.com/acsellers/g510/keyboard"
)

var colors = []color.NRGBA{
	// G1-6
	{255, 0, 0, 0},     // red
	{0, 0, 255, 0},     // blue
	{0, 255, 0, 0},     // green
	{255, 255, 255, 0}, // white
	{128, 128, 128, 0}, // gray
	{0, 0, 0, 0},       // black

	// G7-12
	{220, 76, 60, 0},  // peach
	{26, 188, 156, 0}, // light green
	{52, 152, 219, 0}, // light blue
	{155, 89, 182, 0}, // purple
	{52, 73, 94, 0},   // dark gray
	{241, 196, 15, 0}, // yellow

	// G13-18
	{230, 126, 34, 0},  // orange
	{236, 240, 241, 0}, // white
	{162, 222, 208, 0}, // something
	{211, 84, 0, 0},    // dark orange
	{255, 128, 255, 0}, // purple2
	{128, 64, 0, 0},    //something2
}

func main() {
	k := keyboard.NewKeyboard()
	err := k.Start()
	if err != nil {
		log.Fatal(err)
	}

	k.SetColor(colors[0])

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
