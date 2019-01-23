package user

import (
	"context"

	"../models"
)

// Repository represent the author's repository contract
type Repository interface {
	GetByID(ctx context.Context, id int64) (*models.User, error)
}