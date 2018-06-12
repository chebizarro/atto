package mbgl

/*
#include <mbgl.h>
*/
import "C"

type LatLng struct {
	cptr uintptr
}

func (l *LatLng) cPtr() uintptr {
	return l.cptr
}

func NewLatLng(lat, lon float64) *LatLng {
	lt := LatLng { uintptr(C.mbgl_lat_lng_new(C.double(lat), C.double(lon))) }
	return &lt
}

func (l *LatLng) Destroy() {
	C.mbgl_lat_lng_destroy(C.MbglLatLng(l.cptr))
}
