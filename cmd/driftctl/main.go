package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/davecgh/go-spew/spew"
	"github.com/icrowley/fake"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/nibbleshift/drift/ent"
)

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	client := Open("postgresql://postgres:postgres@127.0.0.1/drift")
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ctx := context.Background()

	steve, err := client.User.
		Create().
		SetUsername("nibbleshift").
		SetFirstName("Steve").
		SetLastName("McDaniel").Save(ctx)

	if err != nil {
		log.Println(err)
	}

	corey, err := client.User.
		Create().
		SetUsername("corey").
		SetFirstName("Corey").
		SetLastName("Gaspard").Save(ctx)

	if err != nil {
		log.Println(err)
	}

	adam, err := client.User.
		Create().
		SetUsername("adam").
		SetFirstName("Adam").
		SetLastName("Allen").Save(ctx)

	if err != nil {
		log.Println(err)
	}

	utter, err := client.Post.Create().SetData(fake.Sentence()).SetOwner(steve).Save(ctx)

	if err != nil {
		log.Println(err)
	}

	steve.Update().AddPosts(utter).Save(ctx)

	time.Sleep(time.Second * 10)

	utter, err = client.Post.Create().SetData(fake.Sentence()).SetOwner(corey).Save(ctx)

	if err != nil {
		log.Println(err)
	}

	corey.Update().AddPosts(utter).Save(ctx)

	time.Sleep(time.Second * 10)
	utter, err = client.Post.Create().SetData(fake.Sentence()).SetOwner(adam).Save(ctx)

	if err != nil {
		log.Println(err)
	}

	adam.Update().AddPosts(utter).Save(ctx)

	spew.Dump(steve, corey, adam)

	users, err := client.User.
		Query().
		WithPosts().
		All(ctx)

	if err != nil {
		log.Println(err)
	}
	spew.Dump(users)

	for _, u := range users {
		for _, p := range u.Edges.Posts {
			fmt.Printf("%s: User(%v):%s -> Post(%v): %s\n", p.CreatedAt.String(), u.ID, u.Username, p.ID, p.Data)
		}
	}

}
