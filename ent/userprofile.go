// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/nibbleshift/drift/ent/userprofile"
)

// UserProfile is the model entity for the UserProfile schema.
type UserProfile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// About holds the value of the "about" field.
	About string `json:"about,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// Dob holds the value of the "dob" field.
	Dob time.Time `json:"dob,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserProfileQuery when eager-loading is set.
	Edges        UserProfileEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserProfileEdges holds the relations/edges for other nodes in the graph.
type UserProfileEdges struct {
	// Links holds the value of the links edge.
	Links []*Link `json:"links,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedLinks map[string][]*Link
}

// LinksOrErr returns the Links value or an error if the edge
// was not loaded in eager-loading.
func (e UserProfileEdges) LinksOrErr() ([]*Link, error) {
	if e.loadedTypes[0] {
		return e.Links, nil
	}
	return nil, &NotLoadedError{edge: "links"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserProfile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userprofile.FieldID:
			values[i] = new(sql.NullInt64)
		case userprofile.FieldAvatar, userprofile.FieldAbout, userprofile.FieldLocation:
			values[i] = new(sql.NullString)
		case userprofile.FieldDob:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserProfile fields.
func (up *UserProfile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userprofile.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			up.ID = int(value.Int64)
		case userprofile.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				up.Avatar = value.String
			}
		case userprofile.FieldAbout:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field about", values[i])
			} else if value.Valid {
				up.About = value.String
			}
		case userprofile.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				up.Location = value.String
			}
		case userprofile.FieldDob:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dob", values[i])
			} else if value.Valid {
				up.Dob = value.Time
			}
		default:
			up.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserProfile.
// This includes values selected through modifiers, order, etc.
func (up *UserProfile) Value(name string) (ent.Value, error) {
	return up.selectValues.Get(name)
}

// QueryLinks queries the "links" edge of the UserProfile entity.
func (up *UserProfile) QueryLinks() *LinkQuery {
	return NewUserProfileClient(up.config).QueryLinks(up)
}

// Update returns a builder for updating this UserProfile.
// Note that you need to call UserProfile.Unwrap() before calling this method if this UserProfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (up *UserProfile) Update() *UserProfileUpdateOne {
	return NewUserProfileClient(up.config).UpdateOne(up)
}

// Unwrap unwraps the UserProfile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (up *UserProfile) Unwrap() *UserProfile {
	_tx, ok := up.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserProfile is not a transactional entity")
	}
	up.config.driver = _tx.drv
	return up
}

// String implements the fmt.Stringer.
func (up *UserProfile) String() string {
	var builder strings.Builder
	builder.WriteString("UserProfile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", up.ID))
	builder.WriteString("avatar=")
	builder.WriteString(up.Avatar)
	builder.WriteString(", ")
	builder.WriteString("about=")
	builder.WriteString(up.About)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(up.Location)
	builder.WriteString(", ")
	builder.WriteString("dob=")
	builder.WriteString(up.Dob.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedLinks returns the Links named value or an error if the edge was not
// loaded in eager-loading with this name.
func (up *UserProfile) NamedLinks(name string) ([]*Link, error) {
	if up.Edges.namedLinks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := up.Edges.namedLinks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (up *UserProfile) appendNamedLinks(name string, edges ...*Link) {
	if up.Edges.namedLinks == nil {
		up.Edges.namedLinks = make(map[string][]*Link)
	}
	if len(edges) == 0 {
		up.Edges.namedLinks[name] = []*Link{}
	} else {
		up.Edges.namedLinks[name] = append(up.Edges.namedLinks[name], edges...)
	}
}

// UserProfiles is a parsable slice of UserProfile.
type UserProfiles []*UserProfile
