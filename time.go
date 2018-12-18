package nu

import "time"

// Duration builds a duration using up to the first 6 provided values as
// hours, minutes, seconds, milliseconds, microseconds, and nanoseconds.
// Any extra values are ignored. Examples:
// - Duration(1, 2, 3) would return the duration equal to "1h2m3s".
// - Duration(1, 2, 3, 4, 5, 6) would return the duration equal to "1h2m3.004005006s".
func Duration(values ...int) time.Duration {
	var duration time.Duration
	for x, value := range maxLength(len(durationScale), values) {
		duration += durationScale[x] * time.Duration(value)
	}
	return duration
}
func maxLength(max int, values []int) []int {
	if len(values) > max {
		return values[:max]
	}
	return values
}

var durationScale = []time.Duration{
	time.Hour,
	time.Minute,
	time.Second,
	time.Millisecond,
	time.Microsecond,
	time.Nanosecond,
}

// UTCDate builds a time.Time in the UTC timezone using up to the first 6 provided values as
// year, month, day, hour, minute, second, and nanoseconds. Any extra values are ignored. Examples:
// - UTCDate(2000, 1, 2) would return a time.Time equal to "2000-01-02 00:00:00 +0000 UTC"
// - UTCDate(2000, 1, 2, 15, 16, 17, 18) would return a time.Time equal to "2000-01-02 15:16:17.000000018 +0000 UTC"
func UTCDate(values ...int) time.Time {
	for len(values) < 7 {
		if len(values) == 1 || len(values) == 2 {
			values = append(values, 1) // Months and days start at 1. Using 0 would pull the date back awkwardly.
		} else {
			values = append(values, 0)
		}
	}
	return time.Date(
		values[0],             // year
		time.Month(values[1]), // month
		values[2],             // day
		values[3],             // hour
		values[4],             // minute
		values[5],             // second
		values[6],             // nanosecond
		time.UTC,
	)
}
