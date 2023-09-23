package pg

import (
	"fmt"
)

type DropColumn struct {
	columnName string
}

func NewDropColumn(columnName string) *DropColumn {
	return &DropColumn{
		columnName: columnName,
	}
}

func (d *DropColumn) ToSQL() string {
	return fmt.Sprintf("DROP COLUMN %s", d.columnName)
}

var _ IQuery = (*DropColumn)(nil)
