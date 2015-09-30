package keyboard

/*
	#cgo pkg-config: x11 gdk-3.0 gtk+-3.0
  #include <X11/XKBlib.h>
	#include <gdk/gdkx.h>
	#include <gtk/gtk.h>
*/
import "C"

import "time"

func (k *Keyboard) initLEDWatch() {
	C.gtk_init(nil, nil)
	win := C.gdk_x11_get_default_xdisplay()
	var state C.uint
	var lstate byte
	for range time.Tick(16 * time.Millisecond) {
		lstate = 0

		// Caps: 0, Num: 1, Scroll: 2
		C.XkbGetIndicatorState(win, C.XkbUseCoreKbd, &state)

		if 0x4&state != 0 {
			lstate += 0x4
		}
		if 0x2&state != 0 {
			lstate += 1
		}
		if 0x1&state != 0 {
			lstate += 2
		}

		// Num: 0, Caps: 1, Scroll: 2
		k.device0.Control(0x21, 0x09, 0x0200, 0, []byte{lstate})
	}
}
