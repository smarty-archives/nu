package nu

// Loop returns a slice of empty structs (zero allocations).
// Convenient for simple loops. See github.com/bradfitz/iter
// for x := range nu.Loop(100) { fmt.Println(x) }
func Loop(n int) []struct{} {
	return make([]struct{}, n)
}
