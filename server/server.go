//  Package server implements the http frontend
package server

import (
	"http"
	"runtime"
	
	"github.com/dimfeld/httptreemux"

	"github.com/go-spatial/atto/mbgl"	
)

var (
	// Version set at runtime from main
	Version string
	// HostName configurable via the atto config.toml file (set in main.go)
	HostName string
	// Port configurable via the atto config.toml file (set in main.go)
	Port string
	//	CORSAllowedOrigin the "Access-Control-Allow-Origin" CORS header
	//	configurable via the atto config.toml file (set in main.go)
	CORSAllowedOrigin = "*"
)

// Start starts the tile server binding to the provided port
func Start(port string) *http.Server {

	//	notify the user the server is starting
	log.Infof("starting atto server on port %v", port)


	r := httptreemux.New()
	group := r.NewGroup("/")

	// endpoints
	group.UsingContext().Handler(
		"GET",
		"/maps/:map_name/:z/:x/:y",
		CORSHandler(HandleMapZXY{}))

	group.UsingContext().Handler(
		"OPTIONS", 
		"/maps/:map_name/:z/:x/:y", 
		CORSHandler(HandleMapZXY{}))
	

	//	start our server
	srv := &http.Server{Addr: port, Handler: r}
	
	go func() {
		loop := mbgl.NewRunLoop()
		defer loop.Destroy()

		fileSource := mbgl.NewOnlineFileSource()
		fileSource.SetAPIBaseUrl(HostName)
		defer fileSource.Destroy()	

		threadPool := mbgl.NewThreadPool(runtime.NumCPU())
		defer threadPool.Destroy()

		
		
		log.Error(srv.ListenAndServe())
		
	}()
	return srv

}
