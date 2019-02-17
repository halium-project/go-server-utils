package uuid

// Producer generate new UUID.
type Producer interface {
	// New UUID.
	New() string
	IsValid(input string) bool
}
