package pg

import (
	"fmt"
	"strings"
)

type columnAttributes struct {
	columnName   string
	columnLength int
	nullable     bool
	unique       bool
	defaultValue *string
}

func (c *columnAttributes) SetLength(length int) {
	c.columnLength = length
}

func (c *columnAttributes) Nullable() {
	c.nullable = true
}

func (c *columnAttributes) Unique() {
	c.unique = true
}

func (c *columnAttributes) NotNullable() {
	c.nullable = false
}

func (c *columnAttributes) Default(defaultValue string) {
	c.defaultValue = &defaultValue
}

func (c *columnAttributes) GetAttributes() string {
	var parts []string

	if c.defaultValue != nil {
		parts = append(parts, fmt.Sprintf("DEFAULT '%s'", *c.defaultValue))
	}

	if c.unique {
		parts = append(parts, "UNIQUE")
	}

	if !c.nullable {
		parts = append(parts, "NOT NULL")
	}

	return strings.Join(parts, " ")
}
