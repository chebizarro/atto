#include <mbgl.h>
#include <mbgl/gl/headless_frontend.hpp>
#include <mbgl/map/map.hpp>
#include <mbgl/map/map_observer.hpp>
#include <mbgl/storage/default_file_source.hpp>
#include <mbgl/storage/online_file_source.hpp>
#include <mbgl/style/style.hpp>
#include <mbgl/util/default_thread_pool.hpp>
#include <mbgl/util/geo.hpp>
#include <mbgl/util/geometry.hpp>
#include <mbgl/util/image.hpp>
#include <mbgl/util/run_loop.hpp>
#include <mbgl/util/size.hpp>

#include <cstdlib>
#include <iostream>
#include <fstream>

using namespace mbgl;

MbglLatLng mbgl_lat_lng_new(double lat, double lon) {
	return reinterpret_cast<MbglLatLng>(new LatLng(lat, lon));
}

void mbgl_lat_lng_destroy(MbglLatLng self) {
	delete reinterpret_cast<LatLng*>(self);
}


MbglHeadlessFrontend mbgl_headless_frontend_new(
	const MbglSize size,
	float pixelRatio,
	MbglFileSource fileSource,
	MbglScheduler scheduler) {
	
	auto fs = reinterpret_cast<FileSource*>(fileSource);
	auto sched = reinterpret_cast<Scheduler*>(scheduler);
	
	return reinterpret_cast<MbglHeadlessFrontend>(
		new HeadlessFrontend(
			{size.width, size.height},
			pixelRatio,
			*fs,
			*sched
		)
	);
}

void mbgl_headless_frontend_destroy(MbglHeadlessFrontend self) {
	delete reinterpret_cast<HeadlessFrontend*>(self);
}

MbglPremultipliedImage mbgl_headless_frontend_render(MbglHeadlessFrontend self, MbglMap map) {
	auto mp = reinterpret_cast<Map*>(map);
	auto image = new PremultipliedImage(reinterpret_cast<HeadlessFrontend*>(self)->render(*mp));
	return reinterpret_cast<MbglPremultipliedImage>(image);
}

void mbgl_headless_frontend_render_to_file(MbglHeadlessFrontend self, MbglMap map, const char* path) {
	auto mp = reinterpret_cast<Map*>(map);
	std::ofstream out(std::string(path), std::ios::binary);
    out << encodePNG(reinterpret_cast<HeadlessFrontend*>(self)->render(*mp));
    out.close();
}

void mbgl_headless_frontend_reset(MbglHeadlessFrontend self) {
	reinterpret_cast<HeadlessFrontend*>(self)->reset();
}


MbglDefaultFileSource mbgl_default_file_source_new(const char* cache_path, const char* asset_root) {
	return reinterpret_cast<MbglDefaultFileSource>(new DefaultFileSource(std::string(cache_path), std::string(asset_root)));
}

void mbgl_default_file_source_destroy(MbglDefaultFileSource self) {
	delete reinterpret_cast<DefaultFileSource*>(self);
}

void mbgl_default_file_source_set_access_token(MbglDefaultFileSource self, const char* token) {
	reinterpret_cast<DefaultFileSource*>(self)->setAccessToken(std::string(token));
}

MbglOnlineFileSource mbgl_online_file_source_new() {
	return reinterpret_cast<MbglOnlineFileSource>(new OnlineFileSource());
}

void mbgl_online_file_source_set_api_base_url(MbglOnlineFileSource self, const char* url) {
	reinterpret_cast<OnlineFileSource*>(self)->setAPIBaseURL(std::string(url));
}

void mbgl_online_file_source_destroy(MbglOnlineFileSource self) {
	delete reinterpret_cast<OnlineFileSource*>(self);
}


MbglMap mbgl_map_new(
	MbglRendererFrontend renderer,
	MbglMapObserver observer,
	MbglSize size,
	float pixelRatio,
	MbglFileSource source,
	MbglScheduler scheduler,
	MbglMapMode mapMode,
	MbglConstrainMode constrainMode,
	MbglViewportMode viewportMode) {

	auto rend = reinterpret_cast<RendererFrontend*>(renderer);
	auto obs = reinterpret_cast<MapObserver*>(observer);
	auto fs = reinterpret_cast<FileSource*>(source);
	auto sched = reinterpret_cast<Scheduler*>(scheduler);

	return reinterpret_cast<MbglMap*>(
		new Map(
			*rend,
			*obs,
			{size.width, size.height},
			pixelRatio,
			*fs,
			*sched,
			static_cast<MapMode>(mapMode),
			static_cast<ConstrainMode>(constrainMode),
			static_cast<ViewportMode>(viewportMode))
		);
}

void mbgl_map_destroy(MbglMap self) {
	delete reinterpret_cast<Map*>(self);
}

MbglStyle mbgl_map_get_style(MbglMap self) {
	style::Style& style = (reinterpret_cast<Map*>(self))->getStyle();
	return reinterpret_cast<MbglStyle>(&style);
}

void mbgl_map_set_lat_lng_zoom(MbglMap self, const MbglLatLng latLng, double zoom) {
	auto llng = reinterpret_cast<const LatLng*>(latLng);
	reinterpret_cast<Map*>(self)->setLatLngZoom(*llng, zoom);		
}

void mbgl_map_set_bearing(MbglMap self, double degrees) {
	(reinterpret_cast<Map*>(self))->setBearing(degrees);	
}

void mbgl_map_set_pitch(MbglMap self, double pitch) {
	(reinterpret_cast<Map*>(self))->setPitch(pitch);
}

void mbgl_map_set_debug(MbglMap self, MbglMapDebugOptions debugOptions) {
	(reinterpret_cast<Map*>(self))->setDebug(
		static_cast<MapDebugOptions>(debugOptions)
	);
}

void mbgl_map_set_size(MbglMap self, MbglSize size) {
	(reinterpret_cast<Map*>(self))->setSize({size.width, size.height});
}

MbglMapObserver mbgl_map_observer_null_observer() {
	return reinterpret_cast<MbglMapObserver>(&MapObserver::nullObserver());
}

void mbgl_map_observer_destroy(MbglMapObserver self) {
	delete reinterpret_cast<MapObserver*>(self);
}

MbglStyle mbgl_style_new(
	MbglScheduler scheduler,
	MbglFileSource fileSource,
	float pixelRatio) {
	
	return reinterpret_cast<MbglStyle>(
		new style::Style(
			reinterpret_cast<Scheduler&>(scheduler),
			reinterpret_cast<FileSource&>(fileSource),
			pixelRatio
		)
	);
}

void mbgl_style_destroy(MbglStyle self) {
	delete reinterpret_cast<style::Style*>(self);
}

void mbgl_style_load_url(MbglStyle self, const char* url) {
	reinterpret_cast<style::Style*>(self)->loadURL(std::string(url));
}

MbglThreadPool mbgl_thread_pool_new(size_t count) {
	return reinterpret_cast<MbglThreadPool>(new ThreadPool(count));
}

void mbgl_thread_pool_destroy(MbglThreadPool self) {
	delete reinterpret_cast<ThreadPool*>(self);
}


void mbgl_image_destroy(MbglImage self) {
	delete reinterpret_cast<PremultipliedImage*>(self);
}

MbglSize mbgl_image_get_size(MbglImage self) {
	auto size = reinterpret_cast<PremultipliedImage*>(self)->size;
	return { size.width, size.height };
}

unsigned char* mbgl_image_get_data(MbglImage self) {
	return &reinterpret_cast<PremultipliedImage*>(self)->data[0];
}

size_t mbgl_image_get_stride(MbglImage self) {
	return reinterpret_cast<PremultipliedImage*>(self)->stride();
}

size_t mbgl_image_get_bytes(MbglImage self) {
	return reinterpret_cast<PremultipliedImage*>(self)->bytes();	
}

const char* mbgl_encode_png(MbglPremultipliedImage image, size_t* size) {
	auto img = encodePNG(reinterpret_cast<const PremultipliedImage&>(image));
	*size = img.length();
	return encodePNG(reinterpret_cast<const PremultipliedImage&>(image)).c_str();
}


MbglRunLoop mbgl_run_loop_new() {
	return reinterpret_cast<MbglRunLoop>(new util::RunLoop());
}

void mbgl_run_loop_destroy(MbglRunLoop self) {
	delete reinterpret_cast<util::RunLoop*>(self);
}
