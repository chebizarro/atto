package mbgl

/*
#include <mbgl.h>
*/
import "C"
import "unsafe"

type DefaultFileSource struct {
	cptr uintptr
}

func (h *DefaultFileSource) cPtr() uintptr {
	return h.cptr
}

func (h *DefaultFileSource) Destroy() {
	C.mbgl_default_file_source_destroy(C.MbglDefaultFileSource(h.cptr))
}

func NewDefaultFileSource(cachePath string, assetRoot string) *DefaultFileSource {
	cpath := C.CString(cachePath)
	aroot := C.CString(assetRoot)
	defer C.free(unsafe.Pointer(cpath)); C.free(unsafe.Pointer(aroot))
	
	dfs := DefaultFileSource{ uintptr(C.mbgl_default_file_source_new(cpath, aroot)) }
	return &dfs
}

func (h *DefaultFileSource) SetAccessToken(token string) {
	ctoken := C.CString(token)
	defer C.free(unsafe.Pointer(ctoken))

	C.mbgl_default_file_source_set_access_token(C.MbglDefaultFileSource(h.cptr), ctoken)
}
