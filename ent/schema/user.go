package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),

		field.String("name").Optional(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// User-to-Profile, O2M - O2O
		edge.To("profiles", Profile.Type),

		// User-to-Organization, M2M - M2M
		edge.To("tenants", Tenant.Type).Required(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
