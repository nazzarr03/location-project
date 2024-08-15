package utils

import "math"

func HaversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	lat1 = lat1 * (math.Pi / 180)
	lon1 = lon1 * (math.Pi / 180)
	lat2 = lat2 * (math.Pi / 180)
	lon2 = lon2 * (math.Pi / 180)

	latDiff := lat2 - lat1
	lonDiff := lon2 - lon1

	a := (1 - math.Cos(latDiff)) / 2
	b := math.Cos(lat1) * math.Cos(lat2) * (1 - math.Cos(lonDiff)) / 2
	d := 2 * 6371 * math.Asin(math.Sqrt(a+b))

	return d
}
