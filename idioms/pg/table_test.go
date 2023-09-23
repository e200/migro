package pg_test

import (
	"testing"

	"github.com/e200/migro/idioms/pg"
	"github.com/stretchr/testify/assert"
)

func TestTableToSQL(t *testing.T) {
	createQuery := pg.NewCreateTable("users")
	createQuery.String("name").Unique().Default("1").Nullable()
	assert.Equal(t, "CREATE TABLE users (name VARCHAR(255) DEFAULT '1' UNIQUE);", createQuery.ToSQL())

	alterQuery := pg.NewAlterTable("users")
	alterQuery.String("name")
	assert.Equal(t, "ALTER TABLE users ADD COLUMN name VARCHAR(255) NOT NULL;", alterQuery.ToSQL())

	alterQuery = pg.NewAlterTable("orders")
	alterQuery.String("order_number").Unique()
	assert.Equal(t, "ALTER TABLE orders ADD COLUMN order_number VARCHAR(255) UNIQUE NOT NULL;", alterQuery.ToSQL())

	dropQuery := pg.NewDropTable("orders")
	assert.Equal(t, "DROP TABLE orders;", dropQuery.ToSQL())
}
