#include "_cgo_export.h"
#include "RGFW.h"

void RFGW_setDebugCallbackCB() {
    RGFW_setDebugCallback((RGFW_debugfunc) goDebugCB);
}