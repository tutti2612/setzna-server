package location

import (
	"testing"
)

func TestDistance(t *testing.T) {
	type args struct {
		lat1 float64
		lng1 float64
		lat2 float64
		lng2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"渋谷と新宿の距離",
			args{35.658482, 139.701441, 35.690224, 139.700089},
			3.523778708998686,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.lat1, tt.args.lng1, tt.args.lat2, tt.args.lng2); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
