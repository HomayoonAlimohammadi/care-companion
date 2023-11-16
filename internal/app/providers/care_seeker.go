package providers

import "github.com/homayoonalimohammadi/care-companion/internal/app/db/queries"

type CareSeeker interface {
	Get(id int64) (*queries.CareSeeker, error)
	Create(*queries.CareSeeker) error
}

type CareSeekerProvider struct{}

func (p *CareSeekerProvider) Get(id int64) (*queries.CareSeeker, error) {
	return nil, nil
}

func (p *CareSeekerProvider) Create(cs *queries.CareSeeker) error {
	return nil
}

func NewCareSeeker(querier queries.Querier) *CareSeekerProvider {
	return &CareSeekerProvider{}
}
