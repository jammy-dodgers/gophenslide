# gophenslide
Go binding for [OpenSlide](https://openslide.org/)

This was a little hacked together for a Uni project and was also my first time writing Go; so the code may not properly follow Go coding conventions.

# Building

Linux (tested on Xubuntu 20.04 LTS):

1. Install dependencies
    - `sudo apt install libopenslide-dev`
1. Grab OpenSlide test data
    - [Download here](http://openslide.cs.cmu.edu/download/openslide-testdata/)
        - You'll want the TIFF image.
    - Place the images in `pkg/testdata`
1. Run tests
    - `go test ./pkg`
