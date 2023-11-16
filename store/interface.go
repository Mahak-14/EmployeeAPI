package store

import (
	"sample/model"

	"gofr.dev/pkg/gofr"
)

type Employee interface {
	// Create inserts a new employee record into the database
	Create(ctx *gofr.Context, student *model.Employee) (*model.Employee, error)

	// GetByID retrieves a employee record based on its ID
	GetByID(ctx *gofr.Context, id int) (*model.Employee, error)

	// Update updates an existing employee record with the provided information
	Update(ctx *gofr.Context, student *model.Employee) (*model.Employee, error)

	// Delete removes a employee record from the database based on its ID
	Delete(ctx *gofr.Context, id int) error
}
