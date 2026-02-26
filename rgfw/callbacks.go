package rgfw

/*
	#include "RFGW_impl.h"
	void RFGW_setWindowMovedCallbackCB();
	void RFGW_setWindowResizedCallbackCB();
	void RFGW_setWindowRestoredCallbackCB();
	void RFGW_setWindowMaximizedCallbackCB();
	void RFGW_setWindowMinimizedCallbackCB();
	void RFGW_setWindowQuitCallbackCB();
	void RFGW_setFocusCallbackCB();
	void RFGW_setMouseNotifyCallbackCB();
	void RFGW_setMousePosCallbackCB();
	void RFGW_setDataDragCallbackCB();
	void RFGW_setWindowRefreshCallbackCB();
	void RFGW_setKeyCallbackCB();
	void RFGW_setMouseButtonCallbackCB();
	void RFGW_setMouseScrollCallbackCB();
	void RFGW_setDataDropCallbackCB();
	void RFGW_setScaleUpdatedCallbackCB();
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
	if goWin.fWindowMoved != nil {
		goWin.fWindowRestored(goWin, int32(x), int32(y), int32(w), int32(h))
	}
}

//export goWindowMaximizedCB
func goWindowMaximizedCB(win *C.RGFW_window, x, y, w, h C.i32) {
	goWin := windows.get(win)
	if goWin.fWindowMaximized != nil {
		goWin.fWindowMaximized(goWin, int32(x), int32(y), int32(w), int32(h))
	}
}

//export goWindowMinimizedCB
func goWindowMinimizedCB(win *C.RGFW_window) {
	goWin := windows.get(win)
	if goWin.fWindowMinimized != nil {
		goWin.fWindowMinimized(goWin)
	}
}

//export goWindowQuitCB
func goWindowQuitCB(win *C.RGFW_window) {
	goWin := windows.get(win)
	if goWin.fWindowQuit != nil {
		goWin.fWindowQuit(goWin)
	}
}

//export goFocusCB
func goFocusCB(win *C.RGFW_window, inFocus C.RGFW_bool) {
	goWin := windows.get(win)
	if goWin.fFocus != nil {
		goWin.fFocus(goWin, inFocus != 0)
	}
}

//export goMouseNotifyCB
func goMouseNotifyCB(win *C.RGFW_window, x, y C.i32, status C.RGFW_bool) {
	goWin := windows.get(win)
	if goWin.fMouseNotify != nil {
		goWin.fMouseNotify(goWin, int32(x), int32(y), status != 0)
	}
}

//export goMousePosCB
func goMousePosCB(win *C.RGFW_window, x, y C.i32, vecX, vecY C.float) {
	goWin := windows.get(win)
	if goWin.fMousePos != nil {
		goWin.fMousePos(goWin, int32(x), int32(y), float32(vecX), float32(vecY))
	}
}

//export goDataDragCB
func goDataDragCB(win *C.RGFW_window, x, y C.i32) {
	goWin := windows.get(win)
	if goWin.fDataDrag != nil {
		goWin.fDataDrag(goWin, int32(x), int32(y))
	}
}

//export goWindowRefreshCB
func goWindowRefreshCB(win *C.RGFW_window) {
	goWin := windows.get(win)
	if goWin.fWindowRefresh != nil {
		goWin.fWindowRefresh(goWin)
	}
}

//export goKeyCB
func goKeyCB(win *C.RGFW_window, key, sym C.u8, mod C.RGFW_keymod, repeat, pressed C.RGFW_bool) {
	goWin := windows.get(win)
	if goWin.fKey != nil {
		goWin.fKey(goWin, byte(key), byte(sym), uint32(mod), repeat != 0, pressed != 0)
	}
}

//export goMouseButtonCB
func goMouseButtonCB(win *C.RGFW_window, button C.RGFW_mouseButton, pressed C.RGFW_bool) {
	goWin := windows.get(win)
	if goWin.fMouseButton != nil {
		goWin.fMouseButton(goWin, uint32(button), pressed != 0)
	}
}

//export goMouseScrollCB
func goMouseScrollCB(win *C.RGFW_window, x, y C.float) {
	goWin := windows.get(win)
	if goWin.fMouseScroll != nil {
		goWin.fMouseScroll(goWin, float32(x), float32(y))
	}
}

//export goDataDropCB
func goDataDropCB(win *C.RGFW_window, files **C.char, count C.size_t) {
	panic("unimplemented")
}

//export goScaleUpdatedCB
func goScaleUpdatedCB(win *C.RGFW_window, scaleX, scaleY C.float) {
	goWin := windows.get(win)
	if goWin.fScaleUpdated != nil {
		goWin.fScaleUpdated(goWin, float32(scaleX), float32(scaleY))
	}
}

type WindowMovedCallback func(win *Window, x, y int32)
type WindowResizedCallback func(win *Window, w, h int32)
type WindowRestoredCallback func(win *Window, x, y, w, h int32)
type WindowMaximizedCallback func(win *Window, x, y, w, h int32)
type WindowMinimizedCallback func(win *Window)
type WindowQuitCallback func(win *Window)
type FocusCallback func(win *Window, inFocus bool)
type MouseNotifyCallback func(win *Window, x, y int32, status bool)
type MousePosCallback func(win *Window, x, y int32, vecX, vecY float32)
type DataDragCallback func(win *Window, x, y int32)
type WindowRefreshCallback func(win *Window)
type KeyCallback func(win *Window, key, sym byte, mod uint32, repeat, pressed bool)
type MouseButtonCallback func(win *Window, button uint32, pressed bool)
type MouseScrollCallback func(win *Window, x, y float32)
type DataDropCallback func(win *Window, files []string)
type ScaleUpdatedCallback func(win *Window, scaleX, scaleY float32)

func (win *Window) SetWindowMovedCallback(callback WindowMovedCallback) (previous WindowMovedCallback) {
	previous = win.fWindowMoved
	win.fWindowMoved = callback
	C.RFGW_setWindowMovedCallbackCB()
	return previous
}

func (win *Window) SetWindowResizedCallback(callback WindowResizedCallback) (previous WindowResizedCallback) {
	previous = win.fWindowResized
	win.fWindowResized = callback
	C.RFGW_setWindowResizedCallbackCB()
	return previous
}

func (win *Window) SetWindowRestoredCallback(callback WindowRestoredCallback) (previous WindowRestoredCallback) {
	previous = win.fWindowRestored
	win.fWindowRestored = callback
	C.RFGW_setWindowRestoredCallbackCB()
	return previous
}

func (win *Window) SetWindowMaximizedCallback(callback WindowMaximizedCallback) (previous WindowMaximizedCallback) {
	previous = win.fWindowMaximized
	win.fWindowMaximized = callback
	C.RFGW_setWindowMaximizedCallbackCB()
	return previous
}

func (win *Window) SetWindowMinimizedCallback(callback WindowMinimizedCallback) (previous WindowMinimizedCallback) {
	previous = win.fWindowMinimized
	win.fWindowMinimized = callback
	C.RFGW_setWindowMinimizedCallbackCB()
	return previous
}

func (win *Window) SetWindowQuitCallback(callback WindowQuitCallback) (previous WindowQuitCallback) {
	previous = win.fWindowQuit
	win.fWindowQuit = callback
	C.RFGW_setWindowQuitCallbackCB()
	return previous
}

func (win *Window) SetFocusCallback(callback FocusCallback) (previous FocusCallback) {
	previous = win.fFocus
	win.fFocus = callback
	C.RFGW_setFocusCallbackCB()
	return previous
}

func (win *Window) SetMouseNotifyCallback(callback MouseNotifyCallback) (previous MouseNotifyCallback) {
	previous = win.fMouseNotify
	win.fMouseNotify = callback
	C.RFGW_setMouseNotifyCallbackCB()
	return previous
}

func (win *Window) SetMousePosCallback(callback MousePosCallback) (previous MousePosCallback) {
	previous = win.fMousePos
	win.fMousePos = callback
	C.RFGW_setMousePosCallbackCB()
	return previous
}

func (win *Window) SetDataDragCallback(callback DataDragCallback) (previous DataDragCallback) {
	previous = win.fDataDrag
	win.fDataDrag = callback
	C.RFGW_setDataDragCallbackCB()
	return previous
}

func (win *Window) SetWindowRefreshCallback(callback WindowRefreshCallback) (previous WindowRefreshCallback) {
	previous = win.fWindowRefresh
	win.fWindowRefresh = callback
	C.RFGW_setWindowRefreshCallbackCB()
	return previous
}

func (win *Window) SetKeyCallback(callback KeyCallback) (previous KeyCallback) {
	previous = win.fKey
	win.fKey = callback
	C.RFGW_setKeyCallbackCB()
	return previous
}

func (win *Window) SetMouseButtonCallback(callback MouseButtonCallback) (previous MouseButtonCallback) {
	previous = win.fMouseButton
	win.fMouseButton = callback
	C.RFGW_setMouseButtonCallbackCB()
	return previous
}

func (win *Window) SetMouseScrollCallback(callback MouseScrollCallback) (previous MouseScrollCallback) {
	previous = win.fMouseScroll
	win.fMouseScroll = callback
	C.RFGW_setMouseScrollCallbackCB()
	return previous
}

func (win *Window) SetDataDropCallback(callback DataDropCallback) (previous DataDropCallback) {
	previous = win.fDataDrop
	win.fDataDrop = callback
	C.RFGW_setDataDropCallbackCB()
	return previous
}

func (win *Window) SetScaleUpdatedCallback(callback ScaleUpdatedCallback) (previous ScaleUpdatedCallback) {
	previous = win.fScaleUpdated
	win.fScaleUpdated = callback
	C.RFGW_setScaleUpdatedCallbackCB()
	return previous
}