package providers

import "github.com/homayoonalimohammadi/care-companion/internal/app/db/queries"

type CareSeek interface {
	Get(id int64) (*queries.CareSeek, error)
	Create(*queries.CareSeek) error
}

type CareSeekProvider struct{}

func (p *CareSeekProvider) Get(id int64) (*queries.CareSeek, error) {
	return nil, nil
}

func (p *CareSeekProvider) Create(cs *queries.CareSeek) error {
	return nil
}

func NewCareSeek(querier queries.Querier) *CareSeekProvider {
	return &CareSeekProvider{}
}
