package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"enttry/entgen"
	"enttry/gqlgen"
	"enttry/graphql"
	"enttry/handlers"

	"enttry/entgen/migrate"
	_ "enttry/entgen/runtime"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := entgen.Open("mysql", "root:root@tcp(localhost:3306)/enttry?parseTime=true")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	err1 := client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false),
	)
	if err1 != nil {
		log.Fatalln("failed creating schema resources: %w", err1)
	}
	defer client.Close()

	resolvers := graphql.Resolver{
		Client: client,
	}

	srv := handler.NewDefaultServer(
		gqlgen.NewExecutableSchema(gqlgen.Config{
			Resolvers: &resolvers,
		}),
	)

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		// send this panic somewhere
		log.Print(err)
		debug.PrintStack()
		return errors.New("user message on panic")
	})

	http.Handle("/playground", playground.Handler("Query", "/query"))
	http.Handle("/query", srv)

	i := handlers.InsertUser{
		Ent: client,
	}

	http.Handle("/create", i)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
