// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ZEQUANR/zhulong/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Account applies equality check predicate on the "account" field. It's identical to AccountEQ.
func Account(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAccount, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRole, v))
}

// AccountEQ applies the EQ predicate on the "account" field.
func AccountEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAccount, v))
}

// AccountNEQ applies the NEQ predicate on the "account" field.
func AccountNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAccount, v))
}

// AccountIn applies the In predicate on the "account" field.
func AccountIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAccount, vs...))
}

// AccountNotIn applies the NotIn predicate on the "account" field.
func AccountNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAccount, vs...))
}

// AccountGT applies the GT predicate on the "account" field.
func AccountGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAccount, v))
}

// AccountGTE applies the GTE predicate on the "account" field.
func AccountGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAccount, v))
}

// AccountLT applies the LT predicate on the "account" field.
func AccountLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAccount, v))
}

// AccountLTE applies the LTE predicate on the "account" field.
func AccountLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAccount, v))
}

// AccountContains applies the Contains predicate on the "account" field.
func AccountContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAccount, v))
}

// AccountHasPrefix applies the HasPrefix predicate on the "account" field.
func AccountHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAccount, v))
}

// AccountHasSuffix applies the HasSuffix predicate on the "account" field.
func AccountHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAccount, v))
}

// AccountEqualFold applies the EqualFold predicate on the "account" field.
func AccountEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAccount, v))
}

// AccountContainsFold applies the ContainsFold predicate on the "account" field.
func AccountContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAccount, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRole, vs...))
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldRole, v))
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldRole, v))
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldRole, v))
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldRole, v))
}

// HasAdministrators applies the HasEdge predicate on the "administrators" edge.
func HasAdministrators() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, AdministratorsTable, AdministratorsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAdministratorsWith applies the HasEdge predicate on the "administrators" edge with a given conditions (other predicates).
func HasAdministratorsWith(preds ...predicate.Administrators) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newAdministratorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudents applies the HasEdge predicate on the "students" edge.
func HasStudents() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, StudentsTable, StudentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudentsWith applies the HasEdge predicate on the "students" edge with a given conditions (other predicates).
func HasStudentsWith(preds ...predicate.Students) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newStudentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeachers applies the HasEdge predicate on the "teachers" edge.
func HasTeachers() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, TeachersTable, TeachersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeachersWith applies the HasEdge predicate on the "teachers" edge with a given conditions (other predicates).
func HasTeachersWith(preds ...predicate.Teachers) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newTeachersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasThesis applies the HasEdge predicate on the "thesis" edge.
func HasThesis() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ThesisTable, ThesisColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasThesisWith applies the HasEdge predicate on the "thesis" edge with a given conditions (other predicates).
func HasThesisWith(preds ...predicate.Thesis) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newThesisStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReviews applies the HasEdge predicate on the "reviews" edge.
func HasReviews() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReviewsTable, ReviewsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReviewsWith applies the HasEdge predicate on the "reviews" edge with a given conditions (other predicates).
func HasReviewsWith(preds ...predicate.Reviews) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newReviewsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasExamineThesis applies the HasEdge predicate on the "examineThesis" edge.
func HasExamineThesis() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ExamineThesisTable, ExamineThesisColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExamineThesisWith applies the HasEdge predicate on the "examineThesis" edge with a given conditions (other predicates).
func HasExamineThesisWith(preds ...predicate.Thesis) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newExamineThesisStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
