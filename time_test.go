package nu

import "testing"

func TestDuration(t *testing.T) {
	assertEqual("1h0m0s", Duration(1).String(), t)
	assertEqual("1h2m3s", Duration(1, 2, 3).String(), t)
	assertEqual("1h2m3.004005006s", Duration(1, 2, 3, 4, 5, 6).String(), t)
	assertEqual("1h2m3.004005006s", Duration(1, 2, 3, 4, 5, 6, /* discarded: */ 7, 8).String(), t)
	assertEqual("1h0m0.000000006s", Duration(1, 0, 0, 0, 0, 6).String(), t)
}

func TestDate(t *testing.T) {
	assertEqual("0000-01-01 00:00:00 +0000 UTC", UTCDate().String(), t)
	assertEqual("2000-01-02 00:00:00 +0000 UTC", UTCDate(2000, 1, 2).String(), t)
	assertEqual("2000-01-02 15:16:17 +0000 UTC", UTCDate(2000, 1, 2, 15, 16, 17).String(), t)
	assertEqual("2000-01-02 15:16:17.000000018 +0000 UTC", UTCDate(2000, 1, 2, 15, 16, 17, 18).String(), t)
	assertEqual("1999-12-31 15:16:17.000000018 +0000 UTC", UTCDate(2000, 1, 0, 15, 16, 17, 18).String(), t)
	assertEqual("1999-12-02 15:16:17.000000018 +0000 UTC", UTCDate(2000, 0, 2, 15, 16, 17, 18).String(), t)
	assertEqual("0000-01-02 15:16:17.000000018 +0000 UTC", UTCDate(0, 1, 2, 15, 16, 17, 18).String(), t)
	assertEqual("2000-01-01 00:00:00 +0000 UTC", UTCDate(2000).String(), t)
}

