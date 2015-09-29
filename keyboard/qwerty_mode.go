package keyboard

type QwertyMode struct {
}

func (qm QwertyMode) DoKeyUp(k Key) {
	if ki, ok := LookupKey(k); ok {
		RawDoKeyUp(ki)
	}

}

func (qm QwertyMode) DoKeyDown(k Key) {
	if ki, ok := LookupKey(k); ok {
		RawDoKeyDown(ki)
	}
}

func (qm QwertyMode) MacroKeyDown(k int) {
}
func (qm QwertyMode) MacroKeyUp(k int) {
}
