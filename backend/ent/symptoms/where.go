// Code generated by entc, DO NOT EDIT.

package symptoms

import (
	"github.com/ADMIN/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Manner applies equality check predicate on the "Manner" field. It's identical to MannerEQ.
func Manner(v string) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManner), v))
	})
}

// MannerEQ applies the EQ predicate on the "Manner" field.
func MannerEQ(v string) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManner), v))
	})
}

// MannerNEQ applies the NEQ predicate on the "Manner" field.
func MannerNEQ(v string) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldManner), v))
	})
}

// MannerIn applies the In predicate on the "Manner" field.
func MannerIn(vs ...string) predicate.Symptoms {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Symptoms(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldManner), v...))
	})
}

// MannerNotIn applies the NotIn predicate on the "Manner" field.
func MannerNotIn(vs ...string) predicate.Symptoms {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Symptoms(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldManner), v...))
	})
}

// MannerGT applies the GT predicate on the "Manner" field.
func MannerGT(v string) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldManner), v))
	})
}

// MannerGTE applies the GTE predicate on the "Manner" field.
func MannerGTE(v string) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldManner), v))
	})
}

// MannerLT applies the LT predicate on the "Manner" field.
func MannerLT(v string) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldManner), v))
	})
}

// MannerLTE applies the LTE predicate on the "Manner" field.
func MannerLTE(v string) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldManner), v))
	})
}

// MannerContains applies the Contains predicate on the "Manner" field.
func MannerContains(v string) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldManner), v))
	})
}

// MannerHasPrefix applies the HasPrefix predicate on the "Manner" field.
func MannerHasPrefix(v string) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldManner), v))
	})
}

// MannerHasSuffix applies the HasSuffix predicate on the "Manner" field.
func MannerHasSuffix(v string) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldManner), v))
	})
}

// MannerEqualFold applies the EqualFold predicate on the "Manner" field.
func MannerEqualFold(v string) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldManner), v))
	})
}

// MannerContainsFold applies the ContainsFold predicate on the "Manner" field.
func MannerContainsFold(v string) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldManner), v))
	})
}

// HasSymptomsDiagnose applies the HasEdge predicate on the "symptoms_diagnose" edge.
func HasSymptomsDiagnose() predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SymptomsDiagnoseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SymptomsDiagnoseTable, SymptomsDiagnoseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSymptomsDiagnoseWith applies the HasEdge predicate on the "symptoms_diagnose" edge with a given conditions (other predicates).
func HasSymptomsDiagnoseWith(preds ...predicate.Diagnose) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SymptomsDiagnoseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SymptomsDiagnoseTable, SymptomsDiagnoseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Symptoms) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Symptoms) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
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
func Not(p predicate.Symptoms) predicate.Symptoms {
	return predicate.Symptoms(func(s *sql.Selector) {
		p(s.Not())
	})
}
