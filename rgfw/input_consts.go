package rgfw

/*
	#include "RFGW_impl.h"
*/
import "C"

type WindowFlag uint

const (
	WindowFlagNoBorder    WindowFlag = C.RGFW_windowNoBorder
	WindowFlagNoResize    WindowFlag = C.RGFW_windowNoResize
	WindowFlagAllowDND    WindowFlag = C.RGFW_windowAllowDND
	WindowFlagHideMouse   WindowFlag = C.RGFW_windowHideMouse
	WindowFlagFullscreen  WindowFlag = C.RGFW_windowFullscreen
	WindowFlagTransparent WindowFlag = C.RGFW_windowTransparent
	WindowFlagCenter      WindowFlag = C.RGFW_windowCenter
	// WindowFlagRawMouse          WindowFlag = C.RGFW_windowRawMouse
	WindowFlagScaleToMonitor WindowFlag = C.RGFW_windowScaleToMonitor
	WindowFlagHide           WindowFlag = C.RGFW_windowHide
	WindowFlagMaximize       WindowFlag = C.RGFW_windowMaximize
	WindowFlagCenterCursor   WindowFlag = C.RGFW_windowCenterCursor
	WindowFlagFloating       WindowFlag = C.RGFW_windowFloating
	WindowFlagFocusOnShow    WindowFlag = C.RGFW_windowFocusOnShow
	WindowFlagMinimize       WindowFlag = C.RGFW_windowMinimize
	WindowFlagFocus          WindowFlag = C.RGFW_windowFocus
	// WindowFlagCaptureMouse      WindowFlag = C.RGFW_windowCaptureMouse
	WindowFlagOpenGL WindowFlag = C.RGFW_windowOpenGL
	WindowFlagEGL    WindowFlag = C.RGFW_windowEGL
	// WindowFlagNoDeinitOnClose   WindowFlag = C.RGFW_noDeinitOnClose
	WindowFlagWindowedFullscreen WindowFlag = C.RGFW_windowedFullscreen
	// WindowFlagCaptureRawMouse   WindowFlag = C.RGFW_windowCaptureRawMouse
)

// NOT AVAILABLE IN CURRENT RGFW VERSION (1.8.5) BUT AVAILABLE IN 2.0.0
// type FlashRequest uint8

// const (
// 	FlashCancel FlashRequest = C.RGFW_flashCancel
// 	FlashBriefly FlashRequest = C.RGFW_flashBriefly
// 	FlashUntilFocused FlashRequest = C.RGFW_flashUntilFocused
// )

type Key uint8

const (
	KeyNULL     Key = C.RGFW_keyNULL
	KeyEscape   Key = C.RGFW_escape
	KeyBacktick Key = C.RGFW_backtick
	Key0        Key = C.RGFW_0
	Key1        Key = C.RGFW_1
	Key2        Key = C.RGFW_2
	Key3        Key = C.RGFW_3
	Key4        Key = C.RGFW_4
	Key5        Key = C.RGFW_5
	Key6        Key = C.RGFW_6
	Key7        Key = C.RGFW_7
	Key8        Key = C.RGFW_8
	Key9        Key = C.RGFW_9
	KeyMinus    Key = C.RGFW_minus
	// KeyEqual      Key = C.RGFW_equals
	KeyEquals       Key = C.RGFW_equals
	KeyBackSpace    Key = C.RGFW_backSpace
	KeyTab          Key = C.RGFW_tab
	KeySpace        Key = C.RGFW_space
	KeyA            Key = C.RGFW_a
	KeyB            Key = C.RGFW_b
	KeyC            Key = C.RGFW_c
	KeyD            Key = C.RGFW_d
	KeyE            Key = C.RGFW_e
	KeyF            Key = C.RGFW_f
	KeyG            Key = C.RGFW_g
	KeyH            Key = C.RGFW_h
	KeyI            Key = C.RGFW_i
	KeyJ            Key = C.RGFW_j
	KeyK            Key = C.RGFW_k
	KeyL            Key = C.RGFW_l
	KeyM            Key = C.RGFW_m
	KeyN            Key = C.RGFW_n
	KeyO            Key = C.RGFW_o
	KeyP            Key = C.RGFW_p
	KeyQ            Key = C.RGFW_q
	KeyR            Key = C.RGFW_r
	KeyS            Key = C.RGFW_s
	KeyT            Key = C.RGFW_t
	KeyU            Key = C.RGFW_u
	KeyV            Key = C.RGFW_v
	KeyW            Key = C.RGFW_w
	KeyX            Key = C.RGFW_x
	KeyY            Key = C.RGFW_y
	KeyZ            Key = C.RGFW_z
	KeyPeriod       Key = C.RGFW_period
	KeyComma        Key = C.RGFW_comma
	KeySlash        Key = C.RGFW_slash
	KeyBracket      Key = C.RGFW_bracket
	KeyCloseBracket Key = C.RGFW_closeBracket
	KeySemicolon    Key = C.RGFW_semicolon
	KeyApostrophe   Key = C.RGFW_apostrophe
	KeyBackSlash    Key = C.RGFW_backSlash
	KeyReturn       Key = C.RGFW_return
	KeyEnter        Key = C.RGFW_enter
	KeyDelete       Key = C.RGFW_delete
	KeyF1           Key = C.RGFW_F1
	KeyF2           Key = C.RGFW_F2
	KeyF3           Key = C.RGFW_F3
	KeyF4           Key = C.RGFW_F4
	KeyF5           Key = C.RGFW_F5
	KeyF6           Key = C.RGFW_F6
	KeyF7           Key = C.RGFW_F7
	KeyF8           Key = C.RGFW_F8
	KeyF9           Key = C.RGFW_F9
	KeyF10          Key = C.RGFW_F10
	KeyF11          Key = C.RGFW_F11
	KeyF12          Key = C.RGFW_F12
	KeyF13          Key = C.RGFW_F13
	KeyF14          Key = C.RGFW_F14
	KeyF15          Key = C.RGFW_F15
	KeyF16          Key = C.RGFW_F16
	KeyF17          Key = C.RGFW_F17
	KeyF18          Key = C.RGFW_F18
	KeyF19          Key = C.RGFW_F19
	KeyF20          Key = C.RGFW_F20
	KeyF21          Key = C.RGFW_F21
	KeyF22          Key = C.RGFW_F22
	KeyF23          Key = C.RGFW_F23
	KeyF24          Key = C.RGFW_F24
	KeyF25          Key = C.RGFW_F25
	KeyCapsLock     Key = C.RGFW_capsLock
	KeyShiftL       Key = C.RGFW_shiftL
	KeyControlL     Key = C.RGFW_controlL
	KeyAltL         Key = C.RGFW_altL
	KeySuperL       Key = C.RGFW_superL
	KeyShiftR       Key = C.RGFW_shiftR
	KeyControlR     Key = C.RGFW_controlR
	KeyAltR         Key = C.RGFW_altR
	KeySuperR       Key = C.RGFW_superR
	KeyUp           Key = C.RGFW_up
	KeyDown         Key = C.RGFW_down
	KeyLeft         Key = C.RGFW_left
	KeyRight        Key = C.RGFW_right
	KeyInsert       Key = C.RGFW_insert
	KeyMenu         Key = C.RGFW_menu
	KeyEnd          Key = C.RGFW_end
	KeyHome         Key = C.RGFW_home
	KeyPageUp       Key = C.RGFW_pageUp
	KeyPageDown     Key = C.RGFW_pageDown
	KeyNumLock      Key = C.RGFW_numLock
	KeyKpSlash      Key = C.RGFW_kpSlash
	KeyKpMultiply   Key = C.RGFW_kpMultiply
	KeyKpPlus       Key = C.RGFW_kpPlus
	KeyKpMinus      Key = C.RGFW_kpMinus
	KeyKpEqual      Key = C.RGFW_kpEqual
	// KeyKpEquals   Key = C.RGFW_kpEqual
	KeyKp1         Key = C.RGFW_kp1
	KeyKp2         Key = C.RGFW_kp2
	KeyKp3         Key = C.RGFW_kp3
	KeyKp4         Key = C.RGFW_kp4
	KeyKp5         Key = C.RGFW_kp5
	KeyKp6         Key = C.RGFW_kp6
	KeyKp7         Key = C.RGFW_kp7
	KeyKp8         Key = C.RGFW_kp8
	KeyKp9         Key = C.RGFW_kp9
	KeyKp0         Key = C.RGFW_kp0
	KeyKpPeriod    Key = C.RGFW_kpPeriod
	KeyKpReturn    Key = C.RGFW_kpReturn
	KeyScrollLock  Key = C.RGFW_scrollLock
	KeyPrintScreen Key = C.RGFW_printScreen
	KeyPause       Key = C.RGFW_pause
	KeyWorld1      Key = C.RGFW_world1
	KeyWorld2      Key = C.RGFW_world2
)

type MouseButton uint8

const (
	MouseLeft   MouseButton = C.RGFW_mouseLeft
	MouseMiddle MouseButton = C.RGFW_mouseMiddle
	MouseRight  MouseButton = C.RGFW_mouseRight
	MouseMisc1  MouseButton = C.RGFW_mouseMisc1
	MouseMisc2  MouseButton = C.RGFW_mouseMisc2
	MouseMisc3  MouseButton = C.RGFW_mouseMisc3
	MouseMisc4  MouseButton = C.RGFW_mouseMisc4
	MouseMisc5  MouseButton = C.RGFW_mouseMisc5
)

type Keymod uint8

const (
	ModCapsLock   Keymod = C.RGFW_modCapsLock
	ModNumLock    Keymod = C.RGFW_modNumLock
	ModControl    Keymod = C.RGFW_modControl
	ModAlt        Keymod = C.RGFW_modAlt
	ModShift      Keymod = C.RGFW_modShift
	ModSuper      Keymod = C.RGFW_modSuper
	ModScrollLock Keymod = C.RGFW_modScrollLock
)
