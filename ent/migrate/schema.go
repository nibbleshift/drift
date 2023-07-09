// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LinksColumns holds the columns for the "links" table.
	LinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "data", Type: field.TypeString, Unique: true},
	}
	// LinksTable holds the schema information for the "links" table.
	LinksTable = &schema.Table{
		Name:       "links",
		Columns:    LinksColumns,
		PrimaryKey: []*schema.Column{LinksColumns[0]},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "data", Type: field.TypeString},
		{Name: "user_posts", Type: field.TypeInt, Nullable: true},
		{Name: "user_profile_links", Type: field.TypeInt, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_user_profiles_links",
				Columns:    []*schema.Column{PostsColumns[4]},
				RefColumns: []*schema.Column{UserProfilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "data", Type: field.TypeString, Unique: true},
		{Name: "post_tags", Type: field.TypeInt, Nullable: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tags_posts_tags",
				Columns:    []*schema.Column{TagsColumns[3]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "post_mentions", Type: field.TypeInt, Nullable: true},
		{Name: "user_profile", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_posts_mentions",
				Columns:    []*schema.Column{UsersColumns[4]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_user_profiles_profile",
				Columns:    []*schema.Column{UsersColumns[5]},
				RefColumns: []*schema.Column{UserProfilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserProfilesColumns holds the columns for the "user_profiles" table.
	UserProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "avatar", Type: field.TypeString},
		{Name: "about", Type: field.TypeString},
		{Name: "location", Type: field.TypeString},
		{Name: "dob", Type: field.TypeTime},
	}
	// UserProfilesTable holds the schema information for the "user_profiles" table.
	UserProfilesTable = &schema.Table{
		Name:       "user_profiles",
		Columns:    UserProfilesColumns,
		PrimaryKey: []*schema.Column{UserProfilesColumns[0]},
	}
	// UserFriendsColumns holds the columns for the "user_friends" table.
	UserFriendsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "friend_id", Type: field.TypeInt},
	}
	// UserFriendsTable holds the schema information for the "user_friends" table.
	UserFriendsTable = &schema.Table{
		Name:       "user_friends",
		Columns:    UserFriendsColumns,
		PrimaryKey: []*schema.Column{UserFriendsColumns[0], UserFriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_friends_user_id",
				Columns:    []*schema.Column{UserFriendsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_friends_friend_id",
				Columns:    []*schema.Column{UserFriendsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserFollowersColumns holds the columns for the "user_followers" table.
	UserFollowersColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "follower_id", Type: field.TypeInt},
	}
	// UserFollowersTable holds the schema information for the "user_followers" table.
	UserFollowersTable = &schema.Table{
		Name:       "user_followers",
		Columns:    UserFollowersColumns,
		PrimaryKey: []*schema.Column{UserFollowersColumns[0], UserFollowersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_followers_user_id",
				Columns:    []*schema.Column{UserFollowersColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_followers_follower_id",
				Columns:    []*schema.Column{UserFollowersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserProfileEmailsColumns holds the columns for the "user_profile_emails" table.
	UserProfileEmailsColumns = []*schema.Column{
		{Name: "user_profile_id", Type: field.TypeInt},
		{Name: "email_id", Type: field.TypeInt},
	}
	// UserProfileEmailsTable holds the schema information for the "user_profile_emails" table.
	UserProfileEmailsTable = &schema.Table{
		Name:       "user_profile_emails",
		Columns:    UserProfileEmailsColumns,
		PrimaryKey: []*schema.Column{UserProfileEmailsColumns[0], UserProfileEmailsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_profile_emails_user_profile_id",
				Columns:    []*schema.Column{UserProfileEmailsColumns[0]},
				RefColumns: []*schema.Column{UserProfilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_profile_emails_email_id",
				Columns:    []*schema.Column{UserProfileEmailsColumns[1]},
				RefColumns: []*schema.Column{UserProfilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserProfileFollowersColumns holds the columns for the "user_profile_followers" table.
	UserProfileFollowersColumns = []*schema.Column{
		{Name: "user_profile_id", Type: field.TypeInt},
		{Name: "follower_id", Type: field.TypeInt},
	}
	// UserProfileFollowersTable holds the schema information for the "user_profile_followers" table.
	UserProfileFollowersTable = &schema.Table{
		Name:       "user_profile_followers",
		Columns:    UserProfileFollowersColumns,
		PrimaryKey: []*schema.Column{UserProfileFollowersColumns[0], UserProfileFollowersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_profile_followers_user_profile_id",
				Columns:    []*schema.Column{UserProfileFollowersColumns[0]},
				RefColumns: []*schema.Column{UserProfilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_profile_followers_follower_id",
				Columns:    []*schema.Column{UserProfileFollowersColumns[1]},
				RefColumns: []*schema.Column{UserProfilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LinksTable,
		PostsTable,
		TagsTable,
		UsersTable,
		UserProfilesTable,
		UserFriendsTable,
		UserFollowersTable,
		UserProfileEmailsTable,
		UserProfileFollowersTable,
	}
)

func init() {
	PostsTable.ForeignKeys[0].RefTable = UsersTable
	PostsTable.ForeignKeys[1].RefTable = UserProfilesTable
	TagsTable.ForeignKeys[0].RefTable = PostsTable
	UsersTable.ForeignKeys[0].RefTable = PostsTable
	UsersTable.ForeignKeys[1].RefTable = UserProfilesTable
	UserFriendsTable.ForeignKeys[0].RefTable = UsersTable
	UserFriendsTable.ForeignKeys[1].RefTable = UsersTable
	UserFollowersTable.ForeignKeys[0].RefTable = UsersTable
	UserFollowersTable.ForeignKeys[1].RefTable = UsersTable
	UserProfileEmailsTable.ForeignKeys[0].RefTable = UserProfilesTable
	UserProfileEmailsTable.ForeignKeys[1].RefTable = UserProfilesTable
	UserProfileFollowersTable.ForeignKeys[0].RefTable = UserProfilesTable
	UserProfileFollowersTable.ForeignKeys[1].RefTable = UserProfilesTable
}
