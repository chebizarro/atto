package mbgl

/*
#include <mbgl.h>
*/
import "C"

import "unsafe"

//type OnlineFileSource C.MbglOnlineFileSource
type OnlineFileSource struct {
	cptr uintptr
}

func (o *OnlineFileSource) cPtr() uintptr {
	return o.cptr
}

func (o *OnlineFileSource) Destroy() {
	C.mbgl_online_file_source_destroy(C.MbglOnlineFileSource(o.cptr))
}

func NewOnlineFileSource() *OnlineFileSource {
	ofs := OnlineFileSource{uintptr(C.mbgl_online_file_source_new())}
	return &ofs
}

func (o *OnlineFileSource) SetAPIBaseUrl(url string) {
	path := C.CString(url)
	defer C.free(unsafe.Pointer(path))
	ofs := C.MbglOnlineFileSource(o.cptr)
	C.mbgl_online_file_source_set_api_base_url(
		ofs,
		path)
}
