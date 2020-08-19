package openslide

import "testing"

const testTiff = "testdata/CMU-1.tiff"

func TestDetectVendor(t *testing.T) {
	vendor, err := DetectVendor(testTiff)
	if err != nil {
		t.Error("Failed to load image: ", err.Error())
	} else if err == nil && vendor == "" {
		t.Error("Err nil but vendor blank")
	}
	t.Log("Vendor: ", vendor)
}

func TestOpen(t *testing.T) {
	slide, err := Open(testTiff)
	defer Close(slide)
	if err != nil {
		t.Error("Failed to load image: ", err.Error())
	}
}
