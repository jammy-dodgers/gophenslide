package openslide

// #cgo CFLAGS: -I/usr/include/openslide
// #cgo LDFLAGS: -lopenslide
// #include <stdio.h>
// #include <stdlib.h>
// #include <openslide.h>
import "C"
import "unsafe"

// DetectVendor
func DetectVendor(filename string) string {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	slide_vendor := C.openslide_detect_vendor(c_filename)
	if slide_vendor == nil {
		return ""
	} else {
		defer C.free(unsafe.Pointer(slide_vendor))
		return C.GoString(slide_vendor)
	}
}
