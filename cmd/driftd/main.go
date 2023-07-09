package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/nibbleshift/drift"
	"github.com/nibbleshift/drift/ent"
)

const defaultPort = "8080"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx := context.Background()
	err := client.Schema.Create(ctx)

	if err != nil {
		log.Println(err)
	}
	srv := handler.NewDefaultServer(drift.NewSchema(client))
	//srv := handler.NewDefaultServer(drift.NewExecutableSchema(drift.Config{Resolvers: &drift.Resolver{}}))

	srv.AddTransport(&transport.Websocket{}) // <---- This is the important part!

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
