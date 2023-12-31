// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateLinkInput represents a mutation input for creating links.
type CreateLinkInput struct {
	CreatedAt *time.Time
	Data      string
}

// Mutate applies the CreateLinkInput on the LinkMutation builder.
func (i *CreateLinkInput) Mutate(m *LinkMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	m.SetData(i.Data)
}

// SetInput applies the change-set in the CreateLinkInput on the LinkCreate builder.
func (c *LinkCreate) SetInput(i CreateLinkInput) *LinkCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreatePostInput represents a mutation input for creating posts.
type CreatePostInput struct {
	CreatedAt  *time.Time
	Data       string
	OwnerID    *int
	TagIDs     []int
	MentionIDs []int
}

// Mutate applies the CreatePostInput on the PostMutation builder.
func (i *CreatePostInput) Mutate(m *PostMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	m.SetData(i.Data)
	if v := i.OwnerID; v != nil {
		m.SetOwnerID(*v)
	}
	if v := i.TagIDs; len(v) > 0 {
		m.AddTagIDs(v...)
	}
	if v := i.MentionIDs; len(v) > 0 {
		m.AddMentionIDs(v...)
	}
}

// SetInput applies the change-set in the CreatePostInput on the PostCreate builder.
func (c *PostCreate) SetInput(i CreatePostInput) *PostCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateTagInput represents a mutation input for creating tags.
type CreateTagInput struct {
	Data    string
	PostIDs []int
}

// Mutate applies the CreateTagInput on the TagMutation builder.
func (i *CreateTagInput) Mutate(m *TagMutation) {
	m.SetData(i.Data)
	if v := i.PostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
}

// SetInput applies the change-set in the CreateTagInput on the TagCreate builder.
func (c *TagCreate) SetInput(i CreateTagInput) *TagCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Username    string
	FirstName   string
	LastName    string
	PostIDs     []int
	FriendIDs   []int
	FollowerIDs []int
	ProfileID   *int
	MentionIDs  []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetUsername(i.Username)
	m.SetFirstName(i.FirstName)
	m.SetLastName(i.LastName)
	if v := i.PostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.FriendIDs; len(v) > 0 {
		m.AddFriendIDs(v...)
	}
	if v := i.FollowerIDs; len(v) > 0 {
		m.AddFollowerIDs(v...)
	}
	if v := i.ProfileID; v != nil {
		m.SetProfileID(*v)
	}
	if v := i.MentionIDs; len(v) > 0 {
		m.AddMentionIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserProfileInput represents a mutation input for creating userprofiles.
type CreateUserProfileInput struct {
	Avatar   string
	About    string
	Location string
	Dob      time.Time
	LinkIDs  []int
}

// Mutate applies the CreateUserProfileInput on the UserProfileMutation builder.
func (i *CreateUserProfileInput) Mutate(m *UserProfileMutation) {
	m.SetAvatar(i.Avatar)
	m.SetAbout(i.About)
	m.SetLocation(i.Location)
	m.SetDob(i.Dob)
	if v := i.LinkIDs; len(v) > 0 {
		m.AddLinkIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserProfileInput on the UserProfileCreate builder.
func (c *UserProfileCreate) SetInput(i CreateUserProfileInput) *UserProfileCreate {
	i.Mutate(c.Mutation())
	return c
}
