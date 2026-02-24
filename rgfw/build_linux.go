//go:build linux && !android && !wayland

package rgfw

/*
	#cgo linux LDFLAGS: -lX11 -lXrandr
*/
import "C"
