package mbgl

/*
#include <mbgl.h>
*/
import "C"

type Map struct {
	cptr uintptr
}

func NewMap(
	renderer RendererFrontend,
	size Size,
	pixelRatio float64,
	source FileSource,
	scheduler Scheduler,
	mapMode MapMode,
	constrainMode ConstrainMode,
	viewportMode ViewportMode) *Map {
	
	nmap := C.mbgl_map_new(
		C.MbglRendererFrontend(renderer.cPtr()),
		C.mbgl_map_observer_null_observer(),
		C.MbglSize{C.uint32_t(size.Width), C.uint32_t(size.Height)},
		C.float(pixelRatio),
		C.MbglFileSource(source.cPtr()),
		C.MbglScheduler(scheduler.cPtr()),
		C.MbglMapMode(mapMode),
		C.MbglConstrainMode(constrainMode),
		C.MbglViewportMode(viewportMode))
	
	gomap := Map{ uintptr(nmap) }
	return &gomap
}

func (m *Map) Destroy() {
	C.mbgl_map_destroy(C.MbglMap(m.cptr))
}

func (m *Map) GetStyle() *Style {
	style := Style{ uintptr(C.mbgl_map_get_style(C.MbglMap(m.cptr))) }
	return &style
}

func (m *Map) SetLatLngZoom(latLng *LatLng, zoom float64) {
	C.mbgl_map_set_lat_lng_zoom(C.MbglMap(m.cptr), C.MbglLatLng(latLng.cptr), C.double(zoom))
}

func (m *Map) SetBearing(degrees float64) {
	C.mbgl_map_set_bearing(C.MbglMap(m.cptr), C.double(degrees))
}

func (m *Map) SetPitch(pitch float64) {
	C.mbgl_map_set_pitch(C.MbglMap(m.cptr), C.double(pitch))
}

func (m *Map) SetDebug(debugOptions MapDebugOptions) {
	C.mbgl_map_set_debug(C.MbglMap(m.cptr), C.MbglMapDebugOptions(debugOptions))
}

func (m *Map) SetSize(size Size) {
	C.mbgl_map_set_size(C.MbglMap(m.cptr), C.MbglSize{C.uint32_t(size.Width), C.uint32_t(size.Height)})
}
