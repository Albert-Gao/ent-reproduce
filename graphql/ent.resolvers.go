package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"enttry/ent/entHelpers"
	"enttry/entgen"
	"enttry/gqlgen"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (entgen.Noder, error) {
	return r.Client.Noder(ctx, id, entgen.WithNodeType(entHelpers.IDToType))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]entgen.Noder, error) {
	return r.Client.Noders(ctx, ids, entgen.WithNodeType(entHelpers.IDToType))
}

// Query returns gqlgen.QueryResolver implementation.
func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
