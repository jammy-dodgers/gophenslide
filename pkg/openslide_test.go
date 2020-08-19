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

func TestLevels(t *testing.T) {
	slide, err := Open(testTiff)
	defer Close(slide)
	if err != nil {
		t.Error("Failed to load image: ", err.Error())
	}
	levels := slide.LevelCount()
	if levels == -1 {
		t.Error("An error has occured :(")
	}
	w, h := slide.LargestLevelDimensions()
	t.Log("Base lvl0 (", w, ", ", h, "): ", slide.LevelDownsample(0))
	for i := int32(1); i < levels; i++ {
		w, h = slide.LevelDimensions(i)
		t.Log("Level ", i, " (", w, ", ", h, "): ", slide.LevelDownsample(i))
	}
	// TODO: make sure these values are actually correct
	// (i don't yet have a program to examine my test files)
}
