package schema

import (
	"enttry/ent/rules"
	"enttry/entgen/privacy"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type Profile struct {
	ent.Schema
}

func (Profile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id", "tenant_id"),
	}
}

func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),

		field.String("nick_name").MaxLen(70),

		field.String("owner_id"),
		field.String("tenant_id"),
	}
}

func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		// User-to-Profile, O2M - O2O
		edge.From("owner", User.Type).Ref("profiles").Field("owner_id").Required().Unique(),

		// Organization to Profile, O2M - O2O
		edge.From("tenant", Tenant.Type).Ref("memberProfiles").Field("tenant_id").Required().Unique(),
	}
}

func (Profile) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rules.FilterProfileRule(),
		},
	}
}

func (Profile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
