package raytrace

import (
	// Uncomment for 64bit floats
	// "math"
	math "github.com/chewxy/math32"
)

type Sphere struct {
	Center       Vector3 `json:"center"`
	Radius       float32 `json:"radius"`
	SurfaceColor Vector3 `json:"surfaceColor"`
	EmitColor    Vector3 `json:"emitColor"`
	Transparency float32 `json:"transparency"`
	Reflection   float32 `json:"reflection"`
}

var posInfinity = math.Inf(1)

func (s *Sphere) RadiusSquared() float32{

	return s.Radius * s.Radius
}

func (s *Sphere) Intersect(rayOrigin Vector3, rayDirection Vector3) (bool, float32, float32) {
	l := s.Center.Subtract(rayOrigin)
	tca := l.DotProduct(rayDirection)
	if tca < 0 {
		return false, posInfinity, posInfinity
	}
	d2 := l.DotProduct(l) - (tca * tca)
	if d2 > s.RadiusSquared() {
		return false, posInfinity, posInfinity
	}
	thc := math.Sqrt(s.RadiusSquared() - d2)
	t0 := tca - thc
	t1 := tca + thc

	return true, t0, t1
}