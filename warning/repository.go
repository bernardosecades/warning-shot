package warning

import (
	"context"

	"../models"
)

// Repository represent the warning's repository contract
type Repository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []*models.Warning, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (*models.Warning, error)
	Update(ctx context.Context, ar *models.Warning) error
	Store(ctx context.Context, a *models.Warning) error
	Delete(ctx context.Context, id int64) error
}
