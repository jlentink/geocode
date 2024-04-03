package geocode

type ICache interface {
	// Exists checks if a key exists in the cache
	Exists(key string) bool

	// Get retrieves a value from the cache
	Get(key string) (*Response, error)

	// Set sets a value in the cache
	Set(key string, value []*Location) error
}
