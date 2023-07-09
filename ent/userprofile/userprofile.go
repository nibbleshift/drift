// Code generated by ent, DO NOT EDIT.

package userprofile

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the userprofile type in the database.
	Label = "user_profile"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldAbout holds the string denoting the about field in the database.
	FieldAbout = "about"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldDob holds the string denoting the dob field in the database.
	FieldDob = "dob"
	// EdgeLinks holds the string denoting the links edge name in mutations.
	EdgeLinks = "links"
	// Table holds the table name of the userprofile in the database.
	Table = "user_profiles"
	// LinksTable is the table that holds the links relation/edge.
	LinksTable = "links"
	// LinksInverseTable is the table name for the Link entity.
	// It exists in this package in order to avoid circular dependency with the "link" package.
	LinksInverseTable = "links"
	// LinksColumn is the table column denoting the links relation/edge.
	LinksColumn = "user_profile_links"
)

// Columns holds all SQL columns for userprofile fields.
var Columns = []string{
	FieldID,
	FieldAvatar,
	FieldAbout,
	FieldLocation,
	FieldDob,
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

// OrderOption defines the ordering options for the UserProfile queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByAbout orders the results by the about field.
func ByAbout(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAbout, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByDob orders the results by the dob field.
func ByDob(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDob, opts...).ToFunc()
}

// ByLinksCount orders the results by links count.
func ByLinksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLinksStep(), opts...)
	}
}

// ByLinks orders the results by links terms.
func ByLinks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLinksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newLinksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LinksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LinksTable, LinksColumn),
	)
}
