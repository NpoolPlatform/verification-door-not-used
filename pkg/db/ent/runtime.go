// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/verification-door/pkg/db/ent/schema"
	"github.com/NpoolPlatform/verification-door/pkg/db/ent/usersecret"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	usersecretFields := schema.UserSecret{}.Fields()
	_ = usersecretFields
	// usersecretDescCreateAt is the schema descriptor for create_at field.
	usersecretDescCreateAt := usersecretFields[3].Descriptor()
	// usersecret.DefaultCreateAt holds the default value on creation for the create_at field.
	usersecret.DefaultCreateAt = usersecretDescCreateAt.Default.(func() uint32)
	// usersecretDescDeleteAt is the schema descriptor for delete_at field.
	usersecretDescDeleteAt := usersecretFields[4].Descriptor()
	// usersecret.DefaultDeleteAt holds the default value on creation for the delete_at field.
	usersecret.DefaultDeleteAt = usersecretDescDeleteAt.Default.(func() uint32)
	// usersecretDescID is the schema descriptor for id field.
	usersecretDescID := usersecretFields[0].Descriptor()
	// usersecret.DefaultID holds the default value on creation for the id field.
	usersecret.DefaultID = usersecretDescID.Default.(func() uuid.UUID)
}
