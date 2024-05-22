// Code generated by ent, DO NOT EDIT.

package reviews

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the reviews type in the database.
	Label = "reviews"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFileName holds the string denoting the file_name field in the database.
	FieldFileName = "file_name"
	// FieldFileURL holds the string denoting the file_url field in the database.
	FieldFileURL = "file_url"
	// FieldUploadTime holds the string denoting the upload_time field in the database.
	FieldUploadTime = "upload_time"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldReviewsTitle holds the string denoting the reviews_title field in the database.
	FieldReviewsTitle = "reviews_title"
	// EdgeUploaders holds the string denoting the uploaders edge name in mutations.
	EdgeUploaders = "uploaders"
	// Table holds the table name of the reviews in the database.
	Table = "reviews"
	// UploadersTable is the table that holds the uploaders relation/edge.
	UploadersTable = "reviews"
	// UploadersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UploadersInverseTable = "users"
	// UploadersColumn is the table column denoting the uploaders relation/edge.
	UploadersColumn = "user_reviews"
)

// Columns holds all SQL columns for reviews fields.
var Columns = []string{
	FieldID,
	FieldFileName,
	FieldFileURL,
	FieldUploadTime,
	FieldCreateTime,
	FieldReviewsTitle,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "reviews"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_reviews",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
)

// OrderOption defines the ordering options for the Reviews queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFileName orders the results by the file_name field.
func ByFileName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileName, opts...).ToFunc()
}

// ByFileURL orders the results by the file_url field.
func ByFileURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileURL, opts...).ToFunc()
}

// ByUploadTime orders the results by the upload_time field.
func ByUploadTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUploadTime, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByReviewsTitle orders the results by the reviews_title field.
func ByReviewsTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReviewsTitle, opts...).ToFunc()
}

// ByUploadersField orders the results by uploaders field.
func ByUploadersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUploadersStep(), sql.OrderByField(field, opts...))
	}
}
func newUploadersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UploadersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UploadersTable, UploadersColumn),
	)
}
