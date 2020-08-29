package raytrace

import (
	"errors"
	"image"
	"image/color"
	// Uncomment for 64bit floats
	// "math"
	math "github.com/chewxy/math32"
)

const maxRayDepth = 15

func mix(a float32, b float32, mix float32) float32 {
	return b * mix + a * (1 - mix)
}

func RenderImage(config *Config, fieldOfView float32) (image.Image, error){
	width := config.Image.Width
	height := config.Image.Height
	r := image.Rect(0, 0, width, height)
	img := image.NewRGBA(r)
	invWidth := 1 / float32(width)
	invHeight := 1 / float32(height)
	aspectRatio := float32(width) / float32(height)

	angle := math.Tan(math.Pi * 0.5 * fieldOfView / 180.0)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c, err := CastRay(x, y, invWidth, invHeight, aspectRatio, angle, config)
			if err != nil {
				return nil, err
			}
			img.Set(x, y, c)
		}
	}
	return img, nil
}

func CastRay(x, y int, invWidth, invHeight, aspectRatio, angle float32, config *Config) (color.Color, error) {
	fx := float32(x)
	fy := float32(y)
	xx := (2 * ((fx + 0.5) * invWidth) - 1) * angle * aspectRatio
	yy := (1 - 2 * ((fy + 0.5) * invHeight)) * angle
	rayDir := Vector3{
		X: xx,
		Y: yy,
		Z: -1,
	}
	rayDir = rayDir.Normalize()
	origin := config.Scene.Camera.Position
	pixel, err := Trace(origin, rayDir, config.Scene.Spheres, 0)

	if err != nil {
		return color.Black, err
	}
	c := color.RGBA{
		R: uint8(math.Min(1, pixel.X) * 255),
		G: uint8(math.Min(1, pixel.Y) * 255),
		B: uint8(math.Min(1, pixel.Z) * 255),
		A: 255,
	}
	return c, nil
}

func Trace(rayOrigin Vector3, rayDirection Vector3, spheres []Sphere, depth int) (Vector3, error) {
	if rayDirection.Length() <= 0.99 || rayDirection.Length() >= 1.01{
		return Vector3{}, errors.New("invalid ray direction length")
	}
	tnear := math.Inf(1)
	var nearest Sphere
	found := false
	for _, s := range spheres {
		intersect, t0, t1 := s.Intersect(rayOrigin, rayDirection)
		if intersect {
			if t0 < 0 {
				t0 = t1
			}
			if t0 < tnear {
				tnear = t0
				nearest = s
				found = true
			}
		}
	}
	if !found {
		return Vector3{
			X: 2,
			Y: 2,
			Z: 2,
		}, nil
	}

	surfaceColor := Vector3{}

	intersect := rayOrigin.Add(rayDirection.ScalarMultiply(tnear))
	normal := intersect.Subtract(nearest.Center).Normalize()

	bias := float32(1e-4) // add some bias to the point from which we will be tracing
	inside := false

	if rayDirection.DotProduct(normal) > 0 {
		normal = normal.ScalarMultiply(-1)
		inside = true
	}

	if (nearest.Transparency > 0 || nearest.Reflection > 0) && depth < maxRayDepth {
		facingRatio := -rayDirection.DotProduct(normal)
		fresnelEffect := mix(math.Pow(1 - facingRatio, 3), 1, 0.1)
		reflectionDir := rayDirection.Subtract(normal.ScalarMultiply(2* rayDirection.DotProduct(normal)))
		reflectionDir = reflectionDir.Normalize()
		newOrigin := intersect.Add(normal.ScalarMultiply(bias))
		reflection, err := Trace(newOrigin, reflectionDir, spheres, depth + 1)
		if err != nil {
			return Vector3{}, err
		}

		refraction := Vector3{}
		if nearest.Transparency > 0 {
			ior := float32(1.1)
			eta := ior
			if !inside {
				eta = 1 / ior
			}
			cosi := -normal.DotProduct(rayDirection)
			k := 1 - eta * eta * (1 - cosi * cosi)
			tmp := rayDirection.ScalarMultiply(eta)
			refractionDir := tmp.Add(normal.ScalarMultiply(eta *  cosi - math.Sqrt(k))).Normalize()
			newOrigin := intersect.Subtract(normal.ScalarMultiply(bias))
			refraction, err = Trace(newOrigin, refractionDir, spheres, depth + 1)
			if err != nil {
				return Vector3{}, err
			}
		}
		surfaceColor = reflection.ScalarMultiply(fresnelEffect).
			Add(refraction.ScalarMultiply((1.0 - fresnelEffect) * nearest.Transparency)).
			Multiply(nearest.SurfaceColor)
	} else {
		for _, s := range spheres {
			if s.EmitColor.X > 0 {
				transmission := Vector3{X: 1, Y: 1, Z: 1}
				lightDirection := s.Center.Subtract(intersect).Normalize()
				for _, t := range spheres {
					if s != t {
						intersect2, _, _ := t.Intersect(intersect.Add(normal.ScalarMultiply(bias)), lightDirection)
						if intersect2 {
							transmission = Vector3{X: 0, Y: 0, Z: 0}
							break
						}
					}
				}
				max := math.Max(0.0, normal.DotProduct(lightDirection))
				surfaceColor = surfaceColor.
					Add(nearest.SurfaceColor.
						Multiply(transmission.
							ScalarMultiply(max).
							Multiply(s.EmitColor)))
			}
		}
	}

	return surfaceColor.Add(nearest.EmitColor), nil
}