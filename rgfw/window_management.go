package rgfw

/*
	#include "RGFW.h"
	#include "window_management.h"
*/
import "C"
import "unsafe"

type EventType uint8

const (
	EventNone        EventType = C.RGFW_eventNone
	EventKeyPressed  EventType = C.RGFW_keyPressed
	EventKeyReleased EventType = C.RGFW_keyReleased
	// EventKeyChar             EventType = C.RGFW_keyChar
	EventMouseButtonPressed  EventType = C.RGFW_mouseButtonPressed
	EventMouseButtonReleased EventType = C.RGFW_mouseButtonReleased
	EventMouseScroll         EventType = C.RGFW_mouseScroll
	EventMousePosChanged     EventType = C.RGFW_mousePosChanged
	EventWindowMoved         EventType = C.RGFW_windowMoved
	EventWindowResized       EventType = C.RGFW_windowResized
	EventFocusIn             EventType = C.RGFW_focusIn
	EventFocusOut            EventType = C.RGFW_focusOut
	EventMouseEnter          EventType = C.RGFW_mouseEnter
	EventMouseLeave          EventType = C.RGFW_mouseLeave
	EventWindowRefresh       EventType = C.RGFW_windowRefresh
	EventQuit                EventType = C.RGFW_quit
	EventDataDrop            EventType = C.RGFW_dataDrop
	EventDataDrag            EventType = C.RGFW_dataDrag
	EventWindowMaximized     EventType = C.RGFW_windowMaximized
	EventWindowMinimized     EventType = C.RGFW_windowMinimized
	EventWindowRestored      EventType = C.RGFW_windowRestored
	EventScaleUpdated        EventType = C.RGFW_scaleUpdated
	// EventMonitorConnected    EventType = C.RGFW_monitorConnected
	// EventMonitorDisconnected EventType = C.RGFW_monitorDisconnected
)

type EventFlag uint

const (
	EventFlagKeyPressed  EventFlag = C.RGFW_keyPressedFlag
	EventFlagKeyReleased EventFlag = C.RGFW_keyReleasedFlag
	// EventFlagKeyChar             EventFlag = C.RGFW_keyCharFlag
	EventFlagMouseScroll         EventFlag = C.RGFW_mouseScrollFlag
	EventFlagMouseButtonPressed  EventFlag = C.RGFW_mouseButtonPressedFlag
	EventFlagMouseButtonReleased EventFlag = C.RGFW_mouseButtonReleasedFlag
	EventFlagMousePosChanged     EventFlag = C.RGFW_mousePosChangedFlag
	EventFlagMouseEnter          EventFlag = C.RGFW_mouseEnterFlag
	EventFlagMouseLeave          EventFlag = C.RGFW_mouseLeaveFlag
	EventFlagWindowMoved         EventFlag = C.RGFW_windowMovedFlag
	EventFlagWindowResized       EventFlag = C.RGFW_windowResizedFlag
	EventFlagFocusIn             EventFlag = C.RGFW_focusInFlag
	EventFlagFocusOut            EventFlag = C.RGFW_focusOutFlag
	EventFlagWindowRefresh       EventFlag = C.RGFW_windowRefreshFlag
	EventFlagWindowMaximized     EventFlag = C.RGFW_windowMaximizedFlag
	EventFlagWindowMinimized     EventFlag = C.RGFW_windowMinimizedFlag
	EventFlagWindowRestored      EventFlag = C.RGFW_windowRestoredFlag
	EventFlagScaleUpdated        EventFlag = C.RGFW_scaleUpdatedFlag
	EventFlagQuit                EventFlag = C.RGFW_quitFlag
	EventFlagDataDrop            EventFlag = C.RGFW_dataDropFlag
	EventFlagDataDrag            EventFlag = C.RGFW_dataDragFlag
	// EventFlagMonitorConnected    EventFlag = C.RGFW_monitorConnectedFlag
	// EventFlagMonitorDisconnected EventFlag = C.RGFW_monitorDisconnectedFlag
	EventFlagKeyEvents      EventFlag = C.RGFW_keyEventsFlag
	EventFlagMouseEvents    EventFlag = C.RGFW_mouseEventsFlag
	EventFlagWindowEvents   EventFlag = C.RGFW_windowEventsFlag
	EventFlagFocusEvents    EventFlag = C.RGFW_focusEventsFlag
	EventFlagDataDropEvents EventFlag = C.RGFW_dataDropEventsFlag
	// EventFlagMonitorEvents       EventFlag = C.RGFW_monitorEventsFlag
	EventFlagAll EventFlag = C.RGFW_allEventFlags
)

type CommonEvent struct {
	Type   EventType
	Window *Window
}

type MouseButtonEvent struct {
	Value  MouseButton
}

type MouseScrollEvent struct {
	X, Y   float32
}

type MousePosEvent struct {
	X, Y       int32
	VecX, VecY float32
}

type KeyEvent struct {
	Value  Key
	Repeat bool
	Mod    Keymod
}

// type KeyCharEvent struct {
// 	Type EventType
// 	Window *Window
// 	Value  uint
// }

type DataDropEvent struct {
	Files  []string // Converted from char** for Go-friendliness
	Count  uint64   // size_t
}

type DataDragEvent struct {
	X, Y   int32
}

type ScaleUpdatedEvent struct {
	X, Y   float32
}

// type MonitorEvent struct {
// 	Type EventType
// 	Window  *Window
// 	Monitor unsafe.Pointer // Or a specific Monitor struct if you've wrapped it
// }

type Event struct {
	Type   EventType
	Common CommonEvent
	Button MouseButtonEvent
	Scroll MouseScrollEvent
	Mouse  MousePosEvent
	Key    KeyEvent
	// KeyChar KeyCharEvent
	Drop  DataDropEvent
	Drag  DataDragEvent
	Scale ScaleUpdatedEvent
	// Monitor MonitorEvent
}

func getEventData(cEvent *C.RGFW_event, win *Window) Event {
	data := C.wrapper_getEventData(cEvent)
	eventType := EventType(data.eventType)

	common := CommonEvent{
		Type:   eventType,
		Window: win,
	}

	goEvent := Event{Type: eventType, Common: common}

	switch eventType {
	case EventKeyPressed, EventKeyReleased:
		goEvent.Key = KeyEvent{
			Value:  Key(data.keyValue),
			Repeat: data.keyRepeat == C.RGFW_TRUE,
			Mod:    Keymod(data.keyMod),
		}
	case EventMouseButtonPressed, EventMouseButtonReleased:
		goEvent.Button = MouseButtonEvent{
			Value:  MouseButton(data.mouseButtonValue),
		}
	case EventMouseScroll:
		goEvent.Scroll = MouseScrollEvent{
			X: float32(data.mouseScrollX),
			Y: float32(data.mouseScrollY),
		}
	case EventMousePosChanged:
		goEvent.Mouse = MousePosEvent{
			X: int32(data.mousePosX),
			Y: int32(data.mousePosY),
			VecX: float32(data.mousePosVecX),
			VecY: float32(data.mousePosVecY),
		}
	case EventMouseEnter, EventMouseLeave:
		goEvent.Mouse = MousePosEvent{
			X: int32(data.mousePosX),
			Y: int32(data.mousePosY),
		}
	}

	return goEvent
}

// Set the window flags (will undo flags if they don't match the old ones)
func (win *Window) SetFlags(flags WindowFlag) {
	C.RGFW_window_setFlags(win.ref, C.RGFW_windowFlags(flags))
}

// Polls and pops the next event with the matching target window in event queue, pushes back events that don't match
//
// NOTE: Using this function without a loop may cause event lag.
// For multi-threaded systems, use RGFW_pollEvents combined with RGFW_window_checkQueuedEvent.
//
// Example:
//
//	 event, foundEvent := win.CheckEvent()
//	 for ;foundEvent; event, foundEvent = win.CheckEvent() {
//			// handle event
//	 }
func (win *Window) CheckEvent() (Event, bool) {
	var cEvent C.RGFW_event
	foundEvent := C.RGFW_window_checkEvent(win.ref, &cEvent)
	if foundEvent == C.RGFW_FALSE {
		return Event{}, false
	}

	goEvent := getEventData(&cEvent, win)

	return goEvent, true
}

// Pops the first queued event for the window
func (win *Window) CheckQueuedEvent() (Event, bool) {
	var cEvent C.RGFW_event
	foundEvent := C.RGFW_window_checkQueuedEvent(win.ref, &cEvent)
	if foundEvent == C.RGFW_FALSE {
		return Event{}, false
	}

	goEvent := getEventData(&cEvent, win)

	return goEvent, true
}

// Checks if a key was pressed while the window is in focus.
// key is the key code to check
func (win *Window) IsKeyPressed(key Key) bool {
	return C.RGFW_window_isKeyPressed(win.ref, C.RGFW_key(key)) == C.RGFW_TRUE
}

// Checks if a key is currently being held down.
// key is the key code to check
func (win *Window) IsKeyDown(key Key) bool {
	return C.RGFW_window_isKeyDown(win.ref, C.RGFW_key(key)) == C.RGFW_TRUE
}

// Checks if a key was released.
// key is the key code to check
func (win *Window) IsKeyReleased(key Key) bool {
	return C.RGFW_window_isKeyReleased(win.ref, C.RGFW_key(key)) == C.RGFW_TRUE
}

// Checks if a mouse button was pressed.
// button is the mouse button code to check
func (win *Window) IsMousePressed(button MouseButton) bool {
	return C.RGFW_window_isMousePressed(win.ref, C.RGFW_mouseButton(button)) == C.RGFW_TRUE
}

// Checks if a mouse button is currently held down.
// button is the mouse button code to check
func (win *Window) IsMouseDown(button MouseButton) bool {
	return C.RGFW_window_isMouseDown(win.ref, C.RGFW_mouseButton(button)) == C.RGFW_TRUE
}

// Checks if a mouse button was released.
// button is the mouse button code to check
func (win *Window) IsMouseReleased(button MouseButton) bool {
	return C.RGFW_window_isMouseReleased(win.ref, C.RGFW_mouseButton(button)) == C.RGFW_TRUE
}

// Checks if the mouse left the window (true only for the first frame)
func (win *Window) DidMouseLeave() bool {
	return C.RGFW_window_didMouseLeave(win.ref) == C.RGFW_TRUE
}

// Checks if the mouse entered the window (true only for the first frame)
func (win *Window) DidMouseEnter() bool {
	return C.RGFW_window_didMouseEnter(win.ref) == C.RGFW_TRUE
}

// Checks if the mouse is currently inside the window bounds
func (win *Window) IsMouseInside() bool {
	return C.RGFW_window_isMouseInside(win.ref) == C.RGFW_TRUE
}

// Checks if there is data being dragged into or within the window
func (win *Window) IsDataDragging() bool {
	return C.RGFW_window_isDataDragging(win.ref) == C.RGFW_TRUE
}

// Gets the position of a data drag
func (win *Window) GetDataDrag() (activeDrag bool, x, y int32) {
	var cX, cY C.i32
	res := C.RGFW_window_getDataDrag(win.ref, &cX, &cY)
	return res == C.RGFW_TRUE, int32(cX), int32(cY)
}

// Checks if a data drop occurred in the window (first frame only)
func (win *Window) DidDataDrop() bool {
	return C.RGFW_window_didDataDrop(win.ref) == C.RGFW_TRUE
}

// Retrieves files from a data drop (drag and drop)
// RGFW_window_getDataDrop

// Closes the window and frees its associated structure
func (win *Window) Close() {
	windows.remove(win.ref)
	C.RGFW_window_close(win.ref)
}

// INTENTIONALLY UNIMPLEMENTED. WILL NOT ADD
// Closes the window without freeing its structure
// RGFW_window_closePtr

// Moves the window to a new position on the screen
func (win *Window) Move(x, y int32) {
	C.RGFW_window_move(win.ref, C.i32(x), C.i32(y))
}

// Moves the window to a specific monitor
func (win *Window) MoveToMonitor(m Monitor) {
	C.RGFW_window_moveToMonitor(win.ref, m.toC())
}

// Resizes the window to the given dimensions
func (win *Window) Resize(w, h int32) {
	C.RGFW_window_resize(win.ref, C.i32(w), C.i32(h))
}

// Sets the aspect ratio of the window
func (win *Window) SetAspectRatio(w, h int32) {
	C.RGFW_window_setAspectRatio(win.ref, C.i32(w), C.i32(h))
}

// Sets the minimum size of the window
func (win *Window) SetMinSize(w, h int32) {
	C.RGFW_window_setMinSize(win.ref, C.i32(w), C.i32(h))
}

// Sets the maximum size of the window
func (win *Window) SetMaxSize(w, h int32) {
	C.RGFW_window_setMaxSize(win.ref, C.i32(w), C.i32(h))
}

// Sets focus to the window
func (win *Window) Focus() {
	C.RGFW_window_focus(win.ref)
}

// Checks if the window is currently in focus
func (win *Window) IsInFocus() bool {
	return C.RGFW_window_isInFocus(win.ref) == C.RGFW_TRUE
}

// Raises the window to the top of the stack
func (win *Window) Raise() {
	C.RGFW_window_raise(win.ref)
}

// Maximizes the window
func (win *Window) Maximize() {
	C.RGFW_window_maximize(win.ref)
}

// Toggles fullscreen mode for the window
func (win *Window) SetFullscreen(fullscreen bool) {
	C.RGFW_window_setFullscreen(win.ref, C.RGFW_bool(boolToInt(fullscreen)))
}

// Centers the window on the screen
func (win *Window) Center() {
	C.RGFW_window_center(win.ref)
}

// Minimizes the window
func (win *Window) Minimize() {
	C.RGFW_window_minimize(win.ref)
}

// Restores the window from minimized state
func (win *Window) Restore() {
	C.RGFW_window_restore(win.ref)
}

// Makes the window a floating window
func (win *Window) SetFloating(floating bool) {
	C.RGFW_window_setFloating(win.ref, C.RGFW_bool(boolToInt(floating)))
}

// Sets the opacity level of the window
func (win *Window) SetOpacity(opacity uint8) {
	C.RGFW_window_setOpacity(win.ref, C.u8(opacity))
}

// Toggles window borders
func (win *Window) SetBorder(border bool) {
	C.RGFW_window_setBorder(win.ref, C.RGFW_bool(boolToInt(border)))
}

// Checks if the window is borderless
func (win *Window) Borderless() bool {
	return C.RGFW_window_borderless(win.ref) == C.RGFW_TRUE
}

// Toggles drag-and-drop (DND) support for the window
func (win *Window) SetDND(allow bool) {
	C.RGFW_window_setDND(win.ref, C.RGFW_bool(boolToInt(allow)))
}

// Checks if drag-and-drop (DND) is allowed
func (win *Window) AllowsDND() bool {
	return C.RGFW_window_allowsDND(win.ref) == C.RGFW_TRUE
}

// Toggles mouse passthrough for the window
func (win *Window) SetMousePassthrough(passthrough bool) {
	C.RGFW_window_setMousePassthrough(win.ref, C.RGFW_bool(boolToInt(passthrough)))
}

// Renames the window
func (win *Window) SetName(name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.RGFW_window_setName(win.ref, cName)
}

// Sets the icon for the window and taskbar
func (win *Window) SetIcon(data []uint8, w, h int32, format Format) bool {
	return C.RGFW_window_setIcon(
		win.ref,
    	(*C.u8)(unsafe.Pointer((unsafe.SliceData(data)))),
        C.i32(w),
        C.i32(h),
        C.RGFW_format(format),
    ) == C.RGFW_TRUE
}

// Sets the icon for the window and/or taskbar
func (win *Window) SetIconEx(data []uint8, w, h int32, format Format, t Icon) bool {
	return C.RGFW_window_setIconEx(
		win.ref,
    	(*C.u8)(unsafe.Pointer((unsafe.SliceData(data)))),
        C.i32(w),
        C.i32(h),
        C.RGFW_format(format),
		C.RGFW_icon(t),
    ) == C.RGFW_TRUE
}

// Sets the mouse icon for the window using a loaded bitmap
func (win *Window) SetMouse(mouse *Mouse) {
	C.RGFW_window_setMouse(win.ref, mouse.ref)
}

// Sets the mouse to a standard system cursor.
func (win *Window) SetMouseStandard(mouse MouseIcons) {
	C.RGFW_window_setMouseStandard(win.ref, C.RGFW_mouseIcons(mouse))
}

// Sets the mouse to the default cursor icon.
func (win *Window) SetMouseDefault() bool {
	return C.RGFW_window_setMouseDefault(win.ref) == C.RGFW_TRUE
}

// Locks the cursor to the center of the window.
func (win *Window) HoldMouse() {
	C.RGFW_window_holdMouse(win.ref)
}

// Returns true if the mouse is currently held by RGFW.
func (win *Window) IsHoldingMouse() bool {
	return C.RGFW_window_isHoldingMouse(win.ref) == C.RGFW_TRUE
}

// Releases the mouse so it can move freely again.
func (win *Window) UnholdMouse() {
	C.RGFW_window_unholdMouse(win.ref)
}

// Hides the window from view.
func (win *Window) Hide() {
	C.RGFW_window_hide(win.ref)
}

// Shows the window if it was hidden.
func (win *Window) Show() {
	C.RGFW_window_show(win.ref)
}

// Request a window flash to get attention from the user
// func (win *Window) Flash(request FlashRequest) {
//     C.RGFW_window_flash(win.ref, C.RGFW_flashRequest(request))
// }

// Sets whether the window should close.
func (win *Window) SetShouldClose(shouldClose bool) {
	C.RGFW_window_setShouldClose(win.ref, C.RGFW_bool(boolToInt(shouldClose)))
}

func GetGlobalMouse() (x, y int32) {
	var cX, cY C.i32
	C.RGFW_getGlobalMouse(&cX, &cY)
	return int32(cX), int32(cY)
}

// Retrieves the mouse position relative to the window.
// success is true if the position was successfully retrieved
func (win *Window) GetMouse() (x, y int32) {
	var cX, cY C.i32
	C.RGFW_window_getMouse(win.ref, &cX, &cY)
	return int32(cX), int32(cY)
}
// Shows or hides the mouse cursor for the window.
func (win *Window) ShowMouse(show bool) {
	C.RGFW_window_showMouse(win.ref, C.RGFW_bool(boolToInt(show)))
}

// Checks if the mouse is currently hidden in the window.
func (win *Window) IsMouseHidden() bool {
	return C.RGFW_window_isMouseHidden(win.ref) == C.RGFW_TRUE
}

// Moves the mouse to the specified position within the window.
func (win *Window) MoveMouse(x, y int32) {
	C.RGFW_window_moveMouse(win.ref, C.i32(x), C.i32(y))
}

// Checks if the window should close.
func (win *Window) ShouldClose() bool {
	return C.RGFW_window_shouldClose(win.ref) == C.RGFW_TRUE
}

// Checks if the window is currently fullscreen.
func (win *Window) IsFullscreen() bool {
	return C.RGFW_window_isFullscreen(win.ref) == C.RGFW_TRUE
}

// Checks if the window is currently hidden.
func (win *Window) IsHidden() bool {
	return C.RGFW_window_isHidden(win.ref) == C.RGFW_TRUE
}

// Checks if the window is minimized.
func (win *Window) IsMinimized() bool {
	return C.RGFW_window_isMinimized(win.ref) == C.RGFW_TRUE
}

// Checks if the window is maximized.
func (win *Window) IsMaximized() bool {
	return C.RGFW_window_isMaximized(win.ref) == C.RGFW_TRUE
}

// Checks if the window is floating.
func (win *Window) IsFloating() bool {
	return C.RGFW_window_isFloating(win.ref) == C.RGFW_TRUE
}
