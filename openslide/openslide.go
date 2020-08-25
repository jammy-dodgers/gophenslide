package openslide

// #cgo CFLAGS: -I/usr/include/openslide
// #cgo LDFLAGS: -lopenslide
// #include <stdio.h>
// #include <stdlib.h>
// #include <stdint.h>
// #include <openslide.h>
// char * str_at(char ** p, int i) { return p[i]; }
import "C"
import (
	"errors"
	"unsafe"
)

// Slide Slides
type Slide struct {
	ptr *C.openslide_t
}

// Open Don't forget to defer Close.
// This is an expensive operation, you will want to cache the result.
func Open(filename string) (Slide, error) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	slideData := C.openslide_open(cFilename)
	if slideData == nil {
		return Slide{nil}, errors.New("File " + filename + " unrecognized.")
	}
	return Slide{slideData}, nil
}

// Close Closes a slide
func (slide Slide) Close() {
	C.openslide_close(slide.ptr)
}

// LevelCount Get the number of levels in the whole slide image.
func (slide Slide) LevelCount() int32 {
	return int32(C.openslide_get_level_count(slide.ptr))
}

// LargestLevelDimensions Get the dimensions of level 0, the largest level (aka get_level0_dimensions)
func (slide Slide) LargestLevelDimensions() (int64, int64) {
	var a, b C.int64_t
	C.openslide_get_level0_dimensions(slide.ptr, &a, &b)
	return int64(a), int64(b)
}

// LevelDimensions Get the dimensions of a level.
func (slide Slide) LevelDimensions(level int32) (int64, int64) {
	var a, b C.int64_t
	C.openslide_get_level_dimensions(slide.ptr, C.int32_t(level), &a, &b)
	return int64(a), int64(b)
}

// LevelDownsample Get the downsampling factor of the given level
func (slide Slide) LevelDownsample(level int32) float64 {
	return float64(C.openslide_get_level_downsample(slide.ptr, C.int32_t(level)))
}

// BestLevelForDownsample Get the best level to use for a particular downsampling factor
func (slide Slide) BestLevelForDownsample(downsample float64) int32 {
	return int32(C.openslide_get_best_level_for_downsample(slide.ptr, C.double(downsample)))
}

// ReadRegion read a region of the image as raw ARGB data
func (slide Slide) ReadRegion(x, y int64, level int32, w, h int64) ([]byte, error) {
	len := w * h * 4
	rawPtr := C.malloc(C.size_t(len))
	defer C.free(rawPtr)
	C.openslide_read_region(slide.ptr, (*C.uint32_t)(rawPtr), C.int64_t(x), C.int64_t(y), C.int32_t(level), C.int64_t(w), C.int64_t(h))
	if txt := C.openslide_get_error(slide.ptr); txt != nil {
		return nil, errors.New(C.GoString(txt))
	}
	return C.GoBytes(rawPtr, C.int(len)), nil
}

// DetectVendor Quickly determine whether a whole slide image is recognized.
func DetectVendor(filename string) (string, error) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	slideVendor := C.openslide_detect_vendor(cFilename)
	if slideVendor == nil {
		return "", errors.New("No vendor for " + filename)
	}
	return C.GoString(slideVendor), nil
}

// PropertyNames Get all property names available for this slide
func (slide Slide) PropertyNames() []string {
	cPropNames := C.openslide_get_property_names(slide.ptr)
	strings := []string{}
	for i := 0; C.str_at(cPropNames, C.int(i)) != nil; i++ {
		strings = append(strings, C.GoString(C.str_at(cPropNames, C.int(i))))
	}
	return strings
}

// PropertyValue Get the value for a specific property
func (slide Slide) PropertyValue(propName string) string {

	cPropName := C.CString(propName)
	defer C.free(unsafe.Pointer(cPropName))
	cPropValue := C.openslide_get_property_value(slide.ptr, cPropName)
	return C.GoString(cPropValue)
}

// Version Get the current version of Openslide as a string
func Version() string {
	cVer := C.openslide_get_version()
	return C.GoString(cVer)
}

// PropBackgroundColor The name of the property containing a slide's background color, if any.
//It is represented as an RGB hex triplet.
const PropBackgroundColor = "openslide.background-color"

// PropBoundsHeight The name of the property containing the height of the rectangle bounding the non-empty region of the slide, if available.
const PropBoundsHeight = "openslide.bounds-height"

// PropBoundsWidth The name of the property containing the width of the rectangle bounding the non-empty region of the slide, if available.
const PropBoundsWidth = "openslide.bounds-width"

// PropBoundsX The name of the property containing the X coordinate of the rectangle bounding the non-empty region of the slide, if available.
const PropBoundsX = "openslide.bounds-x"

// PropBoundsY The name of the property containing the Y coordinate of the rectangle bounding the non-empty region of the slide, if available.
const PropBoundsY = "openslide.bounds-y"

// PropMPPX The name of the property containing the number of microns per pixel in the X dimension of level 0, if known.
const PropMPPX = "openslide.mpp-x"

// PropMPPY The name of the property containing the number of microns per pixel in the Y dimension of level 0, if known.
const PropMPPY = "openslide.mpp-y"

// PropObjectivePower The name of the property containing a slide's objective power, if known.
const PropObjectivePower = "openslide.objective-power"
