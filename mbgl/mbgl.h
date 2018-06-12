#ifndef __MBGL_C_CORE_H
#define __MBGL_C_CORE_H

#include <stdlib.h>
#include "types.h"

#ifdef __cplusplus
extern "C"{
#endif


/**
 * @defgroup MBGL MBGLC: C interface to Mapbox-GL
 *
 * This module exposes parts of the Mapbox-GL library as a C API.
 *
 * @{
 */

/**
 * @defgroup MBGLCCore Core
 *
 * This modules provide an interface to libmbgl-core
 *
 * Many exotic languages can interoperate with C code but have a harder time
 * with C++ due to name mangling. So in addition to C, this interface enables
 * tools written in such languages as Go.
 *
 * @{
 */


/**
 * @defgroup MBGLCCoreTypes Types and Enumerations
 *
 * @{
 */

typedef enum MbglMapMode {
	CONTINUOUS, // continually updating map
	STATIC, // a once-off still image of an arbitrary viewport
	TILE // a once-off still image of a single tile
} MbglMapMode;

typedef enum MbglConstrainMode {
	NONE,
	HEIGHT_ONLY,
	WIDTH_AND_HEIGHT
} MbglConstrainMode;

typedef enum MbglViewportMode {
	DEFAULT,
	FLIPPED_Y
} MbglViewportMode;

typedef enum MbglMapDebugOptions {
	NO_DEBUG     = 0,
	TILE_BORDERS = 1 << 1,
	PARSE_STATUS = 1 << 2,
	TIMESTAMPS   = 1 << 3,
	COLLISION    = 1 << 4,
	OVERDRAW     = 1 << 5
} MbglMapDebugOptions;

/**
 * @}
 */

/**
 * @defgroup MBGLCCoreGeo Geographic primitives
 *
 * @{
 */

MbglLatLng mbgl_lat_lng_new(double lat, double lon);

void mbgl_lat_lng_destroy(MbglLatLng self);

/**
 * @}
 */

/**
 * @defgroup MBGLCCoreFrontend Frontends
 *
 * @{
 */

MbglHeadlessFrontend mbgl_headless_frontend_new(
	const MbglSize size,
	float pixelRatio,
	MbglFileSource fileSource,
	MbglScheduler scheduler);

void mbgl_headless_frontend_destroy(MbglHeadlessFrontend self);

void mbgl_headless_frontend_reset(MbglHeadlessFrontend self);

MbglPremultipliedImage mbgl_headless_frontend_render(MbglHeadlessFrontend self, MbglMap map);

void mbgl_headless_frontend_render_to_file(MbglHeadlessFrontend self, MbglMap map, const char* path);

/**
 * @}
 */

/**
 * @defgroup MBGLCCoreFileSource File Sources
 *
 * @{
 */

MbglDefaultFileSource mbgl_default_file_source_new(const char* cache_path, const char* asset_root);

void mbgl_default_file_source_destroy(MbglDefaultFileSource self);

void mbgl_default_file_source_set_access_token(MbglDefaultFileSource self, const char* token);


MbglOnlineFileSource mbgl_online_file_source_new();

void mbgl_online_file_source_set_api_base_url(MbglOnlineFileSource self, const char* url);

void mbgl_online_file_source_destroy(MbglOnlineFileSource self);

/**
 * @}
 */

/**
 * @defgroup MBGLCCoreMap Map
 *
 * @{
 */

MbglMap mbgl_map_new(
	MbglRendererFrontend renderer,
	MbglMapObserver observer,
	MbglSize size,
	float pixelRatio,
	MbglFileSource source,
	MbglScheduler scheduler,
	MbglMapMode mapMode,
	MbglConstrainMode constrainMode,
	MbglViewportMode viewportMode);

void mbgl_map_destroy(MbglMap self);

MbglStyle mbgl_map_get_style(MbglMap self);

void mbgl_map_set_lat_lng_zoom(MbglMap self, const MbglLatLng latLng, double zoom);

void mbgl_map_set_bearing(MbglMap self, double degrees);

void mbgl_map_set_pitch(MbglMap self, double pitch);

void mbgl_map_set_debug(MbglMap self, MbglMapDebugOptions debugOptions);

void mbgl_map_set_size(MbglMap self, MbglSize size);

/**
 * @}
 */

/**
 * @defgroup MBGLCCoreMapObserver Map Observers
 *
 * @{
 */

MbglMapObserver mbgl_map_observer_null_observer();

void mbgl_map_observer_destroy(MbglMapObserver self);

/**
 * @}
 */

/**
 * @}
 */

/**
 * @defgroup MBGLCCoreStyle Styling
 *
 * @{
 */

MbglStyle mbgl_style_new(MbglScheduler scheduler, MbglFileSource fileSource, float pixelRatio);

void mbgl_style_destroy(MbglStyle self);

void mbgl_style_load_url(MbglStyle self, const char* url);

/**
 * @}
 */

/**
 * @defgroup MBGLCCoreSchedulers Scheduling
 *
 * @{
 */

MbglThreadPool mbgl_thread_pool_new(size_t count);

void mbgl_thread_pool_destroy(MbglThreadPool self);

/**
 * @}
 */

/**
 * @defgroup MBGLCCoreImage Image Handling
 *
 * @{
 */

void mbgl_image_destroy(MbglImage self);

MbglSize mbgl_image_get_size(MbglImage self);

size_t mbgl_image_get_stride(MbglImage self);

size_t mbgl_image_get_bytes(MbglImage self);

unsigned char* mbgl_image_get_data(MbglImage self);

const char* mbgl_encode_png(MbglPremultipliedImage image, size_t* size);

/**
 * @}
 */

/**
 * @defgroup MBGLCCoreRunLoop Run Loop
 *
 * @{
 */

MbglRunLoop mbgl_run_loop_new();

void mbgl_run_loop_destroy(MbglRunLoop self);

/**
 * @}
 */



/**
 * @}
 */

/**
 * @}
 */

#ifdef __cplusplus
}
#endif
#endif /* __MBGL_C_CORE_H */
