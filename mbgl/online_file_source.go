package mbgl

/*
#include <mbgl-c/storage/online_file_source.h>
#include <stdlib.h>
*/
import "C"

import "unsafe"

type OnlineFileSource C.MbglOnlineFileSource

func NewOnlineFileSource() *OnlineFileSource {
	ofs := OnlineFileSource(*C.mbgl_online_file_source_new())
	return &ofs
}

func (fs *OnlineFileSource) SetAPIBaseUrl(url string) {
	path := C.CString(url)
	defer C.free(unsafe.Pointer(path))
	ofs := _Ctype_struct_MbglOnlineFileSource(*fs)
	C.mbgl_online_file_source_set_api_base_url(
		&ofs,
		path)
}