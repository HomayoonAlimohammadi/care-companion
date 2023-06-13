package providers

import "github.com/homayoonalimohammadi/care-companion/internal/app/db/models"

type CareSeekerProvider interface {
	Get(id string) (models.CareSeeker, error)
	Create(models.CareSeeker) error
}
