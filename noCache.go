package geocode

func NewNoCache() *NoCache {
	return &NoCache{}
}

type NoCache struct {
}

func (i NoCache) Exists(key string) bool {
	return false
}

func (i NoCache) Get(key string) (*Response, error) {
	return NewResponse(), ErrNotFound
}

func (i NoCache) Set(key string, value *[]*Location) error {
	return nil
}
