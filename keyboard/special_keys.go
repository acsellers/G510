package keyboard

import (
	"bytes"
	"encoding/binary"
	"time"
)

var activeSpecials = map[int32]bool{}

func (k *Keyboard) HandleSpecialKeys(keys []byte) {
	var keyMask int32
	binary.Read(bytes.NewReader(keys), binary.LittleEndian, &keyMask)
	// byte 1
	if keyMask&SpG1 != 0 {
		activeSpecials[SpG1] = true
		ignoreKeys[KeyF1] = true
		k.CurrentMode.MacroKeyUp(1)
	} else if activeSpecials[SpG1] {
		activeSpecials[SpG1] = false
		ignoreKeys[KeyF1] = false
		k.CurrentMode.MacroKeyDown(1)
	}
	if keyMask&SpG2 != 0 {
		activeSpecials[SpG2] = true
		ignoreKeys[KeyF2] = true
		k.CurrentMode.MacroKeyUp(2)
	} else if activeSpecials[SpG2] {
		activeSpecials[SpG2] = false
		ignoreKeys[KeyF2] = false
		k.CurrentMode.MacroKeyDown(2)
	}
	if keyMask&SpG3 != 0 {
		activeSpecials[SpG3] = true
		ignoreKeys[KeyF3] = true
		k.CurrentMode.MacroKeyUp(3)
	} else if activeSpecials[SpG3] {
		activeSpecials[SpG3] = false
		ignoreKeys[KeyF3] = false
		k.CurrentMode.MacroKeyDown(3)
	}
	if keyMask&SpG4 != 0 {
		activeSpecials[SpG4] = true
		ignoreKeys[KeyF4] = true
		k.CurrentMode.MacroKeyUp(4)
	} else if activeSpecials[SpG4] {
		activeSpecials[SpG4] = false
		ignoreKeys[KeyF4] = false
		k.CurrentMode.MacroKeyDown(4)
	}
	if keyMask&SpG5 != 0 {
		activeSpecials[SpG5] = true
		ignoreKeys[KeyF5] = true
		k.CurrentMode.MacroKeyUp(5)
	} else if activeSpecials[SpG5] {
		activeSpecials[SpG5] = false
		ignoreKeys[KeyF5] = false
		k.CurrentMode.MacroKeyDown(5)
	}
	if keyMask&SpG6 != 0 {
		activeSpecials[SpG6] = true
		ignoreKeys[KeyF6] = true
		k.CurrentMode.MacroKeyUp(6)
	} else if activeSpecials[SpG6] {
		activeSpecials[SpG6] = false
		ignoreKeys[KeyF6] = false
		k.CurrentMode.MacroKeyDown(6)
	}
	if keyMask&SpG7 != 0 {
		activeSpecials[SpG7] = true
		ignoreKeys[KeyF7] = true
		k.CurrentMode.MacroKeyUp(7)
	} else if activeSpecials[SpG7] {
		activeSpecials[SpG7] = false
		ignoreKeys[KeyF7] = false
		k.CurrentMode.MacroKeyDown(7)
	}
	if keyMask&SpG8 != 0 {
		activeSpecials[SpG8] = true
		ignoreKeys[KeyF8] = true
		k.CurrentMode.MacroKeyUp(8)
	} else if activeSpecials[SpG8] {
		activeSpecials[SpG8] = false
		ignoreKeys[KeyF8] = false
		k.CurrentMode.MacroKeyDown(8)
	}

	// byte 2
	if keyMask&SpG9 != 0 {
		activeSpecials[SpG9] = true
		ignoreKeys[KeyF9] = true
		k.CurrentMode.MacroKeyUp(9)
	} else if activeSpecials[SpG9] {
		activeSpecials[SpG9] = false
		ignoreKeys[KeyF9] = false
		k.CurrentMode.MacroKeyDown(9)
	}
	if keyMask&SpG10 != 0 {
		activeSpecials[SpG10] = true
		ignoreKeys[KeyF10] = true
		k.CurrentMode.MacroKeyUp(10)
	} else if activeSpecials[SpG10] {
		activeSpecials[SpG10] = false
		ignoreKeys[KeyF10] = false
		k.CurrentMode.MacroKeyDown(10)
	}
	if keyMask&SpG11 != 0 {
		activeSpecials[SpG11] = true
		ignoreKeys[KeyF11] = true
		k.CurrentMode.MacroKeyUp(11)
	} else if activeSpecials[SpG11] {
		activeSpecials[SpG11] = false
		ignoreKeys[KeyF11] = false
		k.CurrentMode.MacroKeyDown(11)
	}
	if keyMask&SpG12 != 0 {
		activeSpecials[SpG12] = true
		ignoreKeys[KeyF12] = true
		k.CurrentMode.MacroKeyUp(12)
	} else if activeSpecials[SpG12] {
		activeSpecials[SpG12] = false
		ignoreKeys[KeyF12] = false
		k.CurrentMode.MacroKeyDown(12)
	}

	if keyMask&SpG13 != 0 {
		activeSpecials[SpG13] = true
		ignoreKeys[Key1] = true
		k.CurrentMode.MacroKeyUp(13)
	} else if activeSpecials[SpG13] {
		activeSpecials[SpG13] = false
		ignoreKeys[Key1] = false
		k.CurrentMode.MacroKeyDown(13)
	}
	if keyMask&SpG14 != 0 {
		activeSpecials[SpG14] = true
		ignoreKeys[Key2] = true
		k.CurrentMode.MacroKeyUp(14)
	} else if activeSpecials[SpG14] {
		activeSpecials[SpG14] = false
		ignoreKeys[Key2] = false
		k.CurrentMode.MacroKeyDown(14)
	}
	if keyMask&SpG15 != 0 {
		activeSpecials[SpG15] = true
		ignoreKeys[Key3] = true
		k.CurrentMode.MacroKeyUp(15)
	} else if activeSpecials[SpG15] {
		activeSpecials[SpG15] = false
		ignoreKeys[Key3] = false
		k.CurrentMode.MacroKeyDown(15)
	}
	if keyMask&SpG16 != 0 {
		activeSpecials[SpG16] = true
		ignoreKeys[Key4] = true
		k.CurrentMode.MacroKeyUp(16)
	} else if activeSpecials[SpG16] {
		activeSpecials[SpG16] = false
		ignoreKeys[Key4] = false
		k.CurrentMode.MacroKeyDown(16)
	}

	// byte 3
	if keyMask&SpG17 != 0 {
		activeSpecials[SpG17] = true
		ignoreKeys[Key5] = true
		k.CurrentMode.MacroKeyUp(17)
	} else if activeSpecials[SpG17] {
		activeSpecials[SpG17] = false
		ignoreKeys[Key5] = false
		k.CurrentMode.MacroKeyDown(17)
	}
	if keyMask&SpG18 != 0 {
		activeSpecials[SpG18] = true
		ignoreKeys[Key6] = true
		k.CurrentMode.MacroKeyUp(18)
	} else if activeSpecials[SpG18] {
		activeSpecials[SpG18] = false
		ignoreKeys[Key6] = false
		k.CurrentMode.MacroKeyDown(18)
	}

	// TODO MetaOff
	if keyMask&SpMetaOff != 0 {
		activeSpecials[SpMetaOff] = true
	} else if activeSpecials[SpMetaOff] {
		activeSpecials[SpMetaOff] = false
	}

	// Light is
	if keyMask&SpLight != 0 {
		activeSpecials[SpLight] = true
	} else if activeSpecials[SpLight] {
		activeSpecials[SpLight] = false
	}

	if keyMask&SpM1 != 0 {
		activeSpecials[SpM1] = true
	} else if activeSpecials[SpM1] {
		activeSpecials[SpM1] = false
		if k.Modes[0] != nil {
			k.CurrentMode = k.Modes[0]
		}
	}

	if keyMask&SpM2 != 0 {
		activeSpecials[SpM2] = true
	} else if activeSpecials[SpM2] {
		activeSpecials[SpM2] = false
		if k.Modes[1] != nil {
			k.CurrentMode = k.Modes[1]
		}
	}

	// byte 4
	if keyMask&SpM3 != 0 {
		activeSpecials[SpM3] = true
	} else if activeSpecials[SpM3] {
		activeSpecials[SpM3] = false
		if k.Modes[2] != nil {
			k.CurrentMode = k.Modes[2]
		}
	}

	// TODO MR
	if keyMask&SpMR != 0 {
		activeSpecials[SpMR] = true
	} else if activeSpecials[SpMR] {
		activeSpecials[SpMR] = false
	}

	if keyMask&SpLCD_R != 0 {
		activeSpecials[SpLCD_R] = true
	} else if activeSpecials[SpLCD_R] {
		activeSpecials[SpLCD_R] = false
		// swap to next app
		if len(k.Apps) > 1 && k.lastAppChange.Sub(time.Now()) > time.Second/3 {
			k.lastAppChange = time.Now()
			if k.Apps[0] == k.CurrentApp {
				k.CurrentApp = k.Apps[len(k.Apps)-1]
			} else {
				for i := 1; i < len(k.Apps); i++ {
					if k.Apps[i-1] == k.CurrentApp {
						k.CurrentApp = k.Apps[i]
					}
				}
			}
		}
	}

	if keyMask&SpLCD_1 != 0 {
		activeSpecials[SpLCD_1] = true
		if k.CurrentApp != nil {
			k.CurrentApp.LCDKeyDown(1)
		}
	} else if activeSpecials[SpLCD_1] {
		activeSpecials[SpLCD_1] = false
		if k.CurrentApp != nil {
			k.CurrentApp.LCDKeyUp(1)
		}
	}

	if keyMask&SpLCD_2 != 0 {
		activeSpecials[SpLCD_2] = true
		if k.CurrentApp != nil {
			k.CurrentApp.LCDKeyDown(2)
		}
	} else if activeSpecials[SpLCD_2] {
		activeSpecials[SpLCD_2] = false
		if k.CurrentApp != nil {
			k.CurrentApp.LCDKeyUp(2)
		}
	}

	if keyMask&SpLCD_3 != 0 {
		activeSpecials[SpLCD_3] = true
		if k.CurrentApp != nil {
			k.CurrentApp.LCDKeyDown(3)
		}
	} else if activeSpecials[SpLCD_3] {
		activeSpecials[SpLCD_3] = false
		if k.CurrentApp != nil {
			k.CurrentApp.LCDKeyUp(3)
		}
	}

	if keyMask&SpLCD_4 != 0 {
		activeSpecials[SpLCD_4] = true
		if k.CurrentApp != nil {
			k.CurrentApp.LCDKeyDown(4)
		}
	} else if activeSpecials[SpLCD_4] {
		activeSpecials[SpLCD_4] = false
		if k.CurrentApp != nil {
			k.CurrentApp.LCDKeyUp(4)
		}
	}

	// TODO MuteHead
	if keyMask&SpMuteHead != 0 {
		activeSpecials[SpMuteHead] = true
	} else if activeSpecials[SpMuteHead] {
		activeSpecials[SpMuteHead] = false
	}

	// TODO MuteMic
	if keyMask&SpMuteMic != 0 {
		activeSpecials[SpMuteMic] = true
	} else if activeSpecials[SpMuteMic] {
		activeSpecials[SpMuteMic] = false
	}

}

const (
	SpG1       int32 = 0x1
	SpG2             = 0x2
	SpG3             = 0x4
	SpG4             = 0x8
	SpG5             = 0x10
	SpG6             = 0x20
	SpG7             = 0x40
	SpG8             = 0x80
	SpG9             = 0x100
	SpG10            = 0x200
	SpG11            = 0x400
	SpG12            = 0x800
	SpG13            = 0x1000
	SpG14            = 0x2000
	SpG15            = 0x4000
	SpG16            = 0x8000
	SpG17            = 0x10000
	SpG18            = 0x20000
	SpMetaOff        = 0x40000
	SpLight          = 0x80000
	SpM1             = 0x100000
	SpM2             = 0x200000
	SpM3             = 0x400000
	SpMR             = 0x800000
	SpLCD_R          = 0x1000000
	SpLCD_1          = 0x2000000
	SpLCD_2          = 0x4000000
	SpLCD_3          = 0x8000000
	SpLCD_4          = 0x10000000
	SpMuteHead       = 0x20000000
	SpMuteMic        = 0x40000000
)
