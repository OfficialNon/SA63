// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/ADMIN/app/ent/diagnose"
	"github.com/ADMIN/app/ent/symptoms"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// SymptomsCreate is the builder for creating a Symptoms entity.
type SymptomsCreate struct {
	config
	mutation *SymptomsMutation
	hooks    []Hook
}

// SetManner sets the Manner field.
func (sc *SymptomsCreate) SetManner(s string) *SymptomsCreate {
	sc.mutation.SetManner(s)
	return sc
}

// AddSymptomsDiagnoseIDs adds the symptoms_diagnose edge to Diagnose by ids.
func (sc *SymptomsCreate) AddSymptomsDiagnoseIDs(ids ...int) *SymptomsCreate {
	sc.mutation.AddSymptomsDiagnoseIDs(ids...)
	return sc
}

// AddSymptomsDiagnose adds the symptoms_diagnose edges to Diagnose.
func (sc *SymptomsCreate) AddSymptomsDiagnose(d ...*Diagnose) *SymptomsCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return sc.AddSymptomsDiagnoseIDs(ids...)
}

// Mutation returns the SymptomsMutation object of the builder.
func (sc *SymptomsCreate) Mutation() *SymptomsMutation {
	return sc.mutation
}

// Save creates the Symptoms in the database.
func (sc *SymptomsCreate) Save(ctx context.Context) (*Symptoms, error) {
	if _, ok := sc.mutation.Manner(); !ok {
		return nil, &ValidationError{Name: "Manner", err: errors.New("ent: missing required field \"Manner\"")}
	}
	if v, ok := sc.mutation.Manner(); ok {
		if err := symptoms.MannerValidator(v); err != nil {
			return nil, &ValidationError{Name: "Manner", err: fmt.Errorf("ent: validator failed for field \"Manner\": %w", err)}
		}
	}
	var (
		err  error
		node *Symptoms
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SymptomsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SymptomsCreate) SaveX(ctx context.Context) *Symptoms {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *SymptomsCreate) sqlSave(ctx context.Context) (*Symptoms, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *SymptomsCreate) createSpec() (*Symptoms, *sqlgraph.CreateSpec) {
	var (
		s     = &Symptoms{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: symptoms.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: symptoms.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Manner(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: symptoms.FieldManner,
		})
		s.Manner = value
	}
	if nodes := sc.mutation.SymptomsDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   symptoms.SymptomsDiagnoseTable,
			Columns: []string{symptoms.SymptomsDiagnoseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}