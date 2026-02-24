package rgfw

/*
	#include "RFGW_impl.h"
*/
import "C"

type Format uint8

const (
	FormatRGB8  Format = C.RGFW_formatRGB8  /*!< 8-bit RGB (3 channels) */
	FormatBGR8  Format = C.RGFW_formatBGR8  /*!< 8-bit BGR (3 channels) */
	FormatRGBA8 Format = C.RGFW_formatRGBA8 /*!< 8-bit RGBA (4 channels) */
	FormatARGB8 Format = C.RGFW_formatARGB8 /*!< 8-bit RGBA (4 channels) */
	FormatBGRA8 Format = C.RGFW_formatBGRA8 /*!< 8-bit BGRA (4 channels) */
	FormatABGR8 Format = C.RGFW_formatABGR8 /*!< 8-bit BGRA (4 channels) */
)

type Icon uint8

const (
	IconTaskbar Icon = C.RGFW_iconTaskbar
	IconWindow  Icon = C.RGFW_iconWindow
	IconBoth    Icon = C.RGFW_iconBoth
)

type MouseIcon uint8

const (
	MouseIconNormal MouseIcon = C.RGFW_mouseNormal
	MouseIconArrow  MouseIcon = C.RGFW_mouseArrow
	MouseIconIbeam  MouseIcon = C.RGFW_mouseIbeam
	// MouseIconText         MouseIcon = C.RGFW_mouseText
	MouseIconCrosshair    MouseIcon = C.RGFW_mouseCrosshair
	MouseIconPointingHand MouseIcon = C.RGFW_mousePointingHand
	MouseIconResizeEW     MouseIcon = C.RGFW_mouseResizeEW
	MouseIconResizeNS     MouseIcon = C.RGFW_mouseResizeNS
	MouseIconResizeNWSE   MouseIcon = C.RGFW_mouseResizeNWSE
	MouseIconResizeNESW   MouseIcon = C.RGFW_mouseResizeNESW
	// MouseIconResizeNW     MouseIcon = C.RGFW_mouseResizeNW
	// MouseIconResizeN      MouseIcon = C.RGFW_mouseResizeN
	// MouseIconResizeNE     MouseIcon = C.RGFW_mouseResizeNE
	// MouseIconResizeE      MouseIcon = C.RGFW_mouseResizeE
	// MouseIconResizeSE     MouseIcon = C.RGFW_mouseResizeSE
	// MouseIconResizeS      MouseIcon = C.RGFW_mouseResizeS
	// MouseIconResizeSW     MouseIcon = C.RGFW_mouseResizeSW
	// MouseIconResizeW      MouseIcon = C.RGFW_mouseResizeW
	MouseIconResizeAll  MouseIcon = C.RGFW_mouseResizeAll
	MouseIconNotAllowed MouseIcon = C.RGFW_mouseNotAllowed
	// MouseIconWait         MouseIcon = C.RGFW_mouseWait
	// MouseIconProgress     MouseIcon = C.RGFW_mouseProgress
	MouseIconCount MouseIcon = C.RGFW_mouseIconCount
)
