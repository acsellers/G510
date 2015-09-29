package keyboard

func (k *Keyboard) HandleModifierKeys(keys byte) {
	if keys&1 != 0 {
		k.activeModifiers[LeftCtrl] = true
		k.CurrentMode.DoKeyDown(LeftCtrl)
	} else if k.activeModifiers[LeftCtrl] {
		k.activeModifiers[LeftCtrl] = false
		k.CurrentMode.DoKeyUp(LeftCtrl)
	}

	if keys&2 != 0 {
		k.activeModifiers[LeftShift] = true
		k.CurrentMode.DoKeyDown(LeftShift)
	} else if k.activeModifiers[LeftShift] {
		k.activeModifiers[LeftShift] = false
		k.CurrentMode.DoKeyUp(LeftShift)
	}

	if keys&4 != 0 {
		k.activeModifiers[LeftAlt] = true
		k.CurrentMode.DoKeyDown(LeftAlt)
	} else if k.activeModifiers[LeftAlt] {
		k.activeModifiers[LeftAlt] = false
		k.CurrentMode.DoKeyUp(LeftAlt)
	}

	if keys&8 != 0 {
		k.activeModifiers[LeftMeta] = true
		k.CurrentMode.DoKeyDown(LeftMeta)
	} else if k.activeModifiers[LeftMeta] {
		k.activeModifiers[LeftMeta] = false
		k.CurrentMode.DoKeyUp(LeftMeta)
	}

	if keys&16 != 0 {
		k.activeModifiers[RightCtrl] = true
		k.CurrentMode.DoKeyDown(RightCtrl)
	} else if k.activeModifiers[RightCtrl] {
		k.activeModifiers[RightCtrl] = false
		k.CurrentMode.DoKeyUp(RightCtrl)
	}

	if keys&32 != 0 {
		k.activeModifiers[RightShift] = true
		k.CurrentMode.DoKeyDown(RightShift)
	} else if k.activeModifiers[RightShift] {
		k.activeModifiers[RightShift] = false
		k.CurrentMode.DoKeyUp(RightShift)
	}

	if keys&64 != 0 {
		k.activeModifiers[RightAlt] = true
		k.CurrentMode.DoKeyDown(RightAlt)
	} else if k.activeModifiers[RightAlt] {
		k.activeModifiers[RightAlt] = false
		k.CurrentMode.DoKeyUp(RightAlt)
	}

	if keys&128 != 0 {
		k.activeModifiers[RightMeta] = true
		k.CurrentMode.DoKeyDown(RightMeta)
	} else if k.activeModifiers[RightMeta] {
		k.activeModifiers[RightMeta] = false
		k.CurrentMode.DoKeyUp(RightMeta)
	}

}
