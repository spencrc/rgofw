package rgfw

type Key uint8
type Keymod uint8
type MouseButton uint8

const (
	KeyNull      Key = 0
	Escape       Key = 27
	Backtick     Key = 96
	Key0         Key = 48
	Key1         Key = 49
	Key2         Key = 50
	Key3         Key = 51
	Key4         Key = 52
	Key5         Key = 53
	Key6         Key = 54
	Key7         Key = 55
	Key8         Key = 56
	Key9         Key = 57
	Minus        Key = 45
	Equal        Key = 61
	Equals       Key = Equal
	BackSpace    Key = 8
	Tab          Key = 9
	Space        Key = 32
	KeyA         Key = 97
	KeyB         Key = 98
	KeyC         Key = 99
	KeyD         Key = 100
	KeyE         Key = 101
	KeyF         Key = 102
	KeyG         Key = 103
	KeyH         Key = 104
	KeyI         Key = 105
	KeyJ         Key = 106
	KeyK         Key = 107
	KeyL         Key = 108
	KeyM         Key = 109
	KeyN         Key = 110
	KeyO         Key = 111
	KeyP         Key = 112
	KeyQ         Key = 113
	KeyR         Key = 114
	KeyS         Key = 115
	KeyT         Key = 116
	KeyU         Key = 117
	KeyV         Key = 118
	KeyW         Key = 119
	KeyX         Key = 120
	KeyY         Key = 121
	KeyZ         Key = 122
	Period       Key = 46
	Comma        Key = 44
	Slash        Key = 47
	Bracket      Key = 91
	CloseBracket Key = 93
	Semicolon    Key = 59
	Apostrophe   Key = 39
	BackSlash    Key = 92
	Return       Key = 10
	Enter        Key = Return
	Delete       Key = 127
	F1			 Key = 128 + iota
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	F16
	F17
	F18
	F19
	F20
	F21
	F22
	F23
	F24
	F25
	CapsLock
	ShiftL
	ControlL
	AltL
	SuperL
	ShiftR
	ControlR
	AltR
	SuperR
	Up
	Down
	Left
	Right
	Insert
	Menu
	End
	Home
	PageUp
	PageDown
	NumLock
	KpSlash
	KpMultiply
	KpPlus
	KpMinus
	KpEqual
	KpEquals = KpEqual
	Kp1
	Kp2
	Kp3
	Kp4
	Kp5
	Kp6
	Kp7
	Kp8
	Kp9
	Kp0
	KpPeriod
	KpReturn
	ScrollLock
	PrintScreen
	Pause
	World1
	World2
	Last 		 Key = 255 // uint8 max is 255
)

const (
	ModCapsLock   Keymod = 1 << 0 // 1
	ModNumLock    Keymod = 1 << 1 // 2
	ModControl    Keymod = 1 << 2 // 4
	ModAlt        Keymod = 1 << 3 // 8
	ModShift      Keymod = 1 << 4 // 16
	ModSuper      Keymod = 1 << 5 // 32
	ModScrollLock Keymod = 1 << 6 // 64
)

const (
	MouseLeft MouseButton = iota
	MouseMiddle
	MouseRight
	MouseMisc1
	MouseMisc2
	MouseMisc3
	MouseMisc4
	MouseMisc5
	MouseFinal
)