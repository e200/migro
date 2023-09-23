package pg

import (
	"fmt"
)

type createTable struct {
	tableName string
	statements
}

type alterTable struct {
	tableName string
	statements
}

type dropTable struct {
	tableName string
}

func NewCreateTable(tableName string) *createTable {
	return &createTable{
		tableName: tableName,
	}
}

func NewAlterTable(tableName string) *alterTable {
	return &alterTable{
		tableName: tableName,
	}
}

func NewDropTable(tableName string) *dropTable {
	return &dropTable{
		tableName: tableName,
	}
}

func (c *createTable) String(columnName string) *stringColumn {
	stringColumnStatement := NewStringColumn(columnName, 255)

	c.Add(stringColumnStatement)

	return stringColumnStatement
}

func (c *createTable) StringN(columnName string, length int) *createTable {
	c.Add(NewStringColumn(columnName, length))

	return c
}

func (c *createTable) Integer(columnName string) *createTable {
	return nil
}

func (c *createTable) Date(columnName string) *createTable {
	return nil
}

func (c *createTable) Boolean(columnName string) *createTable {
	return nil
}

func (c *createTable) UUID(columnName string) *createTable {
	return nil
}

func (c *createTable) BigInt(columnName string) *createTable {
	return nil
}

func (c *createTable) TinyInt(columnName string) *createTable {
	return nil
}

func (c *createTable) DropColumn(columnName string) {
	c.Add(NewDropColumn(columnName))
}

func (c *createTable) ToSQL() string {
	return fmt.Sprintf(
		"CREATE TABLE %s (%s);",
		c.tableName,
		c.statements.ToSQL(),
	)
}

func (a *alterTable) String(columnName string) *addStringColumn {
	statement := NewAddStringColumn(columnName, 255)

	a.Add(statement)

	return statement
}

func (a *alterTable) StringN(columnName string, length int) *alterTable {
	a.Add(NewAddStringColumn(columnName, length))

	return a
}

func (a *alterTable) Integer(columnName string) *alterTable {
	return nil
}

func (a *alterTable) Date(columnName string) *alterTable {
	return nil
}

func (a *alterTable) Boolean(columnName string) *alterTable {
	return nil
}

func (a *alterTable) UUID(columnName string) *alterTable {
	return nil
}

func (a *alterTable) BigInt(columnName string) *alterTable {
	return nil
}

func (a *alterTable) TinyInt(columnName string) *alterTable {
	return nil
}

func (a *alterTable) DropColumn(columnName string) {
	a.Add(NewDropColumn(columnName))
}

func (a *alterTable) ToSQL() string {
	return fmt.Sprintf(
		"ALTER TABLE %s %s;",
		a.tableName,
		a.statements.ToSQL(),
	)
}

func (a *dropTable) ToSQL() string {
	return fmt.Sprintf("DROP TABLE %s;", a.tableName)
}

var _ IQuery = (*createTable)(nil)
var _ IQuery = (*alterTable)(nil)
var _ IQuery = (*dropTable)(nil)
