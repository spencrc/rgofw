#pragma once
#include "../third_party/RGFW.h"

typedef struct {
    RGFW_eventType eventType;
    RGFW_mouseButton mouseButtonValue;
    float mouseScrollX, mouseScrollY;
    i32 mousePosX, mousePosY;
    float mousePosVecX, mousePosVecY;
    RGFW_key keyValue; 
	RGFW_bool keyRepeat;
	RGFW_keymod keyMod;
    // u32 keyCharValue;
    char** dataDropFiles;
    size_t dataDropCount;
    i32 dataDragX, dataDragY;
    float scaleUpdatedX, scaleUpdatedY;
    // const RGFW_monitor* monitor;
} wrapper_eventData;

wrapper_eventData wrapper_getEventData(RGFW_event* e);