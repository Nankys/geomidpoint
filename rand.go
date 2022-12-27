package geomidpoint

import (
	"math"
	"math/rand"
)

// Circular centered at starting point m
func Circular(startLat, startLng, maxDist float64) (float64, float64, float64) {
	rand2 := rand.Float64()
	radian := math.Pi / 180

	radiusEarth := 6372796.924
	maxDist = maxDist / radiusEarth
	brg := 2 * 180 / math.Pi * rand2
	lat := math.Asin(math.Sin(startLat*radian)*math.Cos(maxDist)+math.Cos(startLat*radian)*math.Sin(maxDist)*math.Cos(brg)) * 180 / math.Pi

	lng := startLng + math.Atan2(math.Sin(brg)*math.Sin(maxDist)*math.Cos(startLat*radian), math.Cos(maxDist)-math.Sin(startLat*radian)*math.Sin(lat))
	if lng < -180 {
		lng = lng + 360
	}
	if lng > 180 {
		lng = lng - 360
	}
	return lat, lng, brg
}

// Rectangular  north limit 90 , south limit -90, west limit -180 and east limit 180
func Rectangular(northLat, southLat, westLon, eastLon float64) (float64, float64) {

	lat := southLat + rand.Float64()*(northLat-southLat)

	lng := westLon + (eastLon-westLon)*rand.Float64()

	return lat, lng
}
