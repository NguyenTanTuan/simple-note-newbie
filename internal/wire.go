//+build wireinject
package internal

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/wire"
	"simple-note-newbie/services"
)

//type Cron struct {
//}

type App struct{
	Resolver *graphql.Resolver
}

var servicesSet = wire.NewSet(
	services.NewNoteService,
)


func GetApp() (*App, func(), error) {
	wire.Build(
		servicesSet,
		wire.Struct(new(graphql.Resolver), "*"),
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil, nil
}

//func GetTestApp() (*Cron, error) {
//	wire.Build(servicesSet)
//	return &Cron{}, nil
//}