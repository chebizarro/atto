package mbgl

/*
#include <mbgl.h>
*/
import "C"

type ThreadPool struct {
	cptr uintptr
}

func (tp *ThreadPool) cPtr() uintptr {
	return tp.cptr
}

func (tp *ThreadPool) Destroy() {
	C.mbgl_thread_pool_destroy(C.MbglThreadPool(tp.cptr))
}

func NewThreadPool(count int) *ThreadPool {
	tp := ThreadPool{ uintptr(C.mbgl_thread_pool_new(C.size_t(count))) };
	return &tp
}
