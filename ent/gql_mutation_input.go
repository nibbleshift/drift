// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Username  string
	FirstName string
	LastName  string
	UtterIDs  []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetUsername(i.Username)
	m.SetFirstName(i.FirstName)
	m.SetLastName(i.LastName)
	if v := i.UtterIDs; len(v) > 0 {
		m.AddUtterIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateUtterInput represents a mutation input for creating utters.
type CreateUtterInput struct {
	CreatedAt *time.Time
	Data      string
	OwnerID   *int
}

// Mutate applies the CreateUtterInput on the UtterMutation builder.
func (i *CreateUtterInput) Mutate(m *UtterMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	m.SetData(i.Data)
	if v := i.OwnerID; v != nil {
		m.SetOwnerID(*v)
	}
}

// SetInput applies the change-set in the CreateUtterInput on the UtterCreate builder.
func (c *UtterCreate) SetInput(i CreateUtterInput) *UtterCreate {
	i.Mutate(c.Mutation())
	return c
}