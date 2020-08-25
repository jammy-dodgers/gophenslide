package openslide

import (
	"io/ioutil"
	"os"
	"testing"
)

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
	defer slide.Close()
	if err != nil {
		t.Error("Failed to load image: ", err.Error())
	}
}

func TestLevels(t *testing.T) {
	slide, err := Open(testTiff)
	defer slide.Close()
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

func TestReadRegion(t *testing.T) {
	slide, err := Open(testTiff)
	defer slide.Close()
	if err != nil {
		t.Error("Failed to load image: ", err.Error())
	}
	bytes, err := slide.ReadRegion(10, 10, 6, 400, 400)
	if err != nil {
		t.Fatal(err.Error())
	}
	const testRawFilename = "testdata/raw_region.data"
	if info, e := os.Stat(testRawFilename); os.IsExist(e) && !info.IsDir() {
		if remErr := os.Remove(testRawFilename); remErr != nil {
			t.Log("Could not remove file ", testRawFilename)
		}
	}
	writeErr := ioutil.WriteFile(testRawFilename, bytes, 0660)
	if writeErr != nil {
		t.Fatal(writeErr.Error())
	}
}

func TestProperties(t *testing.T) {
	slide, err := Open(testTiff)
	defer slide.Close()
	if err != nil {
		t.Error("Failed to load image: ", err.Error())
	}
	props := slide.PropertyNames()
	for i := 0; i < len(props); i++ {
		t.Log(props[i], "=", slide.PropertyValue(props[i]))
	}
}

func TestVersion(t *testing.T) {
	t.Log("Version", Version())
}
