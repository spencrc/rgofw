package rgfw

/*
	#include "RGFW.h"
*/
import "C"

type MonitorMode struct {
	W, H             int32 /*!< monitor workarea size */
	RefreshRate      uint32
	Red, Blue, Green uint8
}

func (m MonitorMode) toC() C.RGFW_monitorMode {
	return C.RGFW_monitorMode{
		w:           C.i32(m.W),
		h:           C.i32(m.H),
		refreshRate: C.u32(m.RefreshRate),
		red:         C.u8(m.Red),
		green:       C.u8(m.Green),
		blue:        C.u8(m.Blue),
	}
}

type Monitor struct {
	X, Y           int32   /*!< x - y of the monitor workarea */
	Name           string  /*!< monitor name */
	ScaleX, ScaleY float32 /*!< monitor content scale */
	PixelRatio     float32 /*!< pixel ratio for monitor (1.0 for regular, 2.0 for hiDPI)  */
	PhysW, PhysH   float32 /*!< monitor physical size in inches */
	Mode           MonitorMode
}

func goMonitor(m C.RGFW_monitor) Monitor {
	mode := MonitorMode{
		W:           int32(m.mode.w),
		H:           int32(m.mode.h),
		RefreshRate: uint32(m.mode.refreshRate),
		Red:         uint8(m.mode.red),
		Blue:        uint8(m.mode.blue),
		Green:       uint8(m.mode.green),
	}
	return Monitor{
		X:          int32(m.x),
		Y:          int32(m.y),
		Name:       C.GoStringN(&m.name[0], 128),
		ScaleX:     float32(m.scaleX),
		ScaleY:     float32(m.scaleY),
		PixelRatio: float32(m.pixelRatio),
		PhysW:      float32(m.physW),
		PhysH:      float32(m.physH),
		Mode:       mode,
	}
}

func (m Monitor) toC() C.RGFW_monitor {
	name := [128]C.char{}

	for i := 0; i < len(m.Name) && i < 128; i++ {
		name[i] = C.char(m.Name[i])
	}

	return C.RGFW_monitor{
		x:          C.i32(m.X),
		y:          C.i32(m.Y),
		name:       name,
		scaleX:     C.float(m.ScaleX),
		scaleY:     C.float(m.ScaleY),
		pixelRatio: C.float(m.PixelRatio),
		physW:      C.float(m.PhysW),
		physH:      C.float(m.PhysH),
		mode:       m.Mode.toC(),
	}
}

type ModeRequest uint8

const (
	MonitorScale   ModeRequest = C.RGFW_monitorScale
	MonitorRefresh ModeRequest = C.RGFW_monitorRefresh
	MonitorRGB     ModeRequest = C.RGFW_monitorRGB
	MonitorAll     ModeRequest = C.RGFW_monitorAll
)

// Scales the window to match its monitor’s resolution.
func (win *Window) ScaleToMonitor() {
	C.RGFW_window_scaleToMonitor(win.ref)
}

// Retrieves the monitor structure associated with the window.
func (win *Window) GetMonitor() Monitor {
	m := C.RGFW_window_getMonitor(win.ref)
	return goMonitor(m)
}
