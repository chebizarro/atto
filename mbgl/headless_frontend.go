package mbgl

/*
#include <mbgl.h>
*/
import "C"

import "unsafe"

type HeadlessFrontend struct {
	cptr uintptr
}

func (h *HeadlessFrontend) cPtr() uintptr {
	return h.cptr
}

func (h *HeadlessFrontend) Destroy() {
	C.mbgl_headless_frontend_destroy(C.MbglHeadlessFrontend(h.cptr))
}

func (h *HeadlessFrontend) Reset() {
	C.mbgl_headless_frontend_reset(C.MbglHeadlessFrontend(h.cptr))
}


func (h *HeadlessFrontend) Render(_map Map) *Image {
	img := C.mbgl_headless_frontend_render(
		C.MbglHeadlessFrontend(h.cptr),
		C.MbglMap(_map.cptr))
	
	return newImage(uintptr(img))
}

func (h *HeadlessFrontend) RenderToFile(_map *Map, path string) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	C.mbgl_headless_frontend_render_to_file(
		C.MbglHeadlessFrontend(h.cptr),
		C.MbglMap(_map.cptr),
		cpath)
}

func NewHeadlessFrontend(size Size, pixelRatio float32, fileSource *FileSource, scheduler *Scheduler) *HeadlessFrontend {
	f:= C.mbgl_headless_frontend_new(
		C.MbglSize{C.uint32_t(size.Width), C.uint32_t(size.Height)},
		C.float(pixelRatio),
		C.MbglFileSource(fileSource.cPtr()),
		C.MbglScheduler(scheduler.cPtr()))
	hf := HeadlessFrontend{uintptr(f)}
	return &hf
}


