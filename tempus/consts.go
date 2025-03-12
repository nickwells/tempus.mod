package tempus

const (
	SecondsPerMinute = 60
	MinutesPerHour   = 60
	SecondsPerHour   = SecondsPerMinute * MinutesPerHour
	HoursPerDay      = 24
	MinutesPerDay    = MinutesPerHour * HoursPerDay
	SecondsPerDay    = SecondsPerMinute * MinutesPerDay
	DaysPerWeek      = 7
	HoursPerWeek     = HoursPerDay * DaysPerWeek
	MinutesPerWeek   = MinutesPerDay * DaysPerWeek
	SecondsPerWeek   = SecondsPerDay * DaysPerWeek
)
