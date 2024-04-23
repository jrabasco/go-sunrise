package sunrise

import (
	"testing"
	"time"
)

var dataSunriseSunset = []struct {
	inLatitude         float64
	inLongitude        float64
	inAltitude         float64
	inYear             int
	inMonth            time.Month
	inDay              int
	outSunrise         time.Time
	outSunset          time.Time
	outSunriseAltitude time.Time
	outSunsetAltitude  time.Time
}{
	// 1970-01-01 - prime meridian
	{
		0, 0, 0,
		1970, time.January, 1,
		time.Date(1970, time.January, 1, 5, 59, 54, 0, time.UTC),
		time.Date(1970, time.January, 1, 18, 07, 8, 0, time.UTC),
		time.Date(1970, time.January, 1, 5, 59, 54, 0, time.UTC),
		time.Date(1970, time.January, 1, 18, 07, 8, 0, time.UTC),
	},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{
		43.65, -79.38, 0,
		2000, time.January, 1,
		time.Date(2000, time.January, 1, 12, 50, 59, 0, time.UTC),
		time.Date(2000, time.January, 1, 21, 50, 37, 0, time.UTC),
		time.Date(2000, time.January, 1, 12, 50, 59, 0, time.UTC),
		time.Date(2000, time.January, 1, 21, 50, 37, 0, time.UTC),
	},
	// 2004-04-01 - (52° N, 5° E)
	{
		52, 5, 0,
		2004, time.April, 1,
		time.Date(2004, time.April, 1, 5, 13, 39, 0, time.UTC),
		time.Date(2004, time.April, 1, 18, 13, 28, 0, time.UTC),
		time.Date(2004, time.April, 1, 5, 13, 39, 0, time.UTC),
		time.Date(2004, time.April, 1, 18, 13, 28, 0, time.UTC),
	},
	// 2020-06-15 - Igloolik, Canada
	{
		69.3321443, -81.6781126, 0,
		2020, time.June, 25,
		time.Time{},
		time.Time{},
		time.Time{},
		time.Time{},
	},
	// 1970-01-01 - prime meridian
	{
		0, 0, 1000,
		1970, time.January, 1,
		time.Date(1970, time.January, 1, 5, 59, 54, 0, time.UTC),
		time.Date(1970, time.January, 1, 18, 07, 8, 0, time.UTC),
		time.Date(1970, time.January, 1, 5, 55, 9, 0, time.UTC),
		time.Date(1970, time.January, 1, 18, 11, 53, 0, time.UTC),
	},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{
		43.65, -79.38, 2000,
		2000, time.January, 1,
		time.Date(2000, time.January, 1, 12, 50, 59, 0, time.UTC),
		time.Date(2000, time.January, 1, 21, 50, 37, 0, time.UTC),
		time.Date(2000, time.January, 1, 12, 41, 01, 0, time.UTC),
		time.Date(2000, time.January, 1, 22, 00, 35, 0, time.UTC),
	},
	// 2004-04-01 - (52° N, 5° E)
	{
		52, 5, 3000,
		2004, time.April, 1,
		time.Date(2004, time.April, 1, 5, 13, 39, 0, time.UTC),
		time.Date(2004, time.April, 1, 18, 13, 28, 0, time.UTC),
		time.Date(2004, time.April, 1, 5, 1, 9, 0, time.UTC),
		time.Date(2004, time.April, 1, 18, 25, 59, 0, time.UTC),
	},
	// 2020-06-15 - Igloolik, Canada
	{
		69.3321443, -81.6781126, 4000,
		2020, time.June, 25,
		time.Time{},
		time.Time{},
		time.Time{},
		time.Time{},
	},
}

func TestSunriseSunset(t *testing.T) {
	for _, tt := range dataSunriseSunset {
		vSunrise, vSunset := SunriseSunset(tt.inLatitude, tt.inLongitude, tt.inYear, tt.inMonth, tt.inDay)
		if vSunrise != tt.outSunrise {
			t.Fatalf("%s != %s", vSunrise.String(), tt.outSunrise.String())
		}
		if vSunset != tt.outSunset {
			t.Fatalf("%s != %s", vSunset.String(), tt.outSunset.String())
		}
		vSunriseAlt, vSunsetAlt := SunriseSunsetAltitude(tt.inLatitude, tt.inLongitude, tt.inAltitude, tt.inYear, tt.inMonth, tt.inDay)
		if vSunriseAlt != tt.outSunriseAltitude {
			t.Fatalf("%s != %s", vSunriseAlt.String(), tt.outSunriseAltitude.String())
		}
		if vSunsetAlt != tt.outSunsetAltitude {
			t.Fatalf("%s != %s", vSunsetAlt.String(), tt.outSunsetAltitude.String())
		}
	}
}
