package rgfw

/*
	#include "RFGW_impl.h"
*/
import "C"
import (
	"unsafe"
)

type Window struct {
	ref *C.RGFW_window

	fWindowMoved     WindowMovedCallback
	fWindowResized   WindowResizedCallback
	fWindowRestored  WindowRestoredCallback
	fWindowMaximized WindowMaximizedCallback
	fWindowMinimized WindowMinimizedCallback
	fWindowQuit      WindowQuitCallback
	fFocus           FocusCallback
	fMouseNotify     MouseNotifyCallback
	fMousePos        MousePosCallback
	fDataDrag        DataDragCallback
	fWindowRefresh   WindowRefreshCallback
	fKey             KeyCallback
	fMouseButton     MouseButtonCallback
	fMouseScroll     MouseScrollCallback
	fDataDrop        DataDropCallback
	fScaleUpdated    ScaleUpdatedCallback
}

// Returns true if the key is pressed during the current frame
// key is the key code of the key you want to check
func IsKeyPressed(key Key) bool {
	return C.RGFW_isKeyPressed(C.u8(key)) == C.RGFW_TRUE
}

// Returns true if the key was released during the current frame
// key is the key code of the key you want to check
func IsKeyReleased(key Key) bool {
	return C.RGFW_isKeyReleased(C.u8(key)) == C.RGFW_TRUE
}

// Returns true if the key is down.
// key is the key code of the key you want to check.
func IsKeyDown(key Key) bool {
	return C.RGFW_isKeyDown(C.u8(key)) == C.RGFW_TRUE
}

// Returns true if the mouse button is pressed during the current frame.
// button is the mouse button code of the button you want to check.
func IsMousePressed(button MouseButton) bool {
	return C.RGFW_isMousePressed(C.u8(button)) == C.RGFW_TRUE
}

// Returns true if the mouse button is released during the current frame.
// button is the mouse button code of the button you want to check.
func IsMouseReleased(button MouseButton) bool {
	return C.RGFW_isMouseReleased(C.u8(button)) == C.RGFW_TRUE
}

// Returns true if the mouse button is down.
// button is the mouse button code of the button you want to check.
func IsMouseDown(button MouseButton) bool {
	return C.RGFW_isMouseDown(C.u8(button)) == C.RGFW_TRUE
}

// Outputs the current x, y position of the mouse.
func GetMouseScroll() (float32, float32) {
	var x, y C.float
	C.RGFW_getMouseScroll(&x, &y)
	return float32(x), float32(y)
}

func GetMouseVector() (float32, float32) {
	var x, y C.float
	C.RGFW_getMouseVector(&x, &y)
	return float32(x), float32(y)
}

// Creates a new window.
func CreateWindow(name string, x, y, w, h int32, flags WindowFlag) *Window {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	ref := C.RGFW_createWindow(cName, C.i32(x), C.i32(y), C.i32(w), C.i32(h), C.uint(flags))

	win := &Window{ref: ref}
	windows.put(win)
	return win
}

// INTENTIONALLY UNIMPLEMENTED. WILL NOT ADD
// Creates a new window using a pre-allocated window structure.
// RGFW_createWindowPtr

// Creates a new surface structure
// Can throw an error! See RGFW_window_createSurfacePtr
// RGFW_window_createSurface

// INTENTIONALLY UNIMPLEMENTED. WILL NOT ADD
// Creates a new surface structure using a pre-allocated surface structure
// RGFW_window_createSurfacePtr

// Blits a surface stucture to the window
// RGFW_window_blitSurface

// Gets the position of the window
func (win *Window) GetPosition() (int32, int32) {
	var x, y C.i32
	C.RGFW_window_getPosition(win.ref, &x, &y)
	return int32(x), int32(y)
}

// Gets the size of the window
func (win *Window) GetSize() (int32, int32) {
	var w, h C.i32
	C.RGFW_window_getSize(win.ref, &w, &h)
	return int32(w), int32(h)
}

// NOT AVAILABLE IN CURRENT RGFW VERSION (1.8.5) BUT AVAILABLE IN 2.0.0
// Gets the size of the window in exact pixels
// func (win *Window) GetSizeInPixels() (int32, int32) {
// 	var w, h C.i32
// 	C.RGFW_window_getSizeInPixels(win.ref, &w, &h)
// 	return int32(w), int32(h)
// }

// Gets the flags of the window
func (win *Window) GetFlags() WindowFlag {
	return WindowFlag(C.RGFW_window_getFlags(win.ref))
}

// Returns the exit key assigned to the window
func (win *Window) GetExitKey() Key {
	return Key(C.RGFW_window_getExitKey(win.ref))
}

// Sets the exit key for the window
func (win *Window) SetExitKey(key Key) {
	C.RGFW_window_setExitKey(win.ref, C.RGFW_key(key))
}

// Sets the types of events you want the window to receive
func (win *Window) SetEnabledEvents(events EventFlag) {
	C.RGFW_window_setEnabledEvents(win.ref, C.RGFW_eventFlag(events))
}

// Gets the currently enabled events for the window
func (win *Window) GetEnabledEvents() EventFlag {
	return EventFlag(C.RGFW_window_getEnabledEvents(win.ref))
}

// Enables all events and disables selected ones
func (win *Window) SetDisabledEvents(events EventFlag) {
	C.RGFW_window_setDisabledEvents(win.ref, C.RGFW_eventFlag(events))
}

// Directly enables or disables a specific event or group of events
func (win *Window) SetEventState(event EventFlag, state bool) {
	C.RGFW_window_setEventState(win.ref, C.RGFW_eventFlag(event), C.RGFW_bool(boolToInt(state)))
}

// Gets the user pointer associated with the window
// func (win *Window) GetUserPtr() unsafe.Pointer {
//     return C.RGFW_window_getUserPtr(win.ref)
// }

// Sets a user pointer for the window
// func (win *Window) SetUserPtr(ptr unsafe.Pointer) {
//     C.RGFW_window_setUserPtr(win.ref, ptr)
// }

// INTENTIONALLY UNIMPLEMENTED. WILL NOT ADD
// Retrieves the platform-specific window source pointer
// RGFW_window_getSrc

// Sets the macOS layer object associated with the window.
// Returns a pointer to the macOS layer object
func (win *Window) SetLayerOSX(layer unsafe.Pointer) {
	C.RGFW_window_setLayer_OSX(win.ref, layer)
}

// Retrieves the macOS view object associated with the window.
// Returns a pointer to the macOS view object, or nil if not on macOS
func (win *Window) GetViewOSX() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_window_getView_OSX(win.ref))
}

// Retrieves the HWND handle for the window.
// Returns a pointer to the Windows HWND handle, or nil if not on Windows
func (win *Window) GetHWND() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_window_getHWND(win.ref))
}

// retrieves the HDC handle for the window.
// Returns a pointer to the Windows HDC handle, or nil if not on Windows
func (win *Window) GetHDC() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_window_getHDC(win.ref))
}

// Retrieves the X11 Window handle for the window.
// Returns the X11 Window handle, or 0 if not on X11
func (win *Window) GetWindowX11() uint32 {
	return uint32(C.RGFW_window_getWindow_X11(win.ref))
}

// Retrieves the Wayland surface handle for the window.
// Returns a pointer to the Wayland surface, or nil if not on Wayland
func (win *Window) GetWindowWayland() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_window_getWindow_Wayland(win.ref))
}
