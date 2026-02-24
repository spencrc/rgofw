package rgfw

/*
	#include "RGFW.h"
*/
import "C"
import (
	"sync"
)

// Internal window logic. Taken from https://github.com/go-gl/glfw/blob/master/v3.3/glfw/window.go
type windowMap struct {
	m     sync.Mutex
	cToGo map[*C.RGFW_window]*Window
}

var windows = windowMap{cToGo: map[*C.RGFW_window]*Window{}}

func (wl *windowMap) put(win *Window) {
	wl.m.Lock()
	defer wl.m.Unlock()
	wl.cToGo[win.ref] = win
}

func (wl *windowMap) remove(ref *C.RGFW_window) {
	wl.m.Lock()
	defer wl.m.Unlock()
	delete(wl.cToGo, ref)
}

func (wl *windowMap) get(ref *C.RGFW_window) *Window {
	wl.m.Lock()
	defer wl.m.Unlock()
	return wl.cToGo[ref]
}
