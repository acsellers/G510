package keyboard

import (
	"fmt"
	"image/color"
	"time"

	"github.com/GeertJohan/go.hid"
)

var Debug bool

type Keyboard struct {
	DeviceInfo  *hid.DeviceInfo
	Device      *hid.Device
	CurrentMode Mode
	Modes       [3]Mode
	CurrentApp  App
	Apps        []App

	lastAppChange   time.Time
	activeModifiers map[Key]bool
}

func NewKeyboard() Keyboard {
	k := Keyboard{
		CurrentMode: QwertyMode{},
		Modes: [3]Mode{
			QwertyMode{},
			QwertyMode{},
			QwertyMode{},
		},

		// private
		lastAppChange:   time.Now(),
		activeModifiers: make(map[Key]bool),
	}
	return k
}

func (k Keyboard) SetColor(c color.NRGBA) {
	data := []byte{0x5, c.R, c.G, c.B}
	k.Device.Control(33, 9, 0x305, 1, data)
}

func ignorePath(path string) {
	d, _ := hid.OpenPath(path)
	b := make([]byte, 8)
	for {
		d.Read(b)
	}
}

func (k *Keyboard) Start() error {
	devs, err := hid.Enumerate(0x046d, 0xc22d)
	if err != nil {
		return err
	}
	for i, dev := range devs {
		if i == 0 {
			go ignorePath(dev.Path)
			continue
		}

		k.DeviceInfo = dev
		k.Device, err = hid.OpenPath(dev.Path)
		if err != nil {
			return err
		}

		go k.process()
	}
	return nil
}

func (k *Keyboard) process() {
	last := time.Now()
	for {
		b := make([]byte, 8)
		k.Device.Read(b)
		if Debug {
			fmt.Printf(
				"%s (%s): %v\n",
				k.DeviceInfo.Path,
				time.Now().Sub(last),
				b,
			)
		}
		switch b[0] {
		case 1:
			k.HandleModifierKeys(b[1])
			k.HandleNormalKeys(b[2:])
		case 2:
			k.HandleMediaKey(b[1])
		case 3:
			k.HandleSpecialKeys(b[1:])
		}
		last = time.Now()
	}
}

type Mode interface {
	DoKeyDown(Key)
	DoKeyUp(Key)
	MacroKeyDown(int)
	MacroKeyUp(int)
}
type App interface {
	LCDKeyDown(int)
	LCDKeyUp(int)
}
