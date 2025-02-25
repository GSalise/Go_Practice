package graph

import (
	"github.com/GSalise/GQLPractice/graph/model"
	"github.com/jackc/pgx/v5"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CharacterStore map[string]model.Character
	DB             *pgx.Conn
}
