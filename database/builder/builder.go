package builder

import (
	"github.com/jacksoom/juno/v4/database"

	"github.com/jacksoom/juno/v4/database/postgresql"
)

// Builder represents a generic Builder implementation that build the proper database
// instance based on the configuration the user has specified
func Builder(ctx *database.Context) (database.Database, error) {
	return postgresql.Builder(ctx)
}
