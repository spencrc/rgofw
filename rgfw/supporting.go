package rgfw

/*
	#include "RGFW.h"
*/
import "C"

// Sets the root (main) RGFW window.
func SetRootWindow(win *Window) {
	C.RGFW_setRootWindow(win.ref)
}

func GetRootWindow() *Window {
	ref := C.RGFW_getRootWindow()
	return windows.get(ref)
}

// Pushes an event into the standard RGFW event queue.
// RGFW_eventQueuePush

func EventQueueFlush() {
	C.RGFW_eventQueueFlush()
}

func (win *Window) EventQueuePop() Event {
	cEvent := C.RGFW_eventQueuePop(win.ref)
	return getEventData(cEvent, win)
}

// Converts an API keycode to the RGFW unmapped (physical) key.
// RGFW_apiKeyToRGFW

// Converts an RGFW keycode to the unmapped (physical) API key.
// RGFW_rgfwToApiKey

// Converts an RGFW keycode to the mapped character representation.
// RGFW_rgfwToKeyChar

// Retrieves the size of the RGFW_info structure.
// RGFW_sizeofInfo

// Initializes the RGFW library.
// RGFW_init

// Deinitializes the RGFW library.
// RGFW_deinit

// Initializes RGFW using a user-provided RGFW_info structure.
// RGFW_init_ptr

// Deinitializes a specific RGFW instance stored in the provided RGFW_info pointer.
// RGFW_deinit_ptr

// Sets the global RGFW_info structure pointer.
// RGFW_setInfo

// Retrieves the global RGFW_info structure pointer.
// RGFW_getInfo
