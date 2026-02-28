//go:build linux && !android && wayland

package rgfw

/*
	#cgo linux CFLAGS: -DRGFW_WAYLAND
	#cgo linux LDFLAGS: -lwayland-client -lwayland-cursor -lwayland-egl -lxkbcommon

	#include "xdg-shell.c"
	#include "xdg-decoration-unstable-v1.c"
	#include "xdg-toplevel-icon-v1.c"
	#include "relative-pointer-unstable-v1.c"
	#include "pointer-constraints-unstable-v1.c"
	#include "xdg-output-unstable-v1.c"
	#include "pointer-warp-v1.c"
*/
import "C"
