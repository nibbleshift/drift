// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/nibbleshift/drift/ent/link"
	"github.com/nibbleshift/drift/ent/post"
	"github.com/nibbleshift/drift/ent/tag"
	"github.com/nibbleshift/drift/ent/user"
	"github.com/nibbleshift/drift/ent/userprofile"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (l *LinkQuery) CollectFields(ctx context.Context, satisfies ...string) (*LinkQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return l, nil
	}
	if err := l.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *LinkQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(link.Columns))
		selectedFields = []string{link.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "createdAt":
			if _, ok := fieldSeen[link.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, link.FieldCreatedAt)
				fieldSeen[link.FieldCreatedAt] = struct{}{}
			}
		case "data":
			if _, ok := fieldSeen[link.FieldData]; !ok {
				selectedFields = append(selectedFields, link.FieldData)
				fieldSeen[link.FieldData] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		l.Select(selectedFields...)
	}
	return nil
}

type linkPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []LinkPaginateOption
}

func newLinkPaginateArgs(rv map[string]any) *linkPaginateArgs {
	args := &linkPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*LinkWhereInput); ok {
		args.opts = append(args.opts, WithLinkFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (po *PostQuery) CollectFields(ctx context.Context, satisfies ...string) (*PostQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return po, nil
	}
	if err := po.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return po, nil
}

func (po *PostQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(post.Columns))
		selectedFields = []string{post.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "owner":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: po.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			po.withOwner = query
		case "tags":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TagClient{config: po.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			po.WithNamedTags(alias, func(wq *TagQuery) {
				*wq = *query
			})
		case "mentions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: po.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			po.WithNamedMentions(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[post.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, post.FieldCreatedAt)
				fieldSeen[post.FieldCreatedAt] = struct{}{}
			}
		case "data":
			if _, ok := fieldSeen[post.FieldData]; !ok {
				selectedFields = append(selectedFields, post.FieldData)
				fieldSeen[post.FieldData] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		po.Select(selectedFields...)
	}
	return nil
}

type postPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PostPaginateOption
}

func newPostPaginateArgs(rv map[string]any) *postPaginateArgs {
	args := &postPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &PostOrder{Field: &PostOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithPostOrder(order))
			}
		case *PostOrder:
			if v != nil {
				args.opts = append(args.opts, WithPostOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*PostWhereInput); ok {
		args.opts = append(args.opts, WithPostFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TagQuery) CollectFields(ctx context.Context, satisfies ...string) (*TagQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TagQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(tag.Columns))
		selectedFields = []string{tag.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "post":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PostClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			t.WithNamedPost(alias, func(wq *PostQuery) {
				*wq = *query
			})
		case "data":
			if _, ok := fieldSeen[tag.FieldData]; !ok {
				selectedFields = append(selectedFields, tag.FieldData)
				fieldSeen[tag.FieldData] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		t.Select(selectedFields...)
	}
	return nil
}

type tagPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TagPaginateOption
}

func newTagPaginateArgs(rv map[string]any) *tagPaginateArgs {
	args := &tagPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*TagWhereInput); ok {
		args.opts = append(args.opts, WithTagFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "posts":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PostClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedPosts(alias, func(wq *PostQuery) {
				*wq = *query
			})
		case "friends":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedFriends(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "followers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedFollowers(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "profile":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserProfileClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.withProfile = query
		case "mentions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PostClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedMentions(alias, func(wq *PostQuery) {
				*wq = *query
			})
		case "username":
			if _, ok := fieldSeen[user.FieldUsername]; !ok {
				selectedFields = append(selectedFields, user.FieldUsername)
				fieldSeen[user.FieldUsername] = struct{}{}
			}
		case "firstName":
			if _, ok := fieldSeen[user.FieldFirstName]; !ok {
				selectedFields = append(selectedFields, user.FieldFirstName)
				fieldSeen[user.FieldFirstName] = struct{}{}
			}
		case "lastName":
			if _, ok := fieldSeen[user.FieldLastName]; !ok {
				selectedFields = append(selectedFields, user.FieldLastName)
				fieldSeen[user.FieldLastName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (up *UserProfileQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserProfileQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return up, nil
	}
	if err := up.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return up, nil
}

func (up *UserProfileQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(userprofile.Columns))
		selectedFields = []string{userprofile.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "links":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&LinkClient{config: up.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			up.WithNamedLinks(alias, func(wq *LinkQuery) {
				*wq = *query
			})
		case "avatar":
			if _, ok := fieldSeen[userprofile.FieldAvatar]; !ok {
				selectedFields = append(selectedFields, userprofile.FieldAvatar)
				fieldSeen[userprofile.FieldAvatar] = struct{}{}
			}
		case "about":
			if _, ok := fieldSeen[userprofile.FieldAbout]; !ok {
				selectedFields = append(selectedFields, userprofile.FieldAbout)
				fieldSeen[userprofile.FieldAbout] = struct{}{}
			}
		case "location":
			if _, ok := fieldSeen[userprofile.FieldLocation]; !ok {
				selectedFields = append(selectedFields, userprofile.FieldLocation)
				fieldSeen[userprofile.FieldLocation] = struct{}{}
			}
		case "dob":
			if _, ok := fieldSeen[userprofile.FieldDob]; !ok {
				selectedFields = append(selectedFields, userprofile.FieldDob)
				fieldSeen[userprofile.FieldDob] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		up.Select(selectedFields...)
	}
	return nil
}

type userprofilePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserProfilePaginateOption
}

func newUserProfilePaginateArgs(rv map[string]any) *userprofilePaginateArgs {
	args := &userprofilePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserProfileWhereInput); ok {
		args.opts = append(args.opts, WithUserProfileFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
