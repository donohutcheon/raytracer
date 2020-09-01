package raytrace

import (
	"image/color"
	"runtime"
	"testing"

	math "github.com/chewxy/math32"
	"github.com/stretchr/testify/assert"
)

func Test_CalculateSurfaceColor(t *testing.T) {
	type params struct {
		reflection    Vector3
		refraction    Vector3
		fresnelEffect float32
		transparency  float32
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
		fov       float32
		config    Config
	}
	tests := []struct {
		name    string
		args    args
		want    color.Color
		wantErr bool
	}{
		{
			name:    "Test 1",
			args:    args{
				rayOrigin: Vector3{},
				x:         149,
				y:         351,
				fov:       30.0,
				config:    Config{
					Image: Image{
						Width:  640,
						Height: 480,
					},
					Scene: Scene{
						Spheres: buildDefaultScene(),
					},
				},
			},
			want:    color.RGBA{R: 147, G: 147, B: 147, A: 255},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invWidth := 1 / float32(tt.args.config.Image.Width)
			invHeight := 1 / float32(tt.args.config.Image.Height)
			aspectRatio := float32(tt.args.config.Image.Width) * invHeight
			angle := math.Tan(math.Pi * (0.5 * tt.args.fov) / 180.0)
			got, err := CastRay(tt.args.x, tt.args.y, invWidth, invHeight, aspectRatio, angle, &tt.args.config)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkRenderImage(b *testing.B) {
	type args struct {
		config      *Config
		fieldOfView float32
		workers     int
	}
	tests := []struct {
		name    string
		args    args
	}{
		{
			name:    "Benchmark 1200*900",
			args:    args{
				config:      &Config{
					Image: Image{
						Width:  1200,
						Height: 900,
					},
					Scene: Scene{
						Spheres: buildDefaultScene(),
					},
				},
				fieldOfView: 30.0,
				workers:     runtime.NumCPU(),
			},
		},
		{
			name:    "Benchmark 640*480",
			args:    args{
				config:      &Config{
					Image: Image{
						Width:  640,
						Height: 480,
					},
					Scene: Scene{
						Spheres: buildDefaultScene(),
					},
				},
				fieldOfView: 30.0,
				workers:     runtime.NumCPU(),
			},
		},
		{
			name:    "Benchmark 1024*768",
			args:    args{
				config:      &Config{
					Image: Image{
						Width:  1024,
						Height: 768,
					},
					Scene: Scene{
						Spheres: buildDefaultScene(),
					},
				},
				fieldOfView: 30.0,
				workers:     runtime.NumCPU(),
			},
		},
	}
	for _, tt := range tests {
		b.Log("Benchmark iterations = ", b.N)
		b.Run(tt.name, func(b *testing.B) {
			RenderImage(tt.args.config, tt.args.fieldOfView, tt.args.workers)
		})
	}
}