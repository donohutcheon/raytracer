package raytrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSphere_Intersect(t *testing.T) {
	type fields struct {
		Center       Vector3
		Radius       float32
		SurfaceColor Vector3
		ReflectColor Vector3
		Transparency float32
		Reflection   float32
	}
	type args struct {
		rayOrigin    Vector3
		rayDirection Vector3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  float32
		want2  float32
	}{
		{
			name:   "Intersect 1",
			fields: fields{
				Center:       Vector3{0,0, -20},
				Radius:       4,
				SurfaceColor: Vector3{1,0.32, 0.36},
				ReflectColor: Vector3{0,0,0},
				Transparency: 0.5,
				Reflection:   1,
			},
			args:   args{
				rayOrigin:    Vector3{},
				rayDirection: Vector3{-0.0217657071, -0.000558109139, -0.999762893},
			},
			want:   true,
			want1:  16.0190334,
			want2:  23.9714832,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sphere{
				Center:       tt.fields.Center,
				Radius:       tt.fields.Radius,
				SurfaceColor: tt.fields.SurfaceColor,
				EmitColor:    tt.fields.ReflectColor,
				Transparency: tt.fields.Transparency,
				Reflection:   tt.fields.Reflection,
			}
			got, t0, t1 := s.Intersect(tt.args.rayOrigin, tt.args.rayDirection)
			assert.InDelta(t, tt.want1, t0, 0.001)
			assert.InDelta(t, tt.want2, t1, 0.001)
			assert.Equal(t, tt.want, got)
		})
	}
}
