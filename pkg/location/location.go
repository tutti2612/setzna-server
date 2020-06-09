package location

import (
	"math"
)

// 2点間の距離を計算する。単位km
func Distance(lat1, lng1, lat2, lng2 float64) float64 {
	// 緯度1度あたり110.9463km、経度1度あたり90.4219km
	// https://s-giken.info/distance/distance.php
	// HACK: 地球の丸みを考慮すると、三平方の定理では不十分なので改善する。
	return math.Sqrt(math.Pow((lat1-lat2)*110.9463, 2) + math.Pow((lng1-lng2)*90.4219, 2))
}
