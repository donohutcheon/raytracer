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

	config, err := buildScene()
	if err != nil {
		panic("Could not load scene " + err.Error())
	}

	img, err := raytrace.RenderImage(config, 30.0)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("example.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func buildScene() (*raytrace.Config, error)  {
	b, err := ioutil.ReadFile("scene.json")
	if err != nil {
		return nil, err
	}

	config := new(raytrace.Config)
	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}