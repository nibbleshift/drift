// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (po *Post) Owner(ctx context.Context) (*User, error) {
	result, err := po.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = po.QueryOwner().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (po *Post) Tags(ctx context.Context) (result []*Tag, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = po.NamedTags(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = po.Edges.TagsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = po.QueryTags().All(ctx)
	}
	return result, err
}

func (po *Post) Mentions(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = po.NamedMentions(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = po.Edges.MentionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = po.QueryMentions().All(ctx)
	}
	return result, err
}

func (t *Tag) Post(ctx context.Context) (result []*Post, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedPost(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.PostOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryPost().All(ctx)
	}
	return result, err
}

func (u *User) Posts(ctx context.Context) (result []*Post, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedPosts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.PostsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryPosts().All(ctx)
	}
	return result, err
}

func (u *User) Friends(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedFriends(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.FriendsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryFriends().All(ctx)
	}
	return result, err
}

func (u *User) Followers(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedFollowers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.FollowersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryFollowers().All(ctx)
	}
	return result, err
}

func (u *User) Profile(ctx context.Context) (*UserProfile, error) {
	result, err := u.Edges.ProfileOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryProfile().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Mentions(ctx context.Context) (result []*Post, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedMentions(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.MentionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryMentions().All(ctx)
	}
	return result, err
}

func (up *UserProfile) Links(ctx context.Context) (result []*Link, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = up.NamedLinks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = up.Edges.LinksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = up.QueryLinks().All(ctx)
	}
	return result, err
}
