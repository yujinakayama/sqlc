// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package querytest_strict

import (
	"database/sql"
)

type Foo struct {
	Bar      string
	MaybeBar sql.NullString
}