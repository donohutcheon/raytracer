package main

import (
	"encoding/json"
	"flag"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/donohutcheon/raytracer/raytrace"
	"github.com/pkg/profile"
)

var workers = flag.Int("workers", 8, "Split the workload across n workers")

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	flag.Parse()
	config, err := buildScene()
	if err != nil {
		panic("Could not load scene " + err.Error())
	}

	img, err := raytrace.RenderImage(config, 30.0, *workers)
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