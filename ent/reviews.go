// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ZEQUANR/zhulong/ent/reviews"
	"github.com/ZEQUANR/zhulong/ent/user"
)

// Reviews is the model entity for the Reviews schema.
type Reviews struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FileName holds the value of the "file_name" field.
	FileName string `json:"file_name,omitempty"`
	// FileURL holds the value of the "file_url" field.
	FileURL string `json:"file_url,omitempty"`
	// UploadTime holds the value of the "upload_time" field.
	UploadTime time.Time `json:"upload_time,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// ReviewsTitle holds the value of the "reviews_title" field.
	ReviewsTitle string `json:"reviews_title,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReviewsQuery when eager-loading is set.
	Edges        ReviewsEdges `json:"edges"`
	user_reviews *int
	selectValues sql.SelectValues
}

// ReviewsEdges holds the relations/edges for other nodes in the graph.
type ReviewsEdges struct {
	// Uploaders holds the value of the uploaders edge.
	Uploaders *User `json:"uploaders,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UploadersOrErr returns the Uploaders value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReviewsEdges) UploadersOrErr() (*User, error) {
	if e.Uploaders != nil {
		return e.Uploaders, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "uploaders"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Reviews) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reviews.FieldID:
			values[i] = new(sql.NullInt64)
		case reviews.FieldFileName, reviews.FieldFileURL, reviews.FieldReviewsTitle:
			values[i] = new(sql.NullString)
		case reviews.FieldUploadTime, reviews.FieldCreateTime:
			values[i] = new(sql.NullTime)
		case reviews.ForeignKeys[0]: // user_reviews
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Reviews fields.
func (r *Reviews) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reviews.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case reviews.FieldFileName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_name", values[i])
			} else if value.Valid {
				r.FileName = value.String
			}
		case reviews.FieldFileURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_url", values[i])
			} else if value.Valid {
				r.FileURL = value.String
			}
		case reviews.FieldUploadTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field upload_time", values[i])
			} else if value.Valid {
				r.UploadTime = value.Time
			}
		case reviews.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				r.CreateTime = value.Time
			}
		case reviews.FieldReviewsTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reviews_title", values[i])
			} else if value.Valid {
				r.ReviewsTitle = value.String
			}
		case reviews.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_reviews", value)
			} else if value.Valid {
				r.user_reviews = new(int)
				*r.user_reviews = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Reviews.
// This includes values selected through modifiers, order, etc.
func (r *Reviews) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryUploaders queries the "uploaders" edge of the Reviews entity.
func (r *Reviews) QueryUploaders() *UserQuery {
	return NewReviewsClient(r.config).QueryUploaders(r)
}

// Update returns a builder for updating this Reviews.
// Note that you need to call Reviews.Unwrap() before calling this method if this Reviews
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Reviews) Update() *ReviewsUpdateOne {
	return NewReviewsClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Reviews entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Reviews) Unwrap() *Reviews {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Reviews is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Reviews) String() string {
	var builder strings.Builder
	builder.WriteString("Reviews(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("file_name=")
	builder.WriteString(r.FileName)
	builder.WriteString(", ")
	builder.WriteString("file_url=")
	builder.WriteString(r.FileURL)
	builder.WriteString(", ")
	builder.WriteString("upload_time=")
	builder.WriteString(r.UploadTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(r.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("reviews_title=")
	builder.WriteString(r.ReviewsTitle)
	builder.WriteByte(')')
	return builder.String()
}

// ReviewsSlice is a parsable slice of Reviews.
type ReviewsSlice []*Reviews
