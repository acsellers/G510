package keyboard

import (
	"fmt"
	"log"

	"github.com/ben-bensdevzone/uinput"
)

var keyMap = map[Key]int{
	KeyA:           uinput.KEY_A,
	KeyB:           uinput.KEY_B,
	KeyC:           uinput.KEY_C,
	KeyD:           uinput.KEY_D,
	KeyE:           uinput.KEY_E,
	KeyF:           uinput.KEY_F,
	KeyG:           uinput.KEY_G,
	KeyH:           uinput.KEY_H,
	KeyI:           uinput.KEY_I,
	KeyJ:           uinput.KEY_J,
	KeyK:           uinput.KEY_K,
	KeyL:           uinput.KEY_L,
	KeyM:           uinput.KEY_M,
	KeyN:           uinput.KEY_N,
	KeyO:           uinput.KEY_O,
	KeyP:           uinput.KEY_P,
	KeyQ:           uinput.KEY_Q,
	KeyR:           uinput.KEY_R,
	KeyS:           uinput.KEY_S,
	KeyT:           uinput.KEY_T,
	KeyU:           uinput.KEY_U,
	KeyV:           uinput.KEY_V,
	KeyW:           uinput.KEY_W,
	KeyX:           uinput.KEY_X,
	KeyY:           uinput.KEY_Y,
	KeyZ:           uinput.KEY_Z,
	Key1:           uinput.KEY_1,
	Key2:           uinput.KEY_2,
	Key3:           uinput.KEY_3,
	Key4:           uinput.KEY_4,
	Key5:           uinput.KEY_5,
	Key6:           uinput.KEY_6,
	Key7:           uinput.KEY_7,
	Key8:           uinput.KEY_8,
	Key9:           uinput.KEY_9,
	Key0:           uinput.KEY_0,
	KeyEnter:       uinput.KEY_ENTER,
	KeyEsc:         uinput.KEY_ESC,
	KeyBackspace:   uinput.KEY_BACKSPACE,
	KeyTab:         uinput.KEY_TAB,
	KeySpace:       uinput.KEY_SPACE,
	KeyMinus:       uinput.KEY_MINUS,
	KeyEqual:       uinput.KEY_EQUAL,
	KeyLeftBrace:   uinput.KEY_LEFTBRACE,
	KeyRightBrace:  uinput.KEY_RIGHTBRACE,
	KeyBackslash:   uinput.KEY_BACKSLASH,
	KeySemicolon:   uinput.KEY_SEMICOLON,
	KeyApostrophe:  uinput.KEY_APOSTROPHE,
	KeyGrave:       uinput.KEY_GRAVE,
	KeyCapsLock:    uinput.KEY_CAPSLOCK,
	KeyComma:       uinput.KEY_COMMA,
	KeyDot:         uinput.KEY_DOT,
	KeySlash:       uinput.KEY_SLASH,
	KeyF1:          uinput.KEY_F1,
	KeyF2:          uinput.KEY_F2,
	KeyF3:          uinput.KEY_F3,
	KeyF4:          uinput.KEY_F4,
	KeyF5:          uinput.KEY_F5,
	KeyF6:          uinput.KEY_F6,
	KeyF7:          uinput.KEY_F7,
	KeyF8:          uinput.KEY_F8,
	KeyF9:          uinput.KEY_F9,
	KeyF10:         uinput.KEY_F10,
	KeyF11:         uinput.KEY_F11,
	KeyF12:         uinput.KEY_F12,
	KeySysRq:       uinput.KEY_SYSRQ,
	KeyScrollLock:  uinput.KEY_SCROLLLOCK,
	KeyPause:       uinput.KEY_PAUSE,
	KeyInsert:      uinput.KEY_INSERT,
	KeyHome:        uinput.KEY_HOME,
	KeyPageUp:      uinput.KEY_PAGEUP,
	KeyDelete:      uinput.KEY_DELETE,
	KeyEnd:         uinput.KEY_END,
	KeyPageDown:    uinput.KEY_PAGEDOWN,
	KeyRight:       uinput.KEY_RIGHT,
	KeyLeft:        uinput.KEY_LEFT,
	KeyDown:        uinput.KEY_DOWN,
	KeyUp:          uinput.KEY_UP,
	KeyNumLock:     uinput.KEY_NUMLOCK,
	KeyNPSlash:     uinput.KEY_KPSLASH,
	KeyNPAsterisk:  uinput.KEY_KPASTERISK,
	KeyNPMinus:     uinput.KEY_KPMINUS,
	KeyNPPlus:      uinput.KEY_KPPLUS,
	KeyNPEnter:     uinput.KEY_KPENTER,
	KeyNP1:         uinput.KEY_KP1,
	KeyNP2:         uinput.KEY_KP2,
	KeyNP3:         uinput.KEY_KP3,
	KeyNP4:         uinput.KEY_KP4,
	KeyNP5:         uinput.KEY_KP5,
	KeyNP6:         uinput.KEY_KP6,
	KeyNP7:         uinput.KEY_KP7,
	KeyNP8:         uinput.KEY_KP8,
	KeyNP9:         uinput.KEY_KP9,
	KeyNP0:         uinput.KEY_KP0,
	KeyNPDot:       uinput.KEY_KPDOT,
	MediaNext:      uinput.KEY_NEXTSONG,
	MediaBack:      uinput.KEY_PREVIOUSSONG,
	MediaStop:      uinput.KEY_STOPCD,
	MediaPlayPause: uinput.KEY_PLAYPAUSE,
	MediaMute:      uinput.KEY_MUTE,
	MediaVolUp:     uinput.KEY_VOLUMEUP,
	MediaVolDown:   uinput.KEY_VOLUMEDOWN,
	// Modifier keys
	LeftCtrl:   uinput.KEY_LEFTCTRL,
	LeftShift:  uinput.KEY_LEFTSHIFT,
	LeftAlt:    uinput.KEY_LEFTALT,
	LeftMeta:   uinput.KEY_LEFTMETA,
	RightCtrl:  uinput.KEY_RIGHTCTRL,
	RightShift: uinput.KEY_RIGHTSHIFT,
	RightAlt:   uinput.KEY_RIGHTALT,
	RightMeta:  uinput.KEY_RIGHTMETA,
}

var kboard uinput.VKeyboard

func init() {
	err := kboard.Create("/dev/uinput")
	if err != nil {
		log.Fatal(err)
	}
}

func LookupKey(k Key) (int, bool) {
	ki, ok := keyMap[k]
	return ki, ok
}

func RawDoKeyDown(k int) {
	if Debug {
		fmt.Println("KeyDown:", k)
	}
	kboard.SendKeyPress(k)
}

func RawDoKeyUp(k int) {
	if Debug {
		fmt.Println("KeyUp:", k)
	}
	kboard.SendKeyRelease(k)
}
