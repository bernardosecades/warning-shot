package warning

import (
	"context"

	"../models"
)

// Usecase represent the warning's usecases
type Usecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*models.Warning, string, error)
	GetByID(ctx context.Context, id int64) (*models.Warning, error)
	Update(ctx context.Context, m *models.Warning) error
	Store(context.Context, *models.Warning) error
	Delete(ctx context.Context, id int64) error
}