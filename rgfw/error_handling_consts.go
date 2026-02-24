package rgfw

/*
	#include "RGFW.h"
*/
import "C"

type DebugType uint8

const (
	TypeError   DebugType = C.RGFW_typeError
	TypeWarning DebugType = C.RGFW_typeWarning
	TypeInfo    DebugType = C.RGFW_typeInfo
)

type ErrorCode uint8

const (
	ErrorNone           ErrorCode = C.RGFW_noError
	ErrorOutOfMemory    ErrorCode = C.RGFW_errOutOfMemory
	ErrorOpenGLContext  ErrorCode = C.RGFW_errOpenGLContext
	ErrorEGLContext     ErrorCode = C.RGFW_errEGLContext
	ErrorWayland        ErrorCode = C.RGFW_errWayland
	ErrorX11            ErrorCode = C.RGFW_errX11
	ErrorDirectXContext ErrorCode = C.RGFW_errDirectXContext
	ErrorIOKit          ErrorCode = C.RGFW_errIOKit
	ErrorClipboard      ErrorCode = C.RGFW_errClipboard
	ErrorFailedFuncLoad ErrorCode = C.RGFW_errFailedFuncLoad
	ErrorBuffer         ErrorCode = C.RGFW_errBuffer
	// ErrorMetal            ErrorCode = C.RGFW_errMetal
	// ErrorPlatform         ErrorCode = C.RGFW_errPlatform
	ErrorEventQueue ErrorCode = C.RGFW_errEventQueue
	InfoWindow      ErrorCode = C.RGFW_infoWindow
	InfoBuffer      ErrorCode = C.RGFW_infoBuffer
	InfoGlobal      ErrorCode = C.RGFW_infoGlobal
	InfoOpenGL      ErrorCode = C.RGFW_infoOpenGL
	WarningWayland  ErrorCode = C.RGFW_warningWayland
	WarningOpenGL   ErrorCode = C.RGFW_warningOpenGL
)
