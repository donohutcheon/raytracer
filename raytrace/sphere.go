package raytrace

import (
	"math"
)

type Sphere struct {
	Center       Vector3 `json:"center"`
	Radius       float64 `json:"radius"`
	SurfaceColor Vector3 `json:"surfaceColor"`
	EmitColor    Vector3 `json:"emitColor"`
	Transparency float64 `json:"transparency"`
	Reflection   float64 `json:"reflection"`
}

var posInfinity = math.Inf(1)

func (s *Sphere) RadiusSquared() float64{
	return s.Radius * s.Radius
}

func (s *Sphere) Intersect(rayOrigin Vector3, rayDirection Vector3) (bool, float64, float64) {
	l := s.Center.Subtract(rayOrigin)
	tca := l.DotProduct(rayDirection)
	if tca < 0 {
		return false, math.Inf(1), math.Inf(1)
	}
	d2 := l.DotProduct(l) - (tca * tca)
	if d2 > s.RadiusSquared() {
		return false, math.Inf(1), math.Inf(1)
	}
	thc := math.Sqrt(s.RadiusSquared() - d2)
	t0 := tca - thc
	t1 := tca + thc

	return true, t0, t1
}