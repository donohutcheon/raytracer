package raytrace

import (
	"encoding/json"
	"fmt"
	"testing"
)

func buildDefaultScene(t *testing.T) []Sphere {
	t.Helper()
	var spheres []Sphere
	spheres = append(spheres, Sphere{
		Center:       Vector3{X: 0, Y: -10004, Z: -20},
		Radius:       10000,
		SurfaceColor: Vector3{X: 0.2, Y: 0.2, Z: 0.2},
		EmitColor:    Vector3{},
		Transparency: 0,
		Reflection:   0.0,
	})
	spheres = append(spheres, Sphere{
		Center:       Vector3{X: 0,Y: 0,Z: -20},
		Radius:       4,
		SurfaceColor: Vector3{X: 1, Y: 0.32, Z: 0.36},
		EmitColor:    Vector3{},
		Transparency: 1,
		Reflection:   0.5,
	})
	spheres = append(spheres, Sphere{
		Center:       Vector3{X: 5, Y: -1, Z: -15,
		},
		Radius:       2,
		SurfaceColor: Vector3{X: 0.9, Y: 0.76, Z: 0.46},
		EmitColor:    Vector3{},
		Transparency: 0,
		Reflection:   1.0,
	})
	spheres = append(spheres, Sphere{
		Center:       Vector3{X: 5, Y: 0, Z: -25},
		Radius:       3,
		SurfaceColor: Vector3{X: 0.65, Y: 0.77, Z: 0.97},
		EmitColor:    Vector3{},
		Transparency: 0,
		Reflection:   1.0,
	})
	spheres = append(spheres, Sphere{
		Center:       Vector3{X: -5.5, Y: 0, Z: -15},
		Radius:       3,
		SurfaceColor: Vector3{X: 0.65, Y: 0.77, Z: 0.97},
		EmitColor:    Vector3{},
		Transparency: 0,
		Reflection:   1.0,
	})
	spheres = append(spheres, Sphere{
		Center:       Vector3{X: 0, Y: 20,	Z: -30},
		Radius:       3,
		SurfaceColor: Vector3{X: 0.0, Y: 0.0, Z: 0.0},
		EmitColor:    Vector3{X: 3, Y: 3, Z: 3},
		Transparency: 0,
		Reflection:   0.0,
	})
	b, _ := json.Marshal(spheres)
	fmt.Println(string(b))
	return spheres
}