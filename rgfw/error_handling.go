package rgfw

/*
	#include "RGFW.h"
	void RFGW_setDebugCallbackCB();
*/
import "C"
import "unsafe"

type DebugCallback func(t DebugType, err ErrorCode, msg string)

var fDebug DebugCallback

//export goDebugCB
func goDebugCB(t C.RGFW_debugType , err C.RGFW_errorCode, msg *C.char) {
	if fDebug != nil {
		fDebug(DebugType(t), ErrorCode(err), C.GoString(msg))
	}
}

// Sets the callback function to handle debug messages from RGFW.
func SetDebugCallback(callback DebugCallback) (previous DebugCallback) {
	previous = fDebug
	fDebug = callback
	C.RFGW_setDebugCallbackCB()
	return previous
}

// Sends a debug message manually through the currently set debug callback.
func SendDebugInfo(t DebugType, err ErrorCode, msg string) {
	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))
	C.RGFW_sendDebugInfo(C.RGFW_debugType(t), C.RGFW_errorCode(err), cMsg)
}
