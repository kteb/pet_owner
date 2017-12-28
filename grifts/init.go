package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kteb/pet_owner/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
