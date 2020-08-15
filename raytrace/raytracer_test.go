package raytrace

import (
	"image/color"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateSurfaceColor(t *testing.T) {
	type params struct {
		reflection    Vector3
		refraction    Vector3
		fresnelEffect float64
		transparency  float64
		surfaceColor  Vector3
	}
	tests := []struct {
		name    string
		params  params
		want    Vector3
		wantErr bool
	}{
		{
			name:    "Test 1",
			params:  params{
				reflection:    Vector3{X: 0.95953, Y: 0.269049, Z: 0.305147},
				refraction:    Vector3{X: 2, Y: 2, Z: 2},
				fresnelEffect: 0.213502,
				transparency:  0.5,
				surfaceColor:  Vector3{X: 1, Y: 0.32, Z: 0.36},
			},
			want:    Vector3{X: 0.99136, Y: 0.270061, Z: 0.306593},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			refl := tt.params.reflection.ScalarMultiply(tt.params.fresnelEffect)
			refr := tt.params.refraction.ScalarMultiply((1.0 - tt.params.fresnelEffect) * tt.params.transparency)
			tmp := refl.Add(refr)
			surfaceColor := tmp.Multiply(tt.params.surfaceColor)
			assert.InDelta(t, tt.want.X, surfaceColor.X, 0.001)
			assert.InDelta(t, tt.want.Y, surfaceColor.Y, 0.001)
			assert.InDelta(t, tt.want.Z, surfaceColor.Z, 0.001)
		})
	}
}

func TestCastRay(t *testing.T) {
	type args struct {
		rayOrigin Vector3
		x         int
		y         int
		width     int
		height    int
		fov       float64
		spheres   []Sphere
	}
	tests := []struct {
		name    string
		args    args
		want    color.Color
		wantErr bool
		skip    bool
	}{
		{
			name:    "Test 1",
			args:    args{
				rayOrigin: Vector3{},
				x:         149,
				y:         351,
				width:     640,
				height:    480,
				fov:       30.0,
				spheres:   buildDefaultScene(t),
			},
			want:    color.RGBA{R: 147, G: 147, B: 147, A: 255},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skip()
			}
			invWidth := 1 / float64(tt.args.width)
			invHeight := 1 / float64(tt.args.height)
			aspectRatio := float64(tt.args.width) * invHeight
			angle := math.Tan(math.Pi * 0.5 * tt.args.fov / 180.0)
			got, err := CastRay(tt.args.x, tt.args.y, invWidth, invHeight, aspectRatio, angle, tt.args.spheres)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}