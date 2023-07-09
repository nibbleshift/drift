package drift

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/nibbleshift/drift/ent"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.LinkWhereInput) (*ent.LinkConnection, error) {
	return r.client.Link.Query().
		Paginate(ctx, after, first, before, last)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	return r.client.Post.Query().
		Paginate(ctx, after, first, before, last)
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.TagWhereInput) (*ent.TagConnection, error) {
	return r.client.Tag.Query().
		Paginate(ctx, after, first, before, last)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last)
}

// UserProfiles is the resolver for the userProfiles field.
func (r *queryResolver) UserProfiles(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.UserProfileWhereInput) (*ent.UserProfileConnection, error) {
	return r.client.UserProfile.Query().
		Paginate(ctx, after, first, before, last)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
