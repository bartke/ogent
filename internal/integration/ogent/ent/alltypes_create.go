// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"ariga.io/ogent/internal/integration/ogent/ent/alltypes"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AllTypesCreate is the builder for creating a AllTypes entity.
type AllTypesCreate struct {
	config
	mutation *AllTypesMutation
	hooks    []Hook
}

// SetInt sets the "int" field.
func (atc *AllTypesCreate) SetInt(i int) *AllTypesCreate {
	atc.mutation.SetInt(i)
	return atc
}

// SetInt8 sets the "int8" field.
func (atc *AllTypesCreate) SetInt8(i int8) *AllTypesCreate {
	atc.mutation.SetInt8(i)
	return atc
}

// SetInt16 sets the "int16" field.
func (atc *AllTypesCreate) SetInt16(i int16) *AllTypesCreate {
	atc.mutation.SetInt16(i)
	return atc
}

// SetInt32 sets the "int32" field.
func (atc *AllTypesCreate) SetInt32(i int32) *AllTypesCreate {
	atc.mutation.SetInt32(i)
	return atc
}

// SetInt64 sets the "int64" field.
func (atc *AllTypesCreate) SetInt64(i int64) *AllTypesCreate {
	atc.mutation.SetInt64(i)
	return atc
}

// SetUint sets the "uint" field.
func (atc *AllTypesCreate) SetUint(u uint) *AllTypesCreate {
	atc.mutation.SetUint(u)
	return atc
}

// SetUint8 sets the "uint8" field.
func (atc *AllTypesCreate) SetUint8(u uint8) *AllTypesCreate {
	atc.mutation.SetUint8(u)
	return atc
}

// SetUint16 sets the "uint16" field.
func (atc *AllTypesCreate) SetUint16(u uint16) *AllTypesCreate {
	atc.mutation.SetUint16(u)
	return atc
}

// SetUint32 sets the "uint32" field.
func (atc *AllTypesCreate) SetUint32(u uint32) *AllTypesCreate {
	atc.mutation.SetUint32(u)
	return atc
}

// SetUint64 sets the "uint64" field.
func (atc *AllTypesCreate) SetUint64(u uint64) *AllTypesCreate {
	atc.mutation.SetUint64(u)
	return atc
}

// SetFloat32 sets the "float32" field.
func (atc *AllTypesCreate) SetFloat32(f float32) *AllTypesCreate {
	atc.mutation.SetFloat32(f)
	return atc
}

// SetFloat64 sets the "float64" field.
func (atc *AllTypesCreate) SetFloat64(f float64) *AllTypesCreate {
	atc.mutation.SetFloat64(f)
	return atc
}

// SetStringType sets the "string_type" field.
func (atc *AllTypesCreate) SetStringType(s string) *AllTypesCreate {
	atc.mutation.SetStringType(s)
	return atc
}

// SetBool sets the "bool" field.
func (atc *AllTypesCreate) SetBool(b bool) *AllTypesCreate {
	atc.mutation.SetBool(b)
	return atc
}

// SetUUID sets the "uuid" field.
func (atc *AllTypesCreate) SetUUID(u uuid.UUID) *AllTypesCreate {
	atc.mutation.SetUUID(u)
	return atc
}

// SetTime sets the "time" field.
func (atc *AllTypesCreate) SetTime(t time.Time) *AllTypesCreate {
	atc.mutation.SetTime(t)
	return atc
}

// SetText sets the "text" field.
func (atc *AllTypesCreate) SetText(s string) *AllTypesCreate {
	atc.mutation.SetText(s)
	return atc
}

// SetState sets the "state" field.
func (atc *AllTypesCreate) SetState(a alltypes.State) *AllTypesCreate {
	atc.mutation.SetState(a)
	return atc
}

// SetBytes sets the "bytes" field.
func (atc *AllTypesCreate) SetBytes(b []byte) *AllTypesCreate {
	atc.mutation.SetBytes(b)
	return atc
}

// SetID sets the "id" field.
func (atc *AllTypesCreate) SetID(u uint32) *AllTypesCreate {
	atc.mutation.SetID(u)
	return atc
}

// Mutation returns the AllTypesMutation object of the builder.
func (atc *AllTypesCreate) Mutation() *AllTypesMutation {
	return atc.mutation
}

// Save creates the AllTypes in the database.
func (atc *AllTypesCreate) Save(ctx context.Context) (*AllTypes, error) {
	var (
		err  error
		node *AllTypes
	)
	if len(atc.hooks) == 0 {
		if err = atc.check(); err != nil {
			return nil, err
		}
		node, err = atc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AllTypesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atc.check(); err != nil {
				return nil, err
			}
			atc.mutation = mutation
			if node, err = atc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(atc.hooks) - 1; i >= 0; i-- {
			if atc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, atc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AllTypes)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AllTypesMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (atc *AllTypesCreate) SaveX(ctx context.Context) *AllTypes {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *AllTypesCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *AllTypesCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *AllTypesCreate) check() error {
	if _, ok := atc.mutation.Int(); !ok {
		return &ValidationError{Name: "int", err: errors.New(`ent: missing required field "AllTypes.int"`)}
	}
	if _, ok := atc.mutation.Int8(); !ok {
		return &ValidationError{Name: "int8", err: errors.New(`ent: missing required field "AllTypes.int8"`)}
	}
	if _, ok := atc.mutation.Int16(); !ok {
		return &ValidationError{Name: "int16", err: errors.New(`ent: missing required field "AllTypes.int16"`)}
	}
	if _, ok := atc.mutation.Int32(); !ok {
		return &ValidationError{Name: "int32", err: errors.New(`ent: missing required field "AllTypes.int32"`)}
	}
	if _, ok := atc.mutation.Int64(); !ok {
		return &ValidationError{Name: "int64", err: errors.New(`ent: missing required field "AllTypes.int64"`)}
	}
	if _, ok := atc.mutation.Uint(); !ok {
		return &ValidationError{Name: "uint", err: errors.New(`ent: missing required field "AllTypes.uint"`)}
	}
	if _, ok := atc.mutation.Uint8(); !ok {
		return &ValidationError{Name: "uint8", err: errors.New(`ent: missing required field "AllTypes.uint8"`)}
	}
	if _, ok := atc.mutation.Uint16(); !ok {
		return &ValidationError{Name: "uint16", err: errors.New(`ent: missing required field "AllTypes.uint16"`)}
	}
	if _, ok := atc.mutation.Uint32(); !ok {
		return &ValidationError{Name: "uint32", err: errors.New(`ent: missing required field "AllTypes.uint32"`)}
	}
	if _, ok := atc.mutation.Uint64(); !ok {
		return &ValidationError{Name: "uint64", err: errors.New(`ent: missing required field "AllTypes.uint64"`)}
	}
	if _, ok := atc.mutation.Float32(); !ok {
		return &ValidationError{Name: "float32", err: errors.New(`ent: missing required field "AllTypes.float32"`)}
	}
	if _, ok := atc.mutation.Float64(); !ok {
		return &ValidationError{Name: "float64", err: errors.New(`ent: missing required field "AllTypes.float64"`)}
	}
	if _, ok := atc.mutation.StringType(); !ok {
		return &ValidationError{Name: "string_type", err: errors.New(`ent: missing required field "AllTypes.string_type"`)}
	}
	if _, ok := atc.mutation.Bool(); !ok {
		return &ValidationError{Name: "bool", err: errors.New(`ent: missing required field "AllTypes.bool"`)}
	}
	if _, ok := atc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "AllTypes.uuid"`)}
	}
	if _, ok := atc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "AllTypes.time"`)}
	}
	if _, ok := atc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "AllTypes.text"`)}
	}
	if _, ok := atc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "AllTypes.state"`)}
	}
	if v, ok := atc.mutation.State(); ok {
		if err := alltypes.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "AllTypes.state": %w`, err)}
		}
	}
	if _, ok := atc.mutation.Bytes(); !ok {
		return &ValidationError{Name: "bytes", err: errors.New(`ent: missing required field "AllTypes.bytes"`)}
	}
	return nil
}

func (atc *AllTypesCreate) sqlSave(ctx context.Context) (*AllTypes, error) {
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (atc *AllTypesCreate) createSpec() (*AllTypes, *sqlgraph.CreateSpec) {
	var (
		_node = &AllTypes{config: atc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: alltypes.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: alltypes.FieldID,
			},
		}
	)
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := atc.mutation.Int(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: alltypes.FieldInt,
		})
		_node.Int = value
	}
	if value, ok := atc.mutation.Int8(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: alltypes.FieldInt8,
		})
		_node.Int8 = value
	}
	if value, ok := atc.mutation.Int16(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: alltypes.FieldInt16,
		})
		_node.Int16 = value
	}
	if value, ok := atc.mutation.Int32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: alltypes.FieldInt32,
		})
		_node.Int32 = value
	}
	if value, ok := atc.mutation.Int64(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: alltypes.FieldInt64,
		})
		_node.Int64 = value
	}
	if value, ok := atc.mutation.Uint(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: alltypes.FieldUint,
		})
		_node.Uint = value
	}
	if value, ok := atc.mutation.Uint8(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: alltypes.FieldUint8,
		})
		_node.Uint8 = value
	}
	if value, ok := atc.mutation.Uint16(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint16,
			Value:  value,
			Column: alltypes.FieldUint16,
		})
		_node.Uint16 = value
	}
	if value, ok := atc.mutation.Uint32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: alltypes.FieldUint32,
		})
		_node.Uint32 = value
	}
	if value, ok := atc.mutation.Uint64(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: alltypes.FieldUint64,
		})
		_node.Uint64 = value
	}
	if value, ok := atc.mutation.Float32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: alltypes.FieldFloat32,
		})
		_node.Float32 = value
	}
	if value, ok := atc.mutation.Float64(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: alltypes.FieldFloat64,
		})
		_node.Float64 = value
	}
	if value, ok := atc.mutation.StringType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alltypes.FieldStringType,
		})
		_node.StringType = value
	}
	if value, ok := atc.mutation.Bool(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: alltypes.FieldBool,
		})
		_node.Bool = value
	}
	if value, ok := atc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: alltypes.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := atc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: alltypes.FieldTime,
		})
		_node.Time = value
	}
	if value, ok := atc.mutation.Text(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alltypes.FieldText,
		})
		_node.Text = value
	}
	if value, ok := atc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: alltypes.FieldState,
		})
		_node.State = value
	}
	if value, ok := atc.mutation.Bytes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: alltypes.FieldBytes,
		})
		_node.Bytes = value
	}
	return _node, _spec
}

// AllTypesCreateBulk is the builder for creating many AllTypes entities in bulk.
type AllTypesCreateBulk struct {
	config
	builders []*AllTypesCreate
}

// Save creates the AllTypes entities in the database.
func (atcb *AllTypesCreateBulk) Save(ctx context.Context) ([]*AllTypes, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*AllTypes, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AllTypesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *AllTypesCreateBulk) SaveX(ctx context.Context) []*AllTypes {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *AllTypesCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *AllTypesCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}
