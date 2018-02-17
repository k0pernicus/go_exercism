package space

type orbitalPeriodEarthDays = float64

/*Planet is an alias for a string that is the name of a planet
 */
type Planet string

const earth orbitalPeriodEarthDays = 31557600.

type orbitalPeriodEarthYears = float64

const (
	mercury orbitalPeriodEarthYears = 0.2408467
	venus   orbitalPeriodEarthYears = 0.61519726
	mars    orbitalPeriodEarthYears = 1.8808158
	jupiter orbitalPeriodEarthYears = 11.862615
	saturn  orbitalPeriodEarthYears = 29.447498
	uranus  orbitalPeriodEarthYears = 84.016846
	neptune orbitalPeriodEarthYears = 164.79132
)

/*Age returns how old someone is in terms of a given planet's solar years.
 */
func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return seconds / earth
	case "Mercury":
		return seconds / (earth * mercury)
	case "Venus":
		return seconds / (earth * venus)
	case "Mars":
		return seconds / (earth * mars)
	case "Jupiter":
		return seconds / (earth * jupiter)
	case "Saturn":
		return seconds / (earth * saturn)
	case "Uranus":
		return seconds / (earth * uranus)
	case "Neptune":
		return seconds / (earth * neptune)
	default:
		return 0.
	}
}
