//go:build ignore
// +build ignore

package main

import (
	"log"

	_ "enttry/entgen/runtime"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	opts := []entc.Option{
		entc.FeatureNames("privacy", "entql"),
	}

	if err := entc.Generate("./ent/schema", &gen.Config{
		Target:  "./entgen",
		Package: "enttry/entgen",
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
