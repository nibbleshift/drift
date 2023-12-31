// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nibbleshift/drift/ent/post"
	"github.com/nibbleshift/drift/ent/tag"
	"github.com/nibbleshift/drift/ent/user"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (pc *PostCreate) SetCreatedAt(t time.Time) *PostCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableCreatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetData sets the "data" field.
func (pc *PostCreate) SetData(s string) *PostCreate {
	pc.mutation.SetData(s)
	return pc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pc *PostCreate) SetOwnerID(id int) *PostCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (pc *PostCreate) SetNillableOwnerID(id *int) *PostCreate {
	if id != nil {
		pc = pc.SetOwnerID(*id)
	}
	return pc
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *PostCreate) SetOwner(u *User) *PostCreate {
	return pc.SetOwnerID(u.ID)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (pc *PostCreate) AddTagIDs(ids ...int) *PostCreate {
	pc.mutation.AddTagIDs(ids...)
	return pc
}

// AddTags adds the "tags" edges to the Tag entity.
func (pc *PostCreate) AddTags(t ...*Tag) *PostCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddTagIDs(ids...)
}

// AddMentionIDs adds the "mentions" edge to the User entity by IDs.
func (pc *PostCreate) AddMentionIDs(ids ...int) *PostCreate {
	pc.mutation.AddMentionIDs(ids...)
	return pc
}

// AddMentions adds the "mentions" edges to the User entity.
func (pc *PostCreate) AddMentions(u ...*User) *PostCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddMentionIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PostCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PostCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PostCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := post.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Post.created_at"`)}
	}
	if _, ok := pc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "Post.data"`)}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(post.Table, sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.Data(); ok {
		_spec.SetField(post.FieldData, field.TypeString, value)
		_node.Data = value
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.OwnerTable,
			Columns: []string{post.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_posts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.TagsTable,
			Columns: post.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MentionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.MentionsTable,
			Columns: post.MentionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Post.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PostUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (pc *PostCreate) OnConflict(opts ...sql.ConflictOption) *PostUpsertOne {
	pc.conflict = opts
	return &PostUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PostCreate) OnConflictColumns(columns ...string) *PostUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PostUpsertOne{
		create: pc,
	}
}

type (
	// PostUpsertOne is the builder for "upsert"-ing
	//  one Post node.
	PostUpsertOne struct {
		create *PostCreate
	}

	// PostUpsert is the "OnConflict" setter.
	PostUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *PostUpsert) SetCreatedAt(v time.Time) *PostUpsert {
	u.Set(post.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PostUpsert) UpdateCreatedAt() *PostUpsert {
	u.SetExcluded(post.FieldCreatedAt)
	return u
}

// SetData sets the "data" field.
func (u *PostUpsert) SetData(v string) *PostUpsert {
	u.Set(post.FieldData, v)
	return u
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *PostUpsert) UpdateData() *PostUpsert {
	u.SetExcluded(post.FieldData)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PostUpsertOne) UpdateNewValues() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PostUpsertOne) Ignore() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PostUpsertOne) DoNothing() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PostCreate.OnConflict
// documentation for more info.
func (u *PostUpsertOne) Update(set func(*PostUpsert)) *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PostUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PostUpsertOne) SetCreatedAt(v time.Time) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateCreatedAt() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetData sets the "data" field.
func (u *PostUpsertOne) SetData(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateData() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateData()
	})
}

// Exec executes the query.
func (u *PostUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PostCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PostUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PostUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PostUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	builders []*PostCreate
	conflict []sql.ConflictOption
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PostCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PostCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Post.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PostUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (pcb *PostCreateBulk) OnConflict(opts ...sql.ConflictOption) *PostUpsertBulk {
	pcb.conflict = opts
	return &PostUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PostCreateBulk) OnConflictColumns(columns ...string) *PostUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PostUpsertBulk{
		create: pcb,
	}
}

// PostUpsertBulk is the builder for "upsert"-ing
// a bulk of Post nodes.
type PostUpsertBulk struct {
	create *PostCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PostUpsertBulk) UpdateNewValues() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PostUpsertBulk) Ignore() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PostUpsertBulk) DoNothing() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PostCreateBulk.OnConflict
// documentation for more info.
func (u *PostUpsertBulk) Update(set func(*PostUpsert)) *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PostUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PostUpsertBulk) SetCreatedAt(v time.Time) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateCreatedAt() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetData sets the "data" field.
func (u *PostUpsertBulk) SetData(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateData() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateData()
	})
}

// Exec executes the query.
func (u *PostUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PostCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PostCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PostUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
