package main

import (
	"fmt"
	"strings"

	"github.com/spencrc/RGoFW/rgfw"
	"github.com/spencrc/RGoFW/wgpurgfw"
)

func testfunc(win *rgfw.Window, x, y int32) {
	fmt.Println("hi man :)")
}

func main() {
	win := rgfw.CreateWindow("a window", 0, 0, 800, 600, rgfw.WindowFlagCenter | rgfw.WindowFlagNoResize)
	defer win.Close()

	win.SetWindowMovedCallback(testfunc)

	s, err := InitState(win, wgpurgfw.GetSurfaceDescriptor(win))
	if err != nil {
		panic(err)
	}
	defer s.Destroy()
	
	for !win.ShouldClose() {
		event, eventFound := win.CheckEvent()
		for ;eventFound; event, eventFound = win.CheckEvent() {
			if event.Type == rgfw.EventQuit {
				break
			}
		}

		err := s.Render()
		if err != nil {
			fmt.Println("error occurred while rendering:", err)

			errstr := err.Error()
			switch {
			case strings.Contains(errstr, "Surface timed out"): // do nothing
			case strings.Contains(errstr, "Surface is outdated"): // do nothing
			case strings.Contains(errstr, "Surface was lost"): // do nothing
			default:
				panic(err)
			}
		}
	}
}