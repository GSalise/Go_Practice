package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/jackc/pgx/v5"
)

type Resolver struct {
	DB *pgx.Conn // Database connection
}
