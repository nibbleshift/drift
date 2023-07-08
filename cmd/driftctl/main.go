package main

import (
	"context"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/icrowley/fake"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nibbleshift/drift/ent"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
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

	for i := 0; i < 10; i++ {

		utter, err := client.Utter.Create().SetData(fake.Sentence()).SetOwner(steve).Save(ctx)

		if err != nil {
			log.Println(err)
		}

		steve.Update().AddUtters(utter).Save(ctx)
		spew.Dump(steve, corey, adam)

	}
	spew.Dump(steve, corey, adam)

	users, err := client.User.
		Query().
		WithUtters().
		All(ctx)

	if err != nil {
		log.Println(err)
	}
	spew.Dump(users)

	for _, u := range users {
		for _, p := range u.Edges.Utters {
			fmt.Printf("%s: User(%v):%s -> Utter(%v): %s\n", p.CreatedAt.String(), u.ID, u.Username, p.ID, p.Data)
		}
	}

}
