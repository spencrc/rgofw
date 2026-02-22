package rgfw

import "unsafe"

type EventType uint8

type CommonEvent struct {
	Type EventType
	Window *Window
}

type MouseButtonEvent struct {
	Type EventType
	Window *Window
	Value  MouseButton 
}

type MouseScrollEvent struct {
	Type EventType
	Window *Window
	X, Y   float32
}

type MousePosEvent struct {
	Type EventType
	Window     *Window
	X, Y       int32
	VecX, VecY float32
}

type KeyEvent struct {
	Type EventType
	Window *Window
	Value  Key
	Repeat bool   
	Mod    Keymod 
}

type KeyCharEvent struct {
	Type EventType
	Window *Window
	Value  uint32
}

type DataDropEvent struct {
	Type EventType
	Window *Window
	Files  []string // Converted from char** for Go-friendliness
	Count  uint64   // size_t
}

type DataDragEvent struct {
	Type EventType
	Window *Window
	X, Y   int32
}

type ScaleUpdatedEvent struct {
	Type EventType
	Window *Window
	X, Y   float32
}

type MonitorEvent struct {
	Type EventType
	Window  *Window
	Monitor unsafe.Pointer // Or a specific Monitor struct if you've wrapped it
}

const (
	EventNone             EventType = iota // 0
	KeyPressed                             // 1
	KeyReleased                            // 2
	KeyChar                                // 3
	MouseButtonPressed                     // 4
	MouseButtonReleased                    // 5
	MouseScroll                            // 6
	MousePosChanged                        // 7
	WindowMoved                            // 8
	WindowResized                          // 9
	FocusIn                                // 10
	FocusOut                               // 11
	MouseEnter                             // 12
	MouseLeave                             // 13
	WindowRefresh                          // 14
	Quit                                   // 15
	DataDrop                               // 16
	DataDrag                               // 17
	WindowMaximized                        // 18
	WindowMinimized                        // 19
	WindowRestored                         // 20
	ScaleUpdated                           // 21
	MonitorConnected                       // 22
	MonitorDisconnected                     // 23
)