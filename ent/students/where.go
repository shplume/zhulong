// Code generated by ent, DO NOT EDIT.

package students

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ZEQUANR/zhulong/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Students {
	return predicate.Students(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Students {
	return predicate.Students(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Students {
	return predicate.Students(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Students {
	return predicate.Students(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Students {
	return predicate.Students(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Students {
	return predicate.Students(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Students {
	return predicate.Students(sql.FieldLTE(FieldID, id))
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldAvatar, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldName, v))
}

// College applies equality check predicate on the "college" field. It's identical to CollegeEQ.
func College(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldCollege, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldPhone, v))
}

// Major applies equality check predicate on the "major" field. It's identical to MajorEQ.
func Major(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldMajor, v))
}

// Class applies equality check predicate on the "class" field. It's identical to ClassEQ.
func Class(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldClass, v))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldNumber, v))
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldAvatar, v))
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldNEQ(FieldAvatar, v))
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldIn(FieldAvatar, vs...))
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldNotIn(FieldAvatar, vs...))
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.Students {
	return predicate.Students(sql.FieldGT(FieldAvatar, v))
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.Students {
	return predicate.Students(sql.FieldGTE(FieldAvatar, v))
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.Students {
	return predicate.Students(sql.FieldLT(FieldAvatar, v))
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.Students {
	return predicate.Students(sql.FieldLTE(FieldAvatar, v))
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.Students {
	return predicate.Students(sql.FieldContains(FieldAvatar, v))
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasPrefix(FieldAvatar, v))
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasSuffix(FieldAvatar, v))
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.Students {
	return predicate.Students(sql.FieldEqualFold(FieldAvatar, v))
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.Students {
	return predicate.Students(sql.FieldContainsFold(FieldAvatar, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Students {
	return predicate.Students(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Students {
	return predicate.Students(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Students {
	return predicate.Students(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Students {
	return predicate.Students(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Students {
	return predicate.Students(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Students {
	return predicate.Students(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Students {
	return predicate.Students(sql.FieldContainsFold(FieldName, v))
}

// CollegeEQ applies the EQ predicate on the "college" field.
func CollegeEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldCollege, v))
}

// CollegeNEQ applies the NEQ predicate on the "college" field.
func CollegeNEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldNEQ(FieldCollege, v))
}

// CollegeIn applies the In predicate on the "college" field.
func CollegeIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldIn(FieldCollege, vs...))
}

// CollegeNotIn applies the NotIn predicate on the "college" field.
func CollegeNotIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldNotIn(FieldCollege, vs...))
}

// CollegeGT applies the GT predicate on the "college" field.
func CollegeGT(v string) predicate.Students {
	return predicate.Students(sql.FieldGT(FieldCollege, v))
}

// CollegeGTE applies the GTE predicate on the "college" field.
func CollegeGTE(v string) predicate.Students {
	return predicate.Students(sql.FieldGTE(FieldCollege, v))
}

// CollegeLT applies the LT predicate on the "college" field.
func CollegeLT(v string) predicate.Students {
	return predicate.Students(sql.FieldLT(FieldCollege, v))
}

// CollegeLTE applies the LTE predicate on the "college" field.
func CollegeLTE(v string) predicate.Students {
	return predicate.Students(sql.FieldLTE(FieldCollege, v))
}

// CollegeContains applies the Contains predicate on the "college" field.
func CollegeContains(v string) predicate.Students {
	return predicate.Students(sql.FieldContains(FieldCollege, v))
}

// CollegeHasPrefix applies the HasPrefix predicate on the "college" field.
func CollegeHasPrefix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasPrefix(FieldCollege, v))
}

// CollegeHasSuffix applies the HasSuffix predicate on the "college" field.
func CollegeHasSuffix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasSuffix(FieldCollege, v))
}

// CollegeEqualFold applies the EqualFold predicate on the "college" field.
func CollegeEqualFold(v string) predicate.Students {
	return predicate.Students(sql.FieldEqualFold(FieldCollege, v))
}

// CollegeContainsFold applies the ContainsFold predicate on the "college" field.
func CollegeContainsFold(v string) predicate.Students {
	return predicate.Students(sql.FieldContainsFold(FieldCollege, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Students {
	return predicate.Students(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Students {
	return predicate.Students(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Students {
	return predicate.Students(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Students {
	return predicate.Students(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Students {
	return predicate.Students(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Students {
	return predicate.Students(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Students {
	return predicate.Students(sql.FieldContainsFold(FieldPhone, v))
}

// MajorEQ applies the EQ predicate on the "major" field.
func MajorEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldMajor, v))
}

// MajorNEQ applies the NEQ predicate on the "major" field.
func MajorNEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldNEQ(FieldMajor, v))
}

// MajorIn applies the In predicate on the "major" field.
func MajorIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldIn(FieldMajor, vs...))
}

// MajorNotIn applies the NotIn predicate on the "major" field.
func MajorNotIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldNotIn(FieldMajor, vs...))
}

// MajorGT applies the GT predicate on the "major" field.
func MajorGT(v string) predicate.Students {
	return predicate.Students(sql.FieldGT(FieldMajor, v))
}

// MajorGTE applies the GTE predicate on the "major" field.
func MajorGTE(v string) predicate.Students {
	return predicate.Students(sql.FieldGTE(FieldMajor, v))
}

// MajorLT applies the LT predicate on the "major" field.
func MajorLT(v string) predicate.Students {
	return predicate.Students(sql.FieldLT(FieldMajor, v))
}

// MajorLTE applies the LTE predicate on the "major" field.
func MajorLTE(v string) predicate.Students {
	return predicate.Students(sql.FieldLTE(FieldMajor, v))
}

// MajorContains applies the Contains predicate on the "major" field.
func MajorContains(v string) predicate.Students {
	return predicate.Students(sql.FieldContains(FieldMajor, v))
}

// MajorHasPrefix applies the HasPrefix predicate on the "major" field.
func MajorHasPrefix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasPrefix(FieldMajor, v))
}

// MajorHasSuffix applies the HasSuffix predicate on the "major" field.
func MajorHasSuffix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasSuffix(FieldMajor, v))
}

// MajorEqualFold applies the EqualFold predicate on the "major" field.
func MajorEqualFold(v string) predicate.Students {
	return predicate.Students(sql.FieldEqualFold(FieldMajor, v))
}

// MajorContainsFold applies the ContainsFold predicate on the "major" field.
func MajorContainsFold(v string) predicate.Students {
	return predicate.Students(sql.FieldContainsFold(FieldMajor, v))
}

// ClassEQ applies the EQ predicate on the "class" field.
func ClassEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldClass, v))
}

// ClassNEQ applies the NEQ predicate on the "class" field.
func ClassNEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldNEQ(FieldClass, v))
}

// ClassIn applies the In predicate on the "class" field.
func ClassIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldIn(FieldClass, vs...))
}

// ClassNotIn applies the NotIn predicate on the "class" field.
func ClassNotIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldNotIn(FieldClass, vs...))
}

// ClassGT applies the GT predicate on the "class" field.
func ClassGT(v string) predicate.Students {
	return predicate.Students(sql.FieldGT(FieldClass, v))
}

// ClassGTE applies the GTE predicate on the "class" field.
func ClassGTE(v string) predicate.Students {
	return predicate.Students(sql.FieldGTE(FieldClass, v))
}

// ClassLT applies the LT predicate on the "class" field.
func ClassLT(v string) predicate.Students {
	return predicate.Students(sql.FieldLT(FieldClass, v))
}

// ClassLTE applies the LTE predicate on the "class" field.
func ClassLTE(v string) predicate.Students {
	return predicate.Students(sql.FieldLTE(FieldClass, v))
}

// ClassContains applies the Contains predicate on the "class" field.
func ClassContains(v string) predicate.Students {
	return predicate.Students(sql.FieldContains(FieldClass, v))
}

// ClassHasPrefix applies the HasPrefix predicate on the "class" field.
func ClassHasPrefix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasPrefix(FieldClass, v))
}

// ClassHasSuffix applies the HasSuffix predicate on the "class" field.
func ClassHasSuffix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasSuffix(FieldClass, v))
}

// ClassEqualFold applies the EqualFold predicate on the "class" field.
func ClassEqualFold(v string) predicate.Students {
	return predicate.Students(sql.FieldEqualFold(FieldClass, v))
}

// ClassContainsFold applies the ContainsFold predicate on the "class" field.
func ClassContainsFold(v string) predicate.Students {
	return predicate.Students(sql.FieldContainsFold(FieldClass, v))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v string) predicate.Students {
	return predicate.Students(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...string) predicate.Students {
	return predicate.Students(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v string) predicate.Students {
	return predicate.Students(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v string) predicate.Students {
	return predicate.Students(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v string) predicate.Students {
	return predicate.Students(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v string) predicate.Students {
	return predicate.Students(sql.FieldLTE(FieldNumber, v))
}

// NumberContains applies the Contains predicate on the "number" field.
func NumberContains(v string) predicate.Students {
	return predicate.Students(sql.FieldContains(FieldNumber, v))
}

// NumberHasPrefix applies the HasPrefix predicate on the "number" field.
func NumberHasPrefix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasPrefix(FieldNumber, v))
}

// NumberHasSuffix applies the HasSuffix predicate on the "number" field.
func NumberHasSuffix(v string) predicate.Students {
	return predicate.Students(sql.FieldHasSuffix(FieldNumber, v))
}

// NumberEqualFold applies the EqualFold predicate on the "number" field.
func NumberEqualFold(v string) predicate.Students {
	return predicate.Students(sql.FieldEqualFold(FieldNumber, v))
}

// NumberContainsFold applies the ContainsFold predicate on the "number" field.
func NumberContainsFold(v string) predicate.Students {
	return predicate.Students(sql.FieldContainsFold(FieldNumber, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Students {
	return predicate.Students(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Students {
	return predicate.Students(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Students) predicate.Students {
	return predicate.Students(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Students) predicate.Students {
	return predicate.Students(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Students) predicate.Students {
	return predicate.Students(sql.NotPredicates(p))
}
