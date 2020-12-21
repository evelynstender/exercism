package space

import (
	"math"
)

type Planet string

var planetsOrbitalPeriodYears = []struct {
	planet        Planet
	orbitalPeriod float64
}{
	{
		planet:        "Earth",
		orbitalPeriod: 1.0,
	},
	{

		planet:        "Mercury",
		orbitalPeriod: 0.2408467,
	},
	{
		planet:        "Venus",
		orbitalPeriod: 0.61519726,
	},
	{
		planet:        "Mars",
		orbitalPeriod: 1.8808158,
	},
	{

		planet:        "Jupiter",
		orbitalPeriod: 11.862615,
	},
	{
		planet:        "Saturn",
		orbitalPeriod: 29.447498,
	},
	{
		planet:        "Uranus",
		orbitalPeriod: 84.016846,
	},
	{

		planet:        "Neptune",
		orbitalPeriod: 164.79132,
	},
}

// Age returns the age in the solar system planets given seconds
func Age(seconds float64, planet Planet) float64 {

	for _, p := range planetsOrbitalPeriodYears {
		if p.planet == planet {
			yearsOnEarth := seconds / 31557600

			if planet == "Earth" {
				return math.Abs(yearsOnEarth)
			}

			return math.Abs(yearsOnEarth / p.orbitalPeriod)
		}
	}

	return -1

}
