package v3

import (
	"github.com/jmoiron/sqlx"

	"github.com/jacksoom/juno/v4/database"
	"github.com/jacksoom/juno/v4/database/postgresql"
)

var _ database.Migrator = &Migrator{}

// Migrator represents the database migrator that should be used to migrate from v2 of the database to v3
type Migrator struct {
	SQL *sqlx.DB
}

func NewMigrator(db *postgresql.Database) *Migrator {
	return &Migrator{
		SQL: db.SQL,
	}
}
