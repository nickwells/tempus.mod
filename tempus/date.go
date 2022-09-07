package tempus

import "time"

// DaysInYear returns the number of days in the year that the time lies
// within.
func DaysInYear(t time.Time) int {
	days := 365
	if IsLeapYear(t) {
		days++
	}
	return days
}

// DaysInMonth returns the number of days in the month that the time lies
// within.
func DaysInMonth(t time.Time) int {
	daysInMonth := []int{
		// Jan Feb Mar Apr May Jun Jul Aug Sep Oct Nov Dec
		0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
	}
	month := t.Month()
	days := daysInMonth[month]
	if month == time.February && IsLeapYear(t) {
		days++
	}
	return days
}

// IsLeapYear returns true if the given year is a leap year.
func IsLeapYear(t time.Time) bool {
	y := t.Year()
	// years divisible by 400 (such as 2000) are leap years
	if y%400 == 0 {
		return true
	}
	// years divisible by 4 but not 100 are leap years
	if y%100 != 0 && y%4 == 0 {
		return true
	}
	// other years are not leap years
	return false
}
