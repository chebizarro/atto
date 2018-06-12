package mbgl

import (
	"testing"
)

func TestRenderPNG(t *testing.T) {

	image := RenderPNG(39.153, -76.275, 1024, 512, "https://osm.tegola.io/", "https://raw.githubusercontent.com/go-spatial/tegola-web-demo/master/styles/camo.json")

	if len(image) == 0 {
		t.Fatal("RenderPNG returned no image data")
	}

}
