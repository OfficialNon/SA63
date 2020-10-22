// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/ADMIN/app/ent/diagnose"
	"github.com/ADMIN/app/ent/predicate"
	"github.com/ADMIN/app/ent/symptoms"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// SymptomsUpdate is the builder for updating Symptoms entities.
type SymptomsUpdate struct {
	config
	hooks      []Hook
	mutation   *SymptomsMutation
	predicates []predicate.Symptoms
}

// Where adds a new predicate for the builder.
func (su *SymptomsUpdate) Where(ps ...predicate.Symptoms) *SymptomsUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetManner sets the Manner field.
func (su *SymptomsUpdate) SetManner(s string) *SymptomsUpdate {
	su.mutation.SetManner(s)
	return su
}

// AddSymptomsDiagnoseIDs adds the symptoms_diagnose edge to Diagnose by ids.
func (su *SymptomsUpdate) AddSymptomsDiagnoseIDs(ids ...int) *SymptomsUpdate {
	su.mutation.AddSymptomsDiagnoseIDs(ids...)
	return su
}

// AddSymptomsDiagnose adds the symptoms_diagnose edges to Diagnose.
func (su *SymptomsUpdate) AddSymptomsDiagnose(d ...*Diagnose) *SymptomsUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return su.AddSymptomsDiagnoseIDs(ids...)
}

// Mutation returns the SymptomsMutation object of the builder.
func (su *SymptomsUpdate) Mutation() *SymptomsMutation {
	return su.mutation
}

// RemoveSymptomsDiagnoseIDs removes the symptoms_diagnose edge to Diagnose by ids.
func (su *SymptomsUpdate) RemoveSymptomsDiagnoseIDs(ids ...int) *SymptomsUpdate {
	su.mutation.RemoveSymptomsDiagnoseIDs(ids...)
	return su
}

// RemoveSymptomsDiagnose removes symptoms_diagnose edges to Diagnose.
func (su *SymptomsUpdate) RemoveSymptomsDiagnose(d ...*Diagnose) *SymptomsUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return su.RemoveSymptomsDiagnoseIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *SymptomsUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := su.mutation.Manner(); ok {
		if err := symptoms.MannerValidator(v); err != nil {
			return 0, &ValidationError{Name: "Manner", err: fmt.Errorf("ent: validator failed for field \"Manner\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SymptomsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SymptomsUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SymptomsUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SymptomsUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SymptomsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   symptoms.Table,
			Columns: symptoms.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: symptoms.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Manner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: symptoms.FieldManner,
		})
	}
	if nodes := su.mutation.RemovedSymptomsDiagnoseIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SymptomsDiagnoseIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{symptoms.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SymptomsUpdateOne is the builder for updating a single Symptoms entity.
type SymptomsUpdateOne struct {
	config
	hooks    []Hook
	mutation *SymptomsMutation
}

// SetManner sets the Manner field.
func (suo *SymptomsUpdateOne) SetManner(s string) *SymptomsUpdateOne {
	suo.mutation.SetManner(s)
	return suo
}

// AddSymptomsDiagnoseIDs adds the symptoms_diagnose edge to Diagnose by ids.
func (suo *SymptomsUpdateOne) AddSymptomsDiagnoseIDs(ids ...int) *SymptomsUpdateOne {
	suo.mutation.AddSymptomsDiagnoseIDs(ids...)
	return suo
}

// AddSymptomsDiagnose adds the symptoms_diagnose edges to Diagnose.
func (suo *SymptomsUpdateOne) AddSymptomsDiagnose(d ...*Diagnose) *SymptomsUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return suo.AddSymptomsDiagnoseIDs(ids...)
}

// Mutation returns the SymptomsMutation object of the builder.
func (suo *SymptomsUpdateOne) Mutation() *SymptomsMutation {
	return suo.mutation
}

// RemoveSymptomsDiagnoseIDs removes the symptoms_diagnose edge to Diagnose by ids.
func (suo *SymptomsUpdateOne) RemoveSymptomsDiagnoseIDs(ids ...int) *SymptomsUpdateOne {
	suo.mutation.RemoveSymptomsDiagnoseIDs(ids...)
	return suo
}

// RemoveSymptomsDiagnose removes symptoms_diagnose edges to Diagnose.
func (suo *SymptomsUpdateOne) RemoveSymptomsDiagnose(d ...*Diagnose) *SymptomsUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return suo.RemoveSymptomsDiagnoseIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *SymptomsUpdateOne) Save(ctx context.Context) (*Symptoms, error) {
	if v, ok := suo.mutation.Manner(); ok {
		if err := symptoms.MannerValidator(v); err != nil {
			return nil, &ValidationError{Name: "Manner", err: fmt.Errorf("ent: validator failed for field \"Manner\": %w", err)}
		}
	}

	var (
		err  error
		node *Symptoms
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SymptomsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SymptomsUpdateOne) SaveX(ctx context.Context) *Symptoms {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *SymptomsUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SymptomsUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SymptomsUpdateOne) sqlSave(ctx context.Context) (s *Symptoms, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   symptoms.Table,
			Columns: symptoms.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: symptoms.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Symptoms.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.Manner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: symptoms.FieldManner,
		})
	}
	if nodes := suo.mutation.RemovedSymptomsDiagnoseIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SymptomsDiagnoseIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Symptoms{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{symptoms.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}