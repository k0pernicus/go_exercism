package clock

import (
	"fmt"
)

/*Hours represents a time hour, between 0 and 24
 */
type Hours = uint8

/*Minutes represents a time minute, between 0 and 60
 */
type Minutes = uint8

const (
	maxHours   int = 24
	minHours   int = 0
	minMinutes int = 0
	maxMinutes int = 60
)

func getHours(hours int) Hours {
	hours = hours % maxHours
	if hours >= minHours {
		return uint8(hours)
	}
	return uint8(maxHours + hours)
}

func getMinutes(minutes int) (int, Minutes) {
	hours, minutes := minutes/maxMinutes, minutes%maxMinutes
	if minutes >= minMinutes {
		return hours, uint8(minutes)
	}
	return hours + ((-60 + minutes) / maxMinutes), uint8(maxMinutes + minutes)
}

/*Clock is a data structure that handles times (without dates).
 *It contains two fields: Hours (time hour), and Minutes (time minutes)
 */
type Clock struct {
	hours   Hours
	minutes Minutes
}

/*New is a function that, based on two integer values, returns a Clock struct
 */
func New(hours, minutes int) Clock {
	sHours, nMinutes := getMinutes(minutes)
	nHours := getHours(hours + sHours)
	return Clock{
		nHours,
		nMinutes,
	}
}

/*String returns the string formatting of a given Clock struct
 */
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

/*Add adds minutes to an existing Clock struct, and returns the new one.
 *This function handles subtraction by accepting negative values.
 */
func (c Clock) Add(m int) Clock {
	iHours, iMinutes := int(c.hours), int(c.minutes)
	sHours, nMinutes := getMinutes(iMinutes + m)
	return Clock{
		getHours(iHours + sHours),
		nMinutes,
	}
}

/*Subtract subtracts minutes to an existing Clock struct, and returns the new one.
 *This function handles subtraction by accepting negative values.
 */
func (c Clock) Subtract(m int) Clock {
	iHours, iMinutes := int(c.hours), int(c.minutes)
	sHours, nMinutes := getMinutes(iMinutes - m)
	return Clock{
		getHours(iHours + sHours),
		nMinutes,
	}
}
