package rgfw

/*
	#include "RGFW.h"
*/
import "C"
import "unsafe"

type Surface struct {
	ref *C.RGFW_surface
}

// Allocates memory using the allocator defined by RGFW_ALLOC at compile time.
// RGFW_alloc

// Frees memory using the deallocator defined by RGFW_FREE at compile time.
// RGFW_free

// Returns the size (in bytes) of the RGFW_window structure.
// RGFW_sizeofWindow

// Returns the size (in bytes) of the RGFW_window_src structure.
// RGFW_sizeofWindowSrc

// (Unix) Toggles the use of Wayland, enabled by default when compiled with RGFW_WAYLAND.
// If not using RGFW_WAYLAND, Wayland functions are not exposed.
// This function can be used to force the use of XWayland.
// RGFW_useWayland

// Returns true if Wayland is currently being used
func UsingWayland() bool {
	return C.RGFW_usingWayland() == C.RGFW_TRUE
}

// Retrieves the current Cocoa layer (macOS only).
// Returns a pointer to the Cocoa layer, or nil if the platform is not in use
func GetLayerOSX() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_getLayer_OSX())
}

// Retrieves the current X11 display connection.
// Returns a pointer to the X11 display, or nil if the platform is not in use
func GetDisplayX11() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_getDisplay_X11())
}

// Retrieves the current Wayland display connection.
// Returns a pointer to the Wayland display, or nil if the platform is not in use
func GetDisplayWayland() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_getDisplay_Wayland())
}

// Sets the class name for X11 and WinAPI windows.
// Windows with the same class name will be grouped by the window manager.
// By default, the class name matches the root window's name.
// RGFW_setClassName

// Sets the X11 instance name.
// By default, the window name will be used as the instance name.
// RGFW_setXInstName

// (macOS only) Changes the current working directory to the application's resource folder.
// RGFW_moveToMacOSResourceDir

// Copy image to another image, respecting each image's format
// func CopyImageData(dest []uint8, w, h int32, destFmt Format, src []uint8, srcFmt Format) {
// 	if len(dest) == 0 {
// 		panic("rgfw: CopyImageData destination buffer nil or empty")
// 	}
// 	if len(src) == 0 {
// 		panic("rgfw: CopyImageData src buffer nil or empty")
// 	}
// 	C.RGFW_copyImageData(
// 		(*C.u8)(&dest[0]), C.i32(w), C.i32(h), C.RGFW_format(destFmt),
//         (*C.u8)(&src[0]), C.RGFW_format(srcFmt),
// 	)
// }

// Returns the size (in bytes) of the RGFW_nativeImage structure.
// RGFW_sizeofNativeImage

// Returns the size (in bytes) of the RGFW_surface structure.
// RGFW_sizeofSurface

// Creates a new surface from raw pixel data.
// NOTE: On X11, this uses the root window's visual, which may fail to render
// on any other window if the visual does not match. RGFW_window_createSurface
// and RGFW_window_createSurfacePtr exist only for X11 to address this issue.
// You can also manually set the root window with RGFW_setRootWindow.
// RGFW_createSurface

// Creates a new surface from raw pixel data.
//
// NOTE: when you create a surface using rgfw.CreateSurface, on X11 it uses the root window's visual
// this means it may fail to render on any other window if the visual does not match
// rgfw.CreateSurface exist only for X11 to address this issues
// Of course, you can also manually set the root window with rgfw.SetRootWindow
// func CreateSurface(data []uint8, w, h int32, format Format) *Surface {
// 	ref := C.RGFW_createSurface(
//         (*C.u8)(unsafe.Pointer(&data[0])),
//         C.i32(w),
//         C.i32(h),
//         C.RGFW_format(format),
//     )
// 	return &Surface{ref}
// }

// INTENTIONALLY UNIMPLEMENTED. WILL NOT ADD
// Creates a surface using a pre-allocated RGFW_surface structure.
// RGFW_createSurfacePtr

// Retrieves the native image associated with a surface.
// RGFW_nativeImage* RGFW_surface_getNativeImage(RGFW_surface* surface);

// Frees the surface pointer and any buffers used for software rendering.
// RGFW_surface_free

// Frees only the internal buffers used for software rendering, leaving the surface struct intact.
// RGFW_surface_freePtr

// Loads a mouse icon from bitmap data (similar to win.SetIcon)
//
// Note: The icon is not resized by default
// RGFW_loadMouse

// Frees the data associated with an RGFW_mouse structure
// RGFW_freeMouse

// Retrieves an array of all available monitors.
// Returns a pointer to an array of RGFW_monitor structures and the number of monitors found (max of 6)
// RGFW_getMonitors

// Retrieves the primary monitor.
func GetPrimaryMonitor() Monitor {
	m := C.RGFW_getPrimaryMonitor()
	return goMonitor(m)
}

// Requests a specific display mode for a monitor.
func (mon *Monitor) RequestMode(mode MonitorMode, request ModeRequest) bool {
	return C.RGFW_monitor_requestMode(mon.toC(), mode.toC(), C.RGFW_modeRequest(request)) == C.RGFW_TRUE
}

// Compares two monitor modes to check if they are equivalent.
// mon is the first monitor mode, and mon2 is the second monitor mode.
// request defines the comparison parameters.
// Returns true if both modes are equivalent
func MonitorModeCompare(mon, mon2 MonitorMode, request ModeRequest) bool {
	return C.RGFW_monitorModeCompare(mon.toC(), mon2.toC(), C.RGFW_modeRequest(request)) == C.RGFW_TRUE
}

// Scales a monitor’s mode to match a window’s size.
// Returns true if the scaling was successful
func (mon *Monitor) ScaleToWindow(win *Window) bool {
	return C.RGFW_monitor_scaleToWindow(mon.toC(), win.ref) == C.RGFW_TRUE
}

// sleep until RGFW gets an event or the timer ends (defined by OS)
func WaitForEvent(waitMS int32) {
	C.RGFW_waitForEvent(C.i32(waitMS))
}

// Set if events should be queued or not (enabled by default if the event queue is checked)
func SetQueueEvents(queue bool) {
	C.RGFW_setQueueEvents(C.RGFW_bool(boolToInt(queue)))
}

// Check all the events until there are none left and updates window structure attributes
func PollEvents() {
	C.RGFW_pollEvents()
}

// Check all the events until there are none left and updates window structure attributes
// queues events if the queue is checked and/or requested
func StopCheckEvents() {
	C.RGFW_stopCheckEvents()
}
