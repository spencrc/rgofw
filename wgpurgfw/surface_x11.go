//go:build linux && !android && !wayland

package wgpurgfw

import (
	"github.com/cogentcore/webgpu/wgpu"
	"github.com/spencrc/RGoFW/rgfw"
)

func GetSurfaceDescriptor(w *rgfw.Window) *wgpu.SurfaceDescriptor {
	return &wgpu.SurfaceDescriptor{
		XlibWindow: &wgpu.SurfaceDescriptorFromXlibWindow{
			Display: rgfw.GetDisplayX11(),
			Window:  w.GetWindowX11(),
		},
	}
}