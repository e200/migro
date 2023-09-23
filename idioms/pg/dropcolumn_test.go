package pg_test

import (
	"testing"

	"github.com/e200/migro/idioms/pg"
	"github.com/stretchr/testify/assert"
)

func TestDropColumnToSQL(t *testing.T) {
	statement := pg.NewDropColumn("name")

	assert.Equal(t, "DROP COLUMN name", statement.ToSQL())
}
