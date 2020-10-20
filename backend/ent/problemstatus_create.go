// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/thanawat/app/ent/problem"
	"github.com/thanawat/app/ent/problemstatus"
)

// ProblemStatusCreate is the builder for creating a ProblemStatus entity.
type ProblemStatusCreate struct {
	config
	mutation *ProblemStatusMutation
	hooks    []Hook
}

// SetProblemstatus sets the problemstatus field.
func (psc *ProblemStatusCreate) SetProblemstatus(s string) *ProblemStatusCreate {
	psc.mutation.SetProblemstatus(s)
	return psc
}

// AddProblemstatus2problemIDs adds the problemstatus2problem edge to Problem by ids.
func (psc *ProblemStatusCreate) AddProblemstatus2problemIDs(ids ...int) *ProblemStatusCreate {
	psc.mutation.AddProblemstatus2problemIDs(ids...)
	return psc
}

// AddProblemstatus2problem adds the problemstatus2problem edges to Problem.
func (psc *ProblemStatusCreate) AddProblemstatus2problem(p ...*Problem) *ProblemStatusCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psc.AddProblemstatus2problemIDs(ids...)
}

// Mutation returns the ProblemStatusMutation object of the builder.
func (psc *ProblemStatusCreate) Mutation() *ProblemStatusMutation {
	return psc.mutation
}

// Save creates the ProblemStatus in the database.
func (psc *ProblemStatusCreate) Save(ctx context.Context) (*ProblemStatus, error) {
	if _, ok := psc.mutation.Problemstatus(); !ok {
		return nil, &ValidationError{Name: "problemstatus", err: errors.New("ent: missing required field \"problemstatus\"")}
	}
	if v, ok := psc.mutation.Problemstatus(); ok {
		if err := problemstatus.ProblemstatusValidator(v); err != nil {
			return nil, &ValidationError{Name: "problemstatus", err: fmt.Errorf("ent: validator failed for field \"problemstatus\": %w", err)}
		}
	}
	var (
		err  error
		node *ProblemStatus
	)
	if len(psc.hooks) == 0 {
		node, err = psc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProblemStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psc.mutation = mutation
			node, err = psc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(psc.hooks) - 1; i >= 0; i-- {
			mut = psc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (psc *ProblemStatusCreate) SaveX(ctx context.Context) *ProblemStatus {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (psc *ProblemStatusCreate) sqlSave(ctx context.Context) (*ProblemStatus, error) {
	ps, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ps.ID = int(id)
	return ps, nil
}

func (psc *ProblemStatusCreate) createSpec() (*ProblemStatus, *sqlgraph.CreateSpec) {
	var (
		ps    = &ProblemStatus{config: psc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: problemstatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: problemstatus.FieldID,
			},
		}
	)
	if value, ok := psc.mutation.Problemstatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: problemstatus.FieldProblemstatus,
		})
		ps.Problemstatus = value
	}
	if nodes := psc.mutation.Problemstatus2problemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   problemstatus.Problemstatus2problemTable,
			Columns: []string{problemstatus.Problemstatus2problemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: problem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return ps, _spec
}
