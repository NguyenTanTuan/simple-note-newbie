//go:generate go run github.com/99designs/gqlgen
package resolvers

import (
	"simple-note-newbie/graph/model"
	"simple-note-newbie/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	notes []*model.Note
	subNotes []*model.SubNote
	services.INoteService
}
