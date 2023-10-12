package main

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -I/usr/include/xfce4/libxfce4panel-2.0
// #cgo LDFLAGS: -L/usr/lib/aarch64-linux-gnu -lxfce4panel-2.0
// #include <libxfce4panel/libxfce4panel.h>
import "C"
import "github.com/gotk3/gotk3/gtk"

func main() {
	a := C.libxfce4panel_check_version(1, 2, 1)
	println(*a)
	println(C.XFCE_SCREEN_POSITION_NONE)
	println(C.XFCE_SCREEN_POSITION_NW_H)

}

//export Construct
func Construct(ptr *C.XfcePanelPluginFunc) {
	println("hello I'm registering my plugin")
	gtk.Init(nil)
}
