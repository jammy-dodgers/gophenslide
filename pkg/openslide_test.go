package openslide

import "testing"

func TestDetectVendor(t *testing.T) {
	const testData = "testdata/CMU-1.tiff"
	vendor, err := DetectVendor(testData)
	if err != nil {
		t.Error("Failed to load image ", err.Error())
	} else if err == nil && vendor == "" {
		t.Error("Err nil but vendor blank")
	}
	t.Log("Vendor: ", vendor)
}
