package geocode

//goland:noinspection GoUnusedExportedFunction
func NewNoCache() *NoCache {
	return &NoCache{}
}

type NoCache struct{}

func (i NoCache) Exists(string) bool {
	return false
}

func (i NoCache) Get(string) (*Response, error) {
	return NewResponse(), ErrNotFound
}

func (i NoCache) Set(string, *[]*Location) error {
	return nil
}
