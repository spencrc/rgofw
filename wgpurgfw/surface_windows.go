//go:build windows

package wgpurgfw

import (
	"syscall"
	"unsafe"

	"github.com/cogentcore/webgpu/wgpu"
	"github.com/spencrc/RGoFW/rgfw"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	pGetModuleHandleW = kernel32.NewProc("GetModuleHandleW")
)

func getModuleHandle() (syscall.Handle, error) {
	ret, _, err := pGetModuleHandleW.Call(uintptr(0))
	if ret == 0 {
		return 0, err
	}
	return syscall.Handle(ret), nil
}

func GetSurfaceDescriptor(w *rgfw.Window) *wgpu.SurfaceDescriptor {
	handle, err := getModuleHandle()
	if err != nil {
		panic(err)
	}

	return &wgpu.SurfaceDescriptor{
		WindowsHWND: &wgpu.SurfaceDescriptorFromWindowsHWND{
			Hwnd:      w.GetHWND(),
			Hinstance: unsafe.Pointer(handle),
		},
	}
}