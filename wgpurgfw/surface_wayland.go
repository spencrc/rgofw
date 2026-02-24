//go:build linux && !android && wayland

package wgpurgfw

import (
	"github.com/cogentcore/webgpu/wgpu"
	"github.com/spencrc/RGoFW/rgfw"
)

func GetSurfaceDescriptor(w *rgfw.Window) *wgpu.SurfaceDescriptor {
	return &wgpu.SurfaceDescriptor{
		WaylandSurface: &wgpu.SurfaceDescriptorFromWaylandSurface{
			Display: rgfw.GetDisplayWayland(),
			Surface: w.GetWindowWayland(),
		},
	}
}
