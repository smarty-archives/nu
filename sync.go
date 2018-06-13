package nu

import "sync"

// WaitGroup returns a *sync.WaitGroup with count as the initial counter value.
func WaitGroup(count int) *sync.WaitGroup {
	group := new(sync.WaitGroup)
	group.Add(count)
	return group
}
