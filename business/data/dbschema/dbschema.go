// Package dbschema contains the database schema.
package dbschema

import (
	_ "embed" // Calls init function.
)

//go:embed schema.sql
var schemaDoc string

// TODO create table from the code.
