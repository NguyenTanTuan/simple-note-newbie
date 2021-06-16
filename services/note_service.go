package services

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"math/rand"
	"simple-note-newbie/graph/model"
)

type INoteService interface {
	AddNote(ctx context.Context, text string, description string) (*model.Note, error)
	GetAllNotes(ctx context.Context) ([]*model.Note, error)
	AddSubNote(ctx context.Context, text string, description string, IDParent string) (*model.SubNote, error)
}

// define note service that implement INoteService
type NoteService struct {
	notes []*model.Note
}

// define provider
//var NewNoteService = wire.NewSet(Init, wire.Bind(new(INoteService), new(NoteService)))

var NewNoteService = wire.NewSet(wire.Struct(new(NoteService),"*"), wire.Bind(new(INoteService), new(*NoteService)))

func Init()  *NoteService{
	return &NoteService{}
}

func (vs *NoteService) AddNote(ctx context.Context, text string, description string) (*model.Note, error) {
	newNote := &model.Note{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Text:        text,
		Description: description,
	}

	vs.notes = append(vs.notes, newNote)
	return newNote, nil
}

func (vs *NoteService) GetAllNotes(ctx context.Context) ([]*model.Note, error) {
	return vs.notes, nil
}

func (vs *NoteService) AddSubNote(ctx context.Context, text string, description string, IDParent string) (*model.SubNote, error) {
	var note *model.Note
	for _, n := range vs.notes {
		if n.ID == IDParent {
			note = n
			break
		}
	}

	if note == nil {
		panic("Parent note not found")
	}

	newSubNote := &model.SubNote{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Text:        text,
		Description: description,
	}

	note.SubNote = append(note.SubNote, newSubNote)
	return newSubNote, nil
}

