package mbgl

/*
#cgo CFLAGS: -fPIC
#cgo CFLAGS: -D_GLIBCXX_USE_CXX11_ABI=1
#cgo CXXFLAGS: -std=c++14 -std=gnu++14
#cgo CXXFLAGS: -g
#cgo CXXFLAGS: -I../mason_packages/.link/include
#cgo LDFLAGS: -L../mason_packages/.link/lib
#cgo LDFLAGS: -lstdc++
#cgo LDFLAGS: -lmbgl-filesource -lmbgl-loop-uv -lmbgl-core
#cgo LDFLAGS: -luv -lrt -lpthread -lnsl -ldl -lsqlite3 -lcurl -lGL -lX11 -lnu -lpng16 -lz -lm -ljpeg -lwebp -licuuc -ldl
*/
import "C"

import (
	"bytes"
	"image/png"
	"runtime"
)

// RenderPNG returns a PNG encoded buffer
func RenderPNG(lat, lng, zoom float64, width, height int, url, style string) *bytes.Buffer {

	loop := NewRunLoop()
	defer loop.Destroy()

	fileSource := NewOnlineFileSource()
	fileSource.SetAPIBaseUrl(url)
	defer fileSource.Destroy()

	threadPool := NewThreadPool(runtime.NumCPU())
	defer threadPool.Destroy()

	size := Size{uint32(width), uint32(height)}

	frontEnd := NewHeadlessFrontend(size, 1, fileSource, threadPool)
	defer frontEnd.Destroy()

	pmap := NewMap(frontEnd, size, 1, fileSource, threadPool, Static, HeightOnly, Default)
	defer pmap.Destroy()

	pmap.GetStyle().LoadURL(style)

	latLng := NewLatLng(lat, lng)
	defer latLng.Destroy()

	pmap.SetLatLngZoom(latLng, zoom)
	//pmap.SetBearing(0)
	//pmap.SetPitch(0)

	image := frontEnd.Render(pmap)
	defer image.Destroy()

	var buf bytes.Buffer
	png.Encode(&buf, image)

	return &buf

}
