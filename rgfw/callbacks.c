#include "_cgo_export.h"
#include "RGFW.h"

void RFGW_setWindowMovedCallbackCB() {
    RGFW_setWindowMovedCallback((RGFW_windowMovedfunc) goWindowMovedCB);
}

void RFGW_setWindowResizedCallbackCB() {
    RGFW_setWindowResizedCallback((RGFW_windowResizedfunc) goWindowResizedCB);
}

void RFGW_setWindowRestoredCallbackCB() { 
    RGFW_setWindowRestoredCallback((RGFW_windowRestoredfunc)goWindowRestoredCB); 
}

void RFGW_setWindowMaximizedCallbackCB(){ 
    RGFW_setWindowMaximizedCallback((RGFW_windowMaximizedfunc)goWindowMaximizedCB); 
}

void RFGW_setWindowMinimizedCallbackCB(){ 
    RGFW_setWindowMinimizedCallback((RGFW_windowMinimizedfunc)goWindowMinimizedCB); 
}

void RFGW_setWindowQuitCallbackCB() { 
    RGFW_setWindowQuitCallback((RGFW_windowQuitfunc)goWindowQuitCB); 
}

void RFGW_setFocusCallbackCB() { 
    RGFW_setFocusCallback((RGFW_focusfunc)goFocusCB); 
}

void RFGW_setMouseNotifyCallbackCB() { 
    RGFW_setMouseNotifyCallback((RGFW_mouseNotifyfunc)goMouseNotifyCB); 
}

void RFGW_setMousePosCallbackCB() { 
    RGFW_setMousePosCallback((RGFW_mousePosfunc)goMousePosCB); 
}

void RFGW_setDataDragCallbackCB() { 
    RGFW_setDataDragCallback((RGFW_dataDragfunc)goDataDragCB); 
}

void RFGW_setWindowRefreshCallbackCB() { 
    RGFW_setWindowRefreshCallback((RGFW_windowRefreshfunc)goWindowRefreshCB); 
}

void RFGW_setKeyCallbackCB() { 
    RGFW_setKeyCallback((RGFW_keyfunc)goKeyCB); 

}

void RFGW_setMouseButtonCallbackCB() { 
    RGFW_setMouseButtonCallback((RGFW_mouseButtonfunc)goMouseButtonCB);
}

void
 RFGW_setMouseScrollCallbackCB() { 
    RGFW_setMouseScrollCallback((RGFW_mouseScrollfunc)goMouseScrollCB); 
}

void RFGW_setDataDropCallbackCB() { 
    RGFW_setDataDropCallback((RGFW_dataDropfunc)goDataDropCB); 
}

void RFGW_setScaleUpdatedCallbackCB() { 
    RGFW_setScaleUpdatedCallback((RGFW_scaleUpdatedfunc)goScaleUpdatedCB);
}