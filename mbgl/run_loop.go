package mbgl

/*
#include <mbgl.h>
*/
import "C"

type RunLoop struct {
	cptr uintptr
}

func (r *RunLoop) cPtr() uintptr {
	return r.cptr
}

func (r *RunLoop) Destroy() {
	C.mbgl_run_loop_destroy(C.MbglRunLoop(r.cptr))
}

func NewRunLoop() *RunLoop {
	rl := RunLoop{ uintptr(C.mbgl_run_loop_new()) }
	return &rl
}
