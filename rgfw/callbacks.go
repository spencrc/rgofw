package rgfw

/*
	#include "RFGW_impl.h"
	void RFGW_setWindowMovedCallbackCB();
*/
import "C"

//export goWindowMovedCB
func goWindowMovedCB(win *C.RGFW_window, x, y C.i32) {
	goWin := windows.get(win)
	if goWin.fWindowMoved != nil {
		goWin.fWindowMoved(goWin, int32(x), int32(y))
	}
}

//export goWindowResizedCB
func goWindowResizedCB(win *C.RGFW_window, w, h C.i32) {
	goWin := windows.get(win)
	if goWin.fWindowMoved != nil {
		goWin.fWindowResized(goWin, int32(w), int32(h))
	}
}

//export goWindowRestoredCB
func goWindowRestoredCB(win *C.RGFW_window, x, y, w, h C.i32) {
	goWin := windows.get(win)
	goWin.fWindowRestored(goWin, int32(x), int32(y), int32(w), int32(h))
}

type WindowMaximizedCallback func(win *C.RGFW_window, x, y, w, h C.i32)

type WindowMinimizedCallback func(win *C.RGFW_window)

type WindowQuitCallback func(win *C.RGFW_window)

type FocusCallback func(win *C.RGFW_window, inFocus C.RGFW_bool)

type MouseNotifyCallback func(win *C.RGFW_window, x, y C.i32, status C.RGFW_bool)

type MousePosCallback func(win *C.RGFW_window, x, y C.i32, vecX, vecY C.float)

type DataDragCallback func(win *C.RGFW_window, x, y C.i32)

type WindowRefreshCallback func(win *C.RGFW_window)

type KeyCallback func(win *C.RGFW_window, key, sym byte, mod C.RGFW_keymod, repeat, pressed C.RGFW_bool)

type MouseButtonCallback func(win *C.RGFW_window, button C.RGFW_mouseButton, pressed C.RGFW_bool)

type MouseScrollCallback func(win *C.RGFW_window, x, y C.float)

type DataDropCallback func(win *C.RGFW_window, files *C.char)

type ScaleUpdatedCallback func(win *C.RGFW_window, scaleX, scaleY C.float)

type WindowMovedCallback func(win *Window, x int32, y int32)

func (win *Window) SetWindowMovedCallback(callback WindowMovedCallback) (previous WindowMovedCallback) {
	previous = win.fWindowMoved
	win.fWindowMoved = callback
	C.RFGW_setWindowMovedCallbackCB()
	return previous
}
