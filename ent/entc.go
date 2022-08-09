//go:build ignore
// +build ignore

package main

import (
	"log"

	_ "enttry/entgen/runtime"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./graphql/schema/ent.graphql"),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.FeatureNames("privacy", "entql", "schema/snapshot"),
		entc.Extensions(ex),
	}

	if err := entc.Generate("./ent/schema", &gen.Config{
		Target:  "./entgen",
		Package: "enttry/entgen",
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
