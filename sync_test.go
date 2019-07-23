package nu

import "testing"

func TestWaitGroup(t *testing.T) {
	group := WaitGroup(1)
	go func() { group.Done() }()
	group.Wait()
	t.Log("Success!")
}
