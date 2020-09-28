package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// CustomEntry is extended entry widget
type CustomEntry struct {
	widget.Entry
	onEnter func()
}

// KeyDown extended KeyDown event for CustomEntry
func (e *CustomEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyReturn:
		e.onEnter()
	default:
		e.Entry.KeyDown(key)
	}

}

// NewCustomEntry creates a new customized entry object with keypress and other extra functions
func NewCustomEntry(onEnterPressed func()) *CustomEntry {
	entry := &CustomEntry{}
	entry.ExtendBaseWidget(entry)
	entry.onEnter = onEnterPressed
	return entry
}
