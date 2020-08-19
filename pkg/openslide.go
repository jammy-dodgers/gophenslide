package openslide

// #cgo CFLAGS: -I/usr/include/openslide
// #cgo LDFLAGS: -lopenslide
// #include <stdio.h>
// #include <stdlib.h>
// #include <openslide.h>
import "C"
import (
	"errors"
	"unsafe"
)

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
