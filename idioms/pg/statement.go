package pg

import (
	"strings"
)

type statements struct {
	statements []IQuery
}

func (s *statements) Add(statement IQuery) *statements {
	s.statements = append(s.statements, statement)

	return s
}

func (s *statements) ToSQL() string {
	parts := make([]string, 0, len(s.statements))

	for _, statement := range s.statements {
		query := statement.ToSQL()

		parts = append(parts, query)
	}

	return strings.Join(parts, ",")
}

var _ IQuery = (*DropColumn)(nil)
