package nu

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	counter := 0
	for range Loop(100) {
		counter++
	}
	assertEqual("100", fmt.Sprint(counter), t)
}
