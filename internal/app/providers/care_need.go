package providers

import "github.com/homayoonalimohammadi/care-companion/internal/app/db/models"

type CareNeedProvider interface {
	Get(id string) (models.CareNeed, error)
	Create(models.CareNeed) error
}
