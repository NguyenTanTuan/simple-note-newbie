#!/usr/bin/env bash



genGql() {
    cd graph/schema
    go run simple-note-newbie/cmd/gqlgen
}

wire() {
    go get github.com/google/wire/cmd/wire
    go run github.com/google/wire/cmd/wire simple-note-newbie/internal
}