// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/ADMIN/app/ent/diagnose"
	"github.com/ADMIN/app/ent/doctor"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DoctorCreate is the builder for creating a Doctor entity.
type DoctorCreate struct {
	config
	mutation *DoctorMutation
	hooks    []Hook
}

// SetDoctorNAME sets the Doctor_NAME field.
func (dc *DoctorCreate) SetDoctorNAME(s string) *DoctorCreate {
	dc.mutation.SetDoctorNAME(s)
	return dc
}

// SetDoctorEmail sets the Doctor_Email field.
func (dc *DoctorCreate) SetDoctorEmail(s string) *DoctorCreate {
	dc.mutation.SetDoctorEmail(s)
	return dc
}

// AddDoctorDiagnoseIDs adds the doctor_diagnose edge to Diagnose by ids.
func (dc *DoctorCreate) AddDoctorDiagnoseIDs(ids ...int) *DoctorCreate {
	dc.mutation.AddDoctorDiagnoseIDs(ids...)
	return dc
}

// AddDoctorDiagnose adds the doctor_diagnose edges to Diagnose.
func (dc *DoctorCreate) AddDoctorDiagnose(d ...*Diagnose) *DoctorCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDoctorDiagnoseIDs(ids...)
}

// Mutation returns the DoctorMutation object of the builder.
func (dc *DoctorCreate) Mutation() *DoctorMutation {
	return dc.mutation
}

// Save creates the Doctor in the database.
func (dc *DoctorCreate) Save(ctx context.Context) (*Doctor, error) {
	if _, ok := dc.mutation.DoctorNAME(); !ok {
		return nil, &ValidationError{Name: "Doctor_NAME", err: errors.New("ent: missing required field \"Doctor_NAME\"")}
	}
	if v, ok := dc.mutation.DoctorNAME(); ok {
		if err := doctor.DoctorNAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "Doctor_NAME", err: fmt.Errorf("ent: validator failed for field \"Doctor_NAME\": %w", err)}
		}
	}
	if _, ok := dc.mutation.DoctorEmail(); !ok {
		return nil, &ValidationError{Name: "Doctor_Email", err: errors.New("ent: missing required field \"Doctor_Email\"")}
	}
	if v, ok := dc.mutation.DoctorEmail(); ok {
		if err := doctor.DoctorEmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "Doctor_Email", err: fmt.Errorf("ent: validator failed for field \"Doctor_Email\": %w", err)}
		}
	}
	var (
		err  error
		node *Doctor
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoctorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DoctorCreate) SaveX(ctx context.Context) *Doctor {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DoctorCreate) sqlSave(ctx context.Context) (*Doctor, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DoctorCreate) createSpec() (*Doctor, *sqlgraph.CreateSpec) {
	var (
		d     = &Doctor{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: doctor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.DoctorNAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorNAME,
		})
		d.DoctorNAME = value
	}
	if value, ok := dc.mutation.DoctorEmail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorEmail,
		})
		d.DoctorEmail = value
	}
	if nodes := dc.mutation.DoctorDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorDiagnoseTable,
			Columns: []string{doctor.DoctorDiagnoseColumn},
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
	return d, _spec
}
