package openslide

// #cgo CFLAGS: -I/usr/include/openslide
// #cgo LDFLAGS: -lopenslide
// #include <stdio.h>
// #include <stdlib.h>
// #include <stdint.h>
// #include <openslide.h>
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
func Close(slide Slide) {
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
