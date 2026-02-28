package rgfw

/*
	#include "RGFW.h"
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

type MouseIcons uint8

const (
	MouseIconNormal MouseIcons = C.RGFW_mouseNormal
	MouseIconArrow  MouseIcons = C.RGFW_mouseArrow
	MouseIconIbeam  MouseIcons = C.RGFW_mouseIbeam
	// MouseIconText         MouseIcons = C.RGFW_mouseText
	MouseIconCrosshair    MouseIcons = C.RGFW_mouseCrosshair
	MouseIconPointingHand MouseIcons = C.RGFW_mousePointingHand
	MouseIconResizeEW     MouseIcons = C.RGFW_mouseResizeEW
	MouseIconResizeNS     MouseIcons = C.RGFW_mouseResizeNS
	MouseIconResizeNWSE   MouseIcons = C.RGFW_mouseResizeNWSE
	MouseIconResizeNESW   MouseIcons = C.RGFW_mouseResizeNESW
	// MouseIconResizeNW     MouseIcons = C.RGFW_mouseResizeNW
	// MouseIconResizeN      MouseIcons = C.RGFW_mouseResizeN
	// MouseIconResizeNE     MouseIcons = C.RGFW_mouseResizeNE
	// MouseIconResizeE      MouseIcons = C.RGFW_mouseResizeE
	// MouseIconResizeSE     MouseIcons = C.RGFW_mouseResizeSE
	// MouseIconResizeS      MouseIcons = C.RGFW_mouseResizeS
	// MouseIconResizeSW     MouseIcons = C.RGFW_mouseResizeSW
	// MouseIconResizeW      MouseIcons = C.RGFW_mouseResizeW
	MouseIconResizeAll  MouseIcons = C.RGFW_mouseResizeAll
	MouseIconNotAllowed MouseIcons = C.RGFW_mouseNotAllowed
	// MouseIconWait         MouseIcons = C.RGFW_mouseWait
	// MouseIconProgress     MouseIcons = C.RGFW_mouseProgress
	MouseIconCount MouseIcons = C.RGFW_mouseIconCount
)
