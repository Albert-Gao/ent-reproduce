package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Tenant struct {
	ent.Schema
}

func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),

		field.String("name").MinLen(1).MaxLen(70),
	}
}

func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{
		// User-to-Organization, M2M - M2M
		edge.From("members", User.Type).Ref("tenants").Annotations(
			entgql.RelayConnection(),
		),

		// Organization to Profile, O2M - O2O
		edge.To("memberProfiles", Profile.Type).Annotations(
			entgql.RelayConnection(),
		),
	}
}

func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
