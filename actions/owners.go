package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kteb/pet_owner/models"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Owner)
// DB Table: Plural (owners)
// Resource: Plural (Owners)
// Path: Plural (/owners)
// View Template Folder: Plural (/templates/owners/)

// OwnersResource is the resource for the Owner model
type OwnersResource struct {
	buffalo.Resource
}

// List gets all Owners. This function is mapped to the path
// GET /owners
func (v OwnersResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	owners := &models.Owners{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Owners from the DB
	if err := q.All(owners); err != nil {
		return errors.WithStack(err)
	}

	// Make Owners available inside the html template
	c.Set("owners", owners)

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("owners/index.html"))
}

// Show gets the data for one Owner. This function is mapped to
// the path GET /owners/{owner_id}
func (v OwnersResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Owner
	owner := &models.Owner{}

	// To find the Owner the parameter owner_id is used.
	if err := tx.Find(owner, c.Param("owner_id")); err != nil {
		return c.Error(404, err)
	}

	// Make owner available inside the html template
	c.Set("owner", owner)

	return c.Render(200, r.HTML("owners/show.html"))
}

// New renders the form for creating a new Owner.
// This function is mapped to the path GET /owners/new
func (v OwnersResource) New(c buffalo.Context) error {
	// Make owner available inside the html template
	c.Set("owner", &models.Owner{})

	return c.Render(200, r.HTML("owners/new.html"))
}

// Create adds a Owner to the DB. This function is mapped to the
// path POST /owners
func (v OwnersResource) Create(c buffalo.Context) error {
	// Allocate an empty Owner
	owner := &models.Owner{}

	// Bind owner to the html form elements
	if err := c.Bind(owner); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(owner)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make owner available inside the html template
		c.Set("owner", owner)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("owners/new.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Owner was created successfully")

	// and redirect to the owners index page
	return c.Redirect(302, "/owners/%s", owner.ID)
}

// Edit renders a edit form for a Owner. This function is
// mapped to the path GET /owners/{owner_id}/edit
func (v OwnersResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Owner
	owner := &models.Owner{}

	if err := tx.Find(owner, c.Param("owner_id")); err != nil {
		return c.Error(404, err)
	}

	// Make owner available inside the html template
	c.Set("owner", owner)
	return c.Render(200, r.HTML("owners/edit.html"))
}

// Update changes a Owner in the DB. This function is mapped to
// the path PUT /owners/{owner_id}
func (v OwnersResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Owner
	owner := &models.Owner{}

	if err := tx.Find(owner, c.Param("owner_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Owner to the html form elements
	if err := c.Bind(owner); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(owner)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make owner available inside the html template
		c.Set("owner", owner)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("owners/edit.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Owner was updated successfully")

	// and redirect to the owners index page
	return c.Redirect(302, "/owners/%s", owner.ID)
}

// Destroy deletes a Owner from the DB. This function is mapped
// to the path DELETE /owners/{owner_id}
func (v OwnersResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Owner
	owner := &models.Owner{}

	// To find the Owner the parameter owner_id is used.
	if err := tx.Find(owner, c.Param("owner_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(owner); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Owner was destroyed successfully")

	// Redirect to the owners index page
	return c.Redirect(302, "/owners")
}
