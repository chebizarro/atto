package main

import (
	"flag"
	"runtime"
	"bytes"
	"image/png"
	
	"github.com/go-spatial/atto/mbgl"
)

func main() {

	debugFlag := flag.Bool("debug", false, "Output debugging information")
	pixelRatioFlag := flag.Float64("r", 1, "Pixel ratio for the generated image")
	widthFlag := flag.Uint("w", 512, "Width in pixels of the generated image")
	heightFlag := flag.Uint("h", 512, "Height in pixels of the generated image")
	baseUrlFlag := flag.String("url", "https://osm.tegola.io/", "Url of the Mapbox Vector Tile server")
	styelUrlFlag := flag.String("style", "https://raw.githubusercontent.com/go-spatial/tegola-web-demo/master/styles/camo.json" , "Url of the style to use to render the image")
	zoomFlag := flag.Float64("z", 1, "Zoom level for the generated image")
	pitchFlag := flag.Float64("p", 0, "Pitch for the generated image")
	bearingFlag := flag.Float64("b", 0, "Bearing for the generated image")
	latFlag := flag.Float64("lat", 39.153, "Latitude for the generated image")
	lngFlag := flag.Float64("lng", -76.275, "Longtitude for the generated image")
	outputFlag := flag.String("o", "out.pdf", "File name for the generated image")
	
	flag.Parse()
	
	pixelRatio := *pixelRatioFlag
	zoom := *zoomFlag
	bearing := *bearingFlag
	pitch := *pitchFlag        

	loop := mbgl.NewRunLoop()
	defer loop.Destroy()

	fileSource := mbgl.NewOnlineFileSource()
	fileSource.SetAPIBaseUrl(*baseUrlFlag)
	defer fileSource.Destroy()

	threadPool := mbgl.NewThreadPool(runtime.NumCPU())
	defer threadPool.Destroy()

	size := mbgl.Size{ uint32(*widthFlag), uint32(*heightFlag) }
	
	frontEnd := mbgl.NewHeadlessFrontend(size, pixelRatio, fileSource, threadPool)
	defer frontEnd.Destroy()
	
	pmap := mbgl.NewMap(frontEnd, size, pixelRatio, fileSource, threadPool, mbgl.Static, mbgl.HeightOnly, mbgl.Default)
	defer pmap.Destroy()
	
	pmap.GetStyle().LoadURL(*styelUrlFlag)
	
	latLng := mbgl.NewLatLng(*latFlag,*lngFlag)
	defer latLng.Destroy()
		
	pmap.SetLatLngZoom(latLng, zoom)
	pmap.SetBearing(bearing)
	pmap.SetPitch(pitch)
	
	if *debugFlag {
		pmap.SetDebug(mbgl.TileBorders | mbgl.ParseStatus)
	}
		
	image := frontEnd.Render(pmap)
	defer image.Destroy()
	
	var buf bytes.Buffer
	png.Encode(&buf, image)
	
	pdf(buf.Bytes(), *outputFlag)
}
