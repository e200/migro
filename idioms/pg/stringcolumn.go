package pg

import (
	"fmt"
)

type stringColumn struct{ columnAttributes }
type addStringColumn struct{ columnAttributes }

func NewStringColumn(columnName string, length int) *stringColumn {
	return &stringColumn{
		columnAttributes: columnAttributes{
			columnName:   columnName,
			columnLength: length,
		},
	}
}

func NewAddStringColumn(columnName string, length int) *addStringColumn {
	return &addStringColumn{
		columnAttributes: columnAttributes{
			columnName:   columnName,
			columnLength: length,
		},
	}
}

func (s *stringColumn) SetLength(length int) *stringColumn {
	s.columnAttributes.SetLength(length)

	return s
}

func (s *stringColumn) Nullable() *stringColumn {
	s.columnAttributes.Nullable()

	return s
}

func (s *stringColumn) Unique() *stringColumn {
	s.columnAttributes.Unique()

	return s
}

func (s *stringColumn) NotNullable() *stringColumn {
	s.columnAttributes.NotNullable()

	return s
}

func (s *stringColumn) Default(defaultValue string) *stringColumn {
	s.columnAttributes.Default(defaultValue)

	return s
}

func (a *stringColumn) ToSQL() string {
	return fmt.Sprintf("%s VARCHAR(%d) %s", a.columnName, a.columnLength, a.columnAttributes.GetAttributes())
}

func (a *addStringColumn) SetLength(length int) *addStringColumn {
	a.columnAttributes.SetLength(length)

	return a
}

func (a *addStringColumn) Nullable() *addStringColumn {
	a.columnAttributes.Nullable()

	return a
}

func (a *addStringColumn) Unique() *addStringColumn {
	a.columnAttributes.Unique()

	return a
}

func (a *addStringColumn) NotNullable() *addStringColumn {
	a.columnAttributes.NotNullable()

	return a
}

func (a *addStringColumn) Default(defaultValue string) *addStringColumn {
	a.columnAttributes.Default(defaultValue)

	return a
}

func (a *addStringColumn) ToSQL() string {
	return fmt.Sprintf("ADD COLUMN %s VARCHAR(%d) %s", a.columnName, a.columnLength, a.columnAttributes.GetAttributes())
}

var _ IQuery = (*stringColumn)(nil)
var _ IQuery = (*addStringColumn)(nil)
