// Code generated by entc, DO NOT EDIT.

package problemtitle

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/thanawat/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Problemtitle applies equality check predicate on the "problemtitle" field. It's identical to ProblemtitleEQ.
func Problemtitle(v string) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProblemtitle), v))
	})
}

// ProblemtitleEQ applies the EQ predicate on the "problemtitle" field.
func ProblemtitleEQ(v string) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProblemtitle), v))
	})
}

// ProblemtitleNEQ applies the NEQ predicate on the "problemtitle" field.
func ProblemtitleNEQ(v string) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProblemtitle), v))
	})
}

// ProblemtitleIn applies the In predicate on the "problemtitle" field.
func ProblemtitleIn(vs ...string) predicate.ProblemTitle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProblemTitle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProblemtitle), v...))
	})
}

// ProblemtitleNotIn applies the NotIn predicate on the "problemtitle" field.
func ProblemtitleNotIn(vs ...string) predicate.ProblemTitle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProblemTitle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProblemtitle), v...))
	})
}

// ProblemtitleGT applies the GT predicate on the "problemtitle" field.
func ProblemtitleGT(v string) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProblemtitle), v))
	})
}

// ProblemtitleGTE applies the GTE predicate on the "problemtitle" field.
func ProblemtitleGTE(v string) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProblemtitle), v))
	})
}

// ProblemtitleLT applies the LT predicate on the "problemtitle" field.
func ProblemtitleLT(v string) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProblemtitle), v))
	})
}

// ProblemtitleLTE applies the LTE predicate on the "problemtitle" field.
func ProblemtitleLTE(v string) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProblemtitle), v))
	})
}

// ProblemtitleContains applies the Contains predicate on the "problemtitle" field.
func ProblemtitleContains(v string) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProblemtitle), v))
	})
}

// ProblemtitleHasPrefix applies the HasPrefix predicate on the "problemtitle" field.
func ProblemtitleHasPrefix(v string) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProblemtitle), v))
	})
}

// ProblemtitleHasSuffix applies the HasSuffix predicate on the "problemtitle" field.
func ProblemtitleHasSuffix(v string) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProblemtitle), v))
	})
}

// ProblemtitleEqualFold applies the EqualFold predicate on the "problemtitle" field.
func ProblemtitleEqualFold(v string) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProblemtitle), v))
	})
}

// ProblemtitleContainsFold applies the ContainsFold predicate on the "problemtitle" field.
func ProblemtitleContainsFold(v string) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProblemtitle), v))
	})
}

// HasProblemtitle2problem applies the HasEdge predicate on the "problemtitle2problem" edge.
func HasProblemtitle2problem() predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Problemtitle2problemTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, Problemtitle2problemTable, Problemtitle2problemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProblemtitle2problemWith applies the HasEdge predicate on the "problemtitle2problem" edge with a given conditions (other predicates).
func HasProblemtitle2problemWith(preds ...predicate.Problem) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Problemtitle2problemInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, Problemtitle2problemTable, Problemtitle2problemColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.ProblemTitle) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.ProblemTitle) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProblemTitle) predicate.ProblemTitle {
	return predicate.ProblemTitle(func(s *sql.Selector) {
		p(s.Not())
	})
}