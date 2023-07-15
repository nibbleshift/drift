//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("ent.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(ex, entviz.Extension{}),
	}

	features := []gen.Feature{
		gen.FeatureEntQL,
		gen.FeatureUpsert,
	}

	if err := entc.Generate("./ent/schema", &gen.Config{Features: features}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
