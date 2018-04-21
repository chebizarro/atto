package mbgl

import (
	"testing"
)

func TestNewMap(t *testing.T) {

	var pixelRatio float32
	pixelRatio = 1.0
	NewRunLoop()
    fileSource := NewDefaultFileSource("testdata/cache.sqlite", ".")
    //fileSource := NewOnlineFileSource()
    //fileSource.SetAPIBaseUrl("https://osm.tegola.io/")
    threadPool := NewThreadPool(4)
    frontEnd := NewHeadlessFrontend(
		Size{ 256, 256 },
		pixelRatio,
		fileSource,
		threadPool)

	if frontEnd == nil {
		t.Fatal("NewHeadlessFronted returned nil")
	}	
	
    tmap := NewMap(
		frontEnd,
		Size{ 256, 256 },
		pixelRatio,
		fileSource,
		threadPool,
		0, 0, 0)
		
	if tmap == nil {
		t.Fatal("NewMap returned nil")
	}
}
