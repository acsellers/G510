package main

var activeMediaKeys = map[Key]bool{}

func HandleMediaKey(keys byte) {
	// next
	if keys&1 != 0 {
		DoKeyDown(MediaNext)
		activeMediaKeys[MediaNext] = true
	} else if activeMediaKeys[MediaNext] {
		DoKeyUp(MediaNext)
		activeMediaKeys[MediaNext] = false
	}

	// back
	if keys&2 != 0 {
		DoKeyDown(MediaBack)
		activeMediaKeys[MediaBack] = true
	} else if activeMediaKeys[MediaBack] {
		DoKeyUp(MediaBack)
		activeMediaKeys[MediaBack] = false
	}

	// stop
	if keys&4 != 0 {
		DoKeyDown(MediaStop)
		activeMediaKeys[MediaStop] = true
	} else if activeMediaKeys[MediaStop] {
		DoKeyUp(MediaStop)
		activeMediaKeys[MediaStop] = false
	}

	// play/pause
	if keys&8 != 0 {
		DoKeyDown(MediaPlayPause)
		activeMediaKeys[MediaPlayPause] = true
	} else if activeMediaKeys[MediaPlayPause] {
		DoKeyUp(MediaPlayPause)
		activeMediaKeys[MediaPlayPause] = false
	}

	// mute
	if keys&16 != 0 {
		DoKeyDown(MediaMute)
		activeMediaKeys[MediaMute] = true
	} else if activeMediaKeys[MediaMute] {
		DoKeyUp(MediaMute)
		activeMediaKeys[MediaMute] = false
	}

	// vol up
	if keys&32 != 0 {
		DoKeyDown(MediaVolUp)
		activeMediaKeys[MediaVolUp] = true
	} else if activeMediaKeys[MediaVolUp] {
		DoKeyUp(MediaVolUp)
		activeMediaKeys[MediaVolUp] = false
	}

	// vol down
	if keys&64 != 0 {
		DoKeyDown(MediaVolDown)
		activeMediaKeys[MediaVolDown] = true
	} else if activeMediaKeys[MediaVolDown] {
		DoKeyUp(MediaVolDown)
		activeMediaKeys[MediaVolDown] = false
	}
}
