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

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Indexes of the Profile.
func (Profile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id", "tenant_id"),
	}
}

// Fields of the Profile.
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
			rules.FilterTenantIdRule("owner_id"),
		},
	}
}

func (Profile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Annotations returns Todo annotations.
//func (Profile) Annotations() []schema.Annotation {
//	return []schema.Annotation{
//		entgql.RelayConnection(),
//		entgql.QueryField(),
//	}
//}
