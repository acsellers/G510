package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/GeertJohan/go.hid"
	"github.com/kylelemons/gousb/usb"
)

var (
	Dev  *usb.Device
	Keys usb.Endpoint
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

func SetColor(c color.NRGBA) {
	data := []byte{0x05, c.R, c.G, c.B}
	Dev.Control(33, 9, 0x305, 1, data)
}

func IgnorePath(path string) {
	d, _ := hid.OpenPath(path)
	b := make([]byte, 8)
	for {
		d.Read(b)
	}
}
func main() {
	devs, err := hid.Enumerate(0x046d, 0xc22d)
	if err != nil {
		log.Fatal(err)
	}
	for i, dev := range devs {
		if i == 0 {
			go IgnorePath(dev.Path)
			continue
		}

		d, err := hid.OpenPath(dev.Path)
		if err != nil {
			log.Fatal(dev.Path, err)
		}

		last := time.Now()
		for {
			b := make([]byte, 8)
			d.Read(b)
			fmt.Printf(
				"%s (%s): %v\n",
				dev.Path,
				time.Now().Sub(last),
				b,
			)
			switch b[0] {
			case 1:
				HandleNormalKeys(b[2:])
			case 2:
				HandleMediaKey(b[1])
			case 3:
				HandleSpecialKeys(b[1:])
			}
			last = time.Now()
		}
	}
}
