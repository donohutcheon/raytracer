package main

import (
	"encoding/json"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/donohutcheon/raytracer/raytrace"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	width := 1700
	height := 900

	spheres, err := buildScene()
	if err != nil {
		panic("Could not load scene " + err.Error())
	}
	img, err := raytrace.RenderImage(width, height, spheres, 30.0)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("outimage.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func buildScene() ([]raytrace.Sphere, error)  {
	b, err := ioutil.ReadFile("scene.json")
	if err != nil {
		return nil, err
	}

	var spheres []raytrace.Sphere
	err = json.Unmarshal(b, &spheres)
	if err != nil {
		return nil, err
	}

	return spheres, nil
}