// Code generated by entc, DO NOT EDIT.

package usersecret

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the usersecret type in the database.
	Label = "user_secret"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the usersecret in the database.
	Table = "user_secrets"
)

// Columns holds all SQL columns for usersecret fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldAppID,
	FieldSecret,
	FieldCreateAt,
	FieldDeleteAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
