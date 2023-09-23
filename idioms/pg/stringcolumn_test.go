package pg_test

import (
	"testing"

	"github.com/e200/migro/idioms/pg"
	"github.com/stretchr/testify/assert"
)

func TestNewStringColumnToSQL(t *testing.T) {
	statement := pg.NewStringColumn("name", 255)

	query := statement.Nullable().Unique().Default("1").ToSQL()
	assert.NotEmpty(t, query)
	assert.Equal(t, "name VARCHAR(255) DEFAULT '1' UNIQUE", query)

	statement = pg.NewStringColumn("name", 255)
	query = statement.ToSQL()
	assert.NotEmpty(t, query)
	assert.Equal(t, "name VARCHAR(255) NOT NULL", query)

	statement = pg.NewStringColumn("name", 255)

	query = statement.Unique().ToSQL()
	assert.NotEmpty(t, query)
	assert.Equal(t, "name VARCHAR(255) UNIQUE NOT NULL", query)

	statement = pg.NewStringColumn("name", 255)

	query = statement.Default("CURRENT_TIMESTAMP").Nullable().ToSQL()
	assert.NotEmpty(t, query)
	assert.Equal(t, "name VARCHAR(255) DEFAULT 'CURRENT_TIMESTAMP'", query)

	statement = pg.NewStringColumn("name", 255)

	query = statement.Unique().ToSQL()
	assert.NotEmpty(t, query)
	assert.Equal(t, "name VARCHAR(255) UNIQUE NOT NULL", query)
}

func TestAddStringColumnToSQL(t *testing.T) {
	statement := pg.NewAddStringColumn("name", 255)

	query := statement.Nullable().Unique().Default("1").ToSQL()
	assert.NotEmpty(t, query)
	assert.Equal(t, "ADD COLUMN name VARCHAR(255) DEFAULT '1' UNIQUE", query)

	statement = pg.NewAddStringColumn("name", 255)
	query = statement.ToSQL()
	assert.NotEmpty(t, query)
	assert.Equal(t, "ADD COLUMN name VARCHAR(255) NOT NULL", query)

	statement = pg.NewAddStringColumn("name", 255)

	query = statement.Unique().ToSQL()
	assert.NotEmpty(t, query)
	assert.Equal(t, "ADD COLUMN name VARCHAR(255) UNIQUE NOT NULL", query)

	statement = pg.NewAddStringColumn("name", 255)

	query = statement.Default("CURRENT_TIMESTAMP").Nullable().ToSQL()
	assert.NotEmpty(t, query)
	assert.Equal(t, "ADD COLUMN name VARCHAR(255) DEFAULT 'CURRENT_TIMESTAMP'", query)

	statement = pg.NewAddStringColumn("name", 255)

	query = statement.Unique().ToSQL()
	assert.NotEmpty(t, query)
	assert.Equal(t, "ADD COLUMN name VARCHAR(255) UNIQUE NOT NULL", query)
}
