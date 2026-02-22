package rgfw

/*
	// This header includes RGFW and RGFW_IMPLEMENTATION! Including RGFW again will freak Go out
	#include "RGFW_impl.h"
*/
import "C"
import "unsafe"

// Retrieves the current X11 display connection.
// Returns a pointer to the X11 display, or nil if the platform is not in use.
func GetDisplayX11() unsafe.Pointer {
	return unsafe.Pointer(C.RGFW_getDisplay_X11())
}