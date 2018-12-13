package nu

import "testing"

func TestSafeError(t *testing.T) {
	assertEqual("", Error().Error(), t)
	assertEqual("", Error("").Error(), t)
	assertEqual("1", Error(1).Error(), t)
	assertEqual("12", Error("1", "2").Error(), t)
	assertEqual("1 2", Error(1, 2).Error(), t)

	assertEqual("\n", Errorln().Error(), t)
	assertEqual("\n", Errorln("").Error(), t)
	assertEqual("1\n", Errorln("1").Error(), t)
	assertEqual("1 2\n", Errorln("1", "2").Error(), t)
	assertEqual("1 2\n", Errorln(1, 2).Error(), t)

	assertEqual("", Errorf("").Error(), t)
	assertEqual("Hi", Errorf("Hi").Error(), t)
	assertEqual("Hi there", Errorf("Hi %s", "there").Error(), t)
}