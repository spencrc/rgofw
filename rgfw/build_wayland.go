//go:build linux && !android && wayland

package rgfw

/*
	#cgo linux CFLAGS: -DRGFW_WAYLAND
	#cgo linux LDFLAGS: -lwayland-client -lwayland-cursor -lwayland-egl -lxkbcommon

	#include "../third_party/xdg-shell.c"
	#include "../third_party/xdg-decoration-unstable-v1.c"
	#include "../third_party/xdg-toplevel-icon-v1.c"
	#include "../third_party/relative-pointer-unstable-v1.c"
	#include "../third_party/pointer-constraints-unstable-v1.c"
	#include "../third_party/xdg-output-unstable-v1.c"
	#include "../third_party/pointer-warp-v1.c"
*/
import "C"
