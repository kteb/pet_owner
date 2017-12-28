package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kteb/pet_owner/models"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

type Selectable struct {
	Value interface{}
	Label string
}

func (s Selectable) SelectValue() interface{} {
	return s.Value
}

func (s Selectable) SelectLabel() string {
	return s.Label
}

func selectOwners(c buffalo.Context) ([]Selectable, error) {
	// TODO:  implement generic selectable

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	owners := &models.Owners{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.Q()

	// Retrieve all Federations from the DB
	if err := q.All(owners); err != nil {
		return nil, errors.WithStack(err)
	}

	result := make([]Selectable, len(*owners))

	for i, element := range *owners {
		result[i].Value = element.ID
		result[i].Label = element.Name
	}
	return result, nil
}
