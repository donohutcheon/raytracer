package raytrace

import (
	"encoding/json"
	// Uncomment for 64bit floats
	// "math"
	math "github.com/chewxy/math32"
)

type Vector3 struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

func (v Vector3) String() (string, error){
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (v Vector3) Normalize() Vector3 {
	normSq := v.LengthSquared()
	newVec := Vector3{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}
	if normSq > 0 {
		invNor := 1 / math.Sqrt(normSq)
		newVec.X *= invNor
		newVec.Y *= invNor
		newVec.Z *= invNor
	}
	return newVec
}

func (v Vector3) LengthSquared() float32{
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector3) Length() float32{
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector3) ScalarMultiply(f float32) Vector3 {
	return Vector3{
		X: v.X * f,
		Y: v.Y * f,
		Z: v.Z * f,
	}
}

func (v Vector3) DotProduct(v2 Vector3) float32 {
	return v.X * v2.X + v.Y * v2.Y + v.Z * v2.Z
}

func (v Vector3) Multiply(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X * v2.X,
		Y: v.Y * v2.Y,
		Z: v.Z * v2.Z,
	}
}

func (v Vector3) Subtract(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
		Z: v.Z - v2.Z,
	}
}

func (v Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{
		X: v2.X + v.X,
		Y: v2.Y + v.Y,
		Z: v2.Z + v.Z,
	}
}

func (v Vector3) Negate(v2 Vector3) Vector3 {
	return Vector3{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}