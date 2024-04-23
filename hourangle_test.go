package sunrise

import (
	"testing"
)

var dataHourAngle = []struct {
	inLatitude    float64
	inDeclination float64
	inAltitude    float64
	out           float64
	outAltitude   float64
}{
	// 1970-01-01 - prime meridian
	{0, -22.97753, 0, 90.904793, 90.904793},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{43.65, -23.01689, 0, 67.453649, 67.453649},
	// 2004-04-01 - (52° N, 5° E)
	{52, 4.75374, 0, 97.477355, 97.477355},
	// 1970-01-01 - prime meridian
	{0, -22.97753, 1000, 90.904793, 92.0933},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{43.65, -23.01689, 2000, 67.453649, 69.946618},
	// 2004-04-01 - (52° N, 5° E)
	{52, 4.75374, 4000, 97.477355, 101.089676},
}

func TestHourAngle(t *testing.T) {
	for _, tt := range dataHourAngle {
		v := HourAngle(tt.inLatitude, tt.inDeclination)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}

		vAlt := HourAngleAltitude(tt.inLatitude, tt.inDeclination, tt.inAltitude)
		if Round(vAlt, DefaultPlaces) != Round(tt.outAltitude, DefaultPlaces) {
			t.Fatalf("%f != %f", vAlt, tt.outAltitude)
		}
	}
}
