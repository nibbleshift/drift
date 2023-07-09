// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/nibbleshift/drift/ent/schema"
	"github.com/nibbleshift/drift/ent/utter"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	utterFields := schema.Utter{}.Fields()
	_ = utterFields
	// utterDescCreatedAt is the schema descriptor for created_at field.
	utterDescCreatedAt := utterFields[0].Descriptor()
	// utter.DefaultCreatedAt holds the default value on creation for the created_at field.
	utter.DefaultCreatedAt = utterDescCreatedAt.Default.(func() time.Time)
}