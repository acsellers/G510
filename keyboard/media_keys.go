package keyboard

var activeMediaKeys = map[Key]bool{}

func (k Keyboard) HandleMediaKey(keys byte) {
	// next
	if keys&1 != 0 {
		k.CurrentMode.DoKeyDown(MediaNext)
		activeMediaKeys[MediaNext] = true
	} else if activeMediaKeys[MediaNext] {
		k.CurrentMode.DoKeyUp(MediaNext)
		activeMediaKeys[MediaNext] = false
	}

	// back
	if keys&2 != 0 {
		k.CurrentMode.DoKeyDown(MediaBack)
		activeMediaKeys[MediaBack] = true
	} else if activeMediaKeys[MediaBack] {
		k.CurrentMode.DoKeyUp(MediaBack)
		activeMediaKeys[MediaBack] = false
	}

	// stop
	if keys&4 != 0 {
		k.CurrentMode.DoKeyDown(MediaStop)
		activeMediaKeys[MediaStop] = true
	} else if activeMediaKeys[MediaStop] {
		k.CurrentMode.DoKeyUp(MediaStop)
		activeMediaKeys[MediaStop] = false
	}

	// play/pause
	if keys&8 != 0 {
		k.CurrentMode.DoKeyDown(MediaPlayPause)
		activeMediaKeys[MediaPlayPause] = true
	} else if activeMediaKeys[MediaPlayPause] {
		k.CurrentMode.DoKeyUp(MediaPlayPause)
		activeMediaKeys[MediaPlayPause] = false
	}

	// mute
	if keys&16 != 0 {
		k.CurrentMode.DoKeyDown(MediaMute)
		activeMediaKeys[MediaMute] = true
	} else if activeMediaKeys[MediaMute] {
		k.CurrentMode.DoKeyUp(MediaMute)
		activeMediaKeys[MediaMute] = false
	}

	// vol up
	if keys&32 != 0 {
		k.CurrentMode.DoKeyDown(MediaVolUp)
		activeMediaKeys[MediaVolUp] = true
	} else if activeMediaKeys[MediaVolUp] {
		k.CurrentMode.DoKeyUp(MediaVolUp)
		activeMediaKeys[MediaVolUp] = false
	}

	// vol down
	if keys&64 != 0 {
		k.CurrentMode.DoKeyDown(MediaVolDown)
		activeMediaKeys[MediaVolDown] = true
	} else if activeMediaKeys[MediaVolDown] {
		k.CurrentMode.DoKeyUp(MediaVolDown)
		activeMediaKeys[MediaVolDown] = false
	}
}
