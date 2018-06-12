package mbgl

import (
	"testing"
)

func TestNewMap(t *testing.T) {

	var pixelRatio float64
	pixelRatio = 1.0
	
	// The RunLoop is a hidden dependency of ThreadPool
	loop := NewRunLoop()
    defer loop.Destroy()
    
    //fileSource := NewOnlineFileSource()
    //fileSource.SetAPIBaseUrl("https://osm.tegola.io/")

    fileSource := NewDefaultFileSource("testdata/cache.sqlite", ".")
    defer fileSource.Destroy()
        
    threadPool := NewThreadPool(4)
    defer threadPool.Destroy()
    
    frontEnd := NewHeadlessFrontend(
		Size{ 512, 512 },
		pixelRatio,
		fileSource,
		threadPool)

	defer frontEnd.Destroy()

	if &frontEnd.cptr == nil {
		t.Fatal("NewHeadlessFronted returned nil")
	}	
	
    tmap := NewMap(
		frontEnd,
		Size{ 512, 512 },
		pixelRatio,
		fileSource,
		threadPool,
		0, 0, 0)
	
	defer tmap.Destroy()
	
	if &tmap.cptr == nil {
		t.Fatal("NewMap returned nil")
	}
}
