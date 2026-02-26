package rgfw

/*
	#include "RGFW.h"
*/
import "C"
import "unsafe"

// Reads clipboard data. Returns clipboard data as string
func ReadClipboard() string {
	var cSize C.size_t
	cStr := C.RGFW_readClipboard(&cSize)
	return C.GoStringN(cStr, C.int(cSize))
}

// INTENTIONALLY UNIMPLEMENTED. WILL NOT ADD
// Reads clipboard data into a provided buffer, or returns the required length if str is NULL.
// RGFW_readClipboardPtr

// Writes text to the clipboard.
func WriteClipboard(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	C.RGFW_writeClipboard(cText, C.u32(len(text)))
}
