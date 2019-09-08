package space

// Planet represents the name of the planet on which to check your age
type Planet string

// earthRev represents the number of seconds it takes the earth to complete a revolution around the sun
const earthRev = 31557600

func Age(ageInSecs float64, planet Planet) float64 {

	Planets := map[Planet]float64{
		"Earth":   earthRev,
		"Mercury": earthRev * 0.2408467,
		"Venus":   earthRev * 0.61519726,
		"Mars":    earthRev * 1.8808158,
		"Jupiter": earthRev * 11.862615,
		"Saturn":  earthRev * 29.447498,
		"Uranus":  earthRev * 84.016846,
		"Neptune": earthRev * 164.79132,
	}

	return calcAge(Planets[planet], ageInSecs)
}

func calcAge(planetRev float64, ageInSecs float64) float64 {
	return ageInSecs / planetRev
}
