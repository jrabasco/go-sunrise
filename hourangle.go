package sunrise

import (
	"math"
)

// HourAngle calculates the second of the two angles required to locate a point
// on the celestial sphere in the equatorial coordinate system.
func HourAngle(latitude, declination float64) float64 {
	return HourAngleAltitude(latitude, declination, 0)
}

// HourAngleAltitude calculates the second of the two angles required to
// locate a point on the celestial sphere in the equatorial coordinate system
// while correcting for the observer's altitude (in meters)
func HourAngleAltitude(latitude, declination, altitude float64) float64 {
	var (
		latitudeRad    = latitude * Degree
		declinationRad = declination * Degree
		numRad         = -0.0145385927 // -0.833 degrees to radians
		altCorrection  = radCorrection(altitude)
		numerator      = math.Sin(numRad+altCorrection) - math.Sin(latitudeRad)*math.Sin(declinationRad)
		denominator    = math.Cos(latitudeRad) * math.Cos(declinationRad)
	)

	// Check for no sunrise/sunset
	if numerator/denominator > 1 {
		// Sun never rises
		return math.MaxFloat64
	}

	if numerator/denominator < -1 {
		// Sun never sets
		return -1 * math.MaxFloat64
	}

	return math.Acos(numerator/denominator) / Degree
}

func radCorrection(altitude float64) float64 {
	degrees := -2.076 * math.Sqrt(altitude) / 60
	return (degrees / 360) * math.Pi * 2
}
