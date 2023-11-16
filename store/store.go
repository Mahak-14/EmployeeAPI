package store

import (
	"database/sql"
	"fmt"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"sample/model"
)

type employee struct{}

func New() *employee {
	return &employee{}
}

// Create inserts a new employee record into the database
func (s *employee) Create(ctx *gofr.Context, emp *model.Employee) (*model.Employee, error) {
	_, err := ctx.DB().ExecContext(ctx, createQuery, emp.ID, emp.Name, emp.Dept)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return emp, nil
}

// GetByID retrieves a employee record based on its ID
func (s *employee) GetByID(ctx *gofr.Context, id int) (*model.Employee, error) {
	var resp model.Employee

	err := ctx.DB().QueryRowContext(ctx, getByIDQuery, id).
		Scan(&resp.ID, &resp.Name, &resp.Dept)
	switch err {
	case sql.ErrNoRows:
		return nil, errors.EntityNotFound{Entity: "employee", ID: fmt.Sprintf("%v", id)}
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// Update updates an existing employee record with the provided information
func (s *employee) Update(ctx *gofr.Context, emp *model.Employee) (*model.Employee, error) {
	_, err := ctx.DB().ExecContext(ctx, updateQuery, emp.Name, emp.Dept, emp.ID)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return emp, nil
}

// Delete removes a employee record from the database based on its ID
func (s *employee) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, deleteQuery, id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
