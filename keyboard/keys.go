package keyboard

type Key uint8

const (
	// letter keys
	KeyA Key = iota + 4
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	Key1 // num row
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	Key0
	KeyEnter // actions
	KeyEsc
	KeyBackspace
	KeyTab
	KeySpace
	KeyMinus
	KeyEqual
	KeyLeftBrace
	KeyRightBrace
	KeyBackslash
	_
	KeySemicolon
	KeyApostrophe
	KeyGrave
	KeyComma
	KeyDot
	KeySlash
	KeyCapsLock
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeySysRq
	KeyScrollLock
	KeyPause
	KeyInsert
	KeyHome
	KeyPageUp
	KeyDelete
	KeyEnd
	KeyPageDown
	KeyRight
	KeyLeft
	KeyDown
	KeyUp
	KeyNumLock
	_ // num pad things
	KeyNPSlash
	KeyNPAsterisk
	KeyNPMinus
	KeyNPPlus
	KeyNPEnter
	KeyNP1
	KeyNP2
	KeyNP3
	KeyNP4
	KeyNP5
	KeyNP6
	KeyNP7
	KeyNP8
	KeyNP9
	KeyNP0
	KeyNPDot
	_ // media keys
	MediaNext
	MediaBack
	MediaStop
	MediaPlayPause
	MediaMute
	MediaVolUp
	MediaVolDown

	// Modifier keys
	LeftCtrl
	LeftShift
	LeftAlt
	LeftMeta
	RightCtrl
	RightShift
	RightAlt
	RightMeta
)
