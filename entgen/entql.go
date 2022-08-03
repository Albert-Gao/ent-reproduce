// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"enttry/entgen/predicate"
	"enttry/entgen/profile"
	"enttry/entgen/tenant"
	"enttry/entgen/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 3)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: profile.FieldID,
			},
		},
		Type: "Profile",
		Fields: map[string]*sqlgraph.FieldSpec{
			profile.FieldCreateTime: {Type: field.TypeTime, Column: profile.FieldCreateTime},
			profile.FieldUpdateTime: {Type: field.TypeTime, Column: profile.FieldUpdateTime},
			profile.FieldNickName:   {Type: field.TypeString, Column: profile.FieldNickName},
			profile.FieldOwnerID:    {Type: field.TypeString, Column: profile.FieldOwnerID},
			profile.FieldTenantID:   {Type: field.TypeString, Column: profile.FieldTenantID},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   tenant.Table,
			Columns: tenant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tenant.FieldID,
			},
		},
		Type: "Tenant",
		Fields: map[string]*sqlgraph.FieldSpec{
			tenant.FieldCreateTime: {Type: field.TypeTime, Column: tenant.FieldCreateTime},
			tenant.FieldUpdateTime: {Type: field.TypeTime, Column: tenant.FieldUpdateTime},
			tenant.FieldName:       {Type: field.TypeString, Column: tenant.FieldName},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldCreateTime: {Type: field.TypeTime, Column: user.FieldCreateTime},
			user.FieldUpdateTime: {Type: field.TypeTime, Column: user.FieldUpdateTime},
			user.FieldName:       {Type: field.TypeString, Column: user.FieldName},
		},
	}
	graph.MustAddE(
		"owner",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profile.OwnerTable,
			Columns: []string{profile.OwnerColumn},
			Bidi:    false,
		},
		"Profile",
		"User",
	)
	graph.MustAddE(
		"tenant",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profile.TenantTable,
			Columns: []string{profile.TenantColumn},
			Bidi:    false,
		},
		"Profile",
		"Tenant",
	)
	graph.MustAddE(
		"members",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tenant.MembersTable,
			Columns: tenant.MembersPrimaryKey,
			Bidi:    false,
		},
		"Tenant",
		"User",
	)
	graph.MustAddE(
		"memberProfiles",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tenant.MemberProfilesTable,
			Columns: []string{tenant.MemberProfilesColumn},
			Bidi:    false,
		},
		"Tenant",
		"Profile",
	)
	graph.MustAddE(
		"profiles",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProfilesTable,
			Columns: []string{user.ProfilesColumn},
			Bidi:    false,
		},
		"User",
		"Profile",
	)
	graph.MustAddE(
		"tenants",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TenantsTable,
			Columns: user.TenantsPrimaryKey,
			Bidi:    false,
		},
		"User",
		"Tenant",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (pq *ProfileQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ProfileQuery builder.
func (pq *ProfileQuery) Filter() *ProfileFilter {
	return &ProfileFilter{config: pq.config, predicateAdder: pq}
}

// addPredicate implements the predicateAdder interface.
func (m *ProfileMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ProfileMutation builder.
func (m *ProfileMutation) Filter() *ProfileFilter {
	return &ProfileFilter{config: m.config, predicateAdder: m}
}

// ProfileFilter provides a generic filtering capability at runtime for ProfileQuery.
type ProfileFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ProfileFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *ProfileFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(profile.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *ProfileFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(profile.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *ProfileFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(profile.FieldUpdateTime))
}

// WhereNickName applies the entql string predicate on the nick_name field.
func (f *ProfileFilter) WhereNickName(p entql.StringP) {
	f.Where(p.Field(profile.FieldNickName))
}

// WhereOwnerID applies the entql string predicate on the owner_id field.
func (f *ProfileFilter) WhereOwnerID(p entql.StringP) {
	f.Where(p.Field(profile.FieldOwnerID))
}

// WhereTenantID applies the entql string predicate on the tenant_id field.
func (f *ProfileFilter) WhereTenantID(p entql.StringP) {
	f.Where(p.Field(profile.FieldTenantID))
}

// WhereHasOwner applies a predicate to check if query has an edge owner.
func (f *ProfileFilter) WhereHasOwner() {
	f.Where(entql.HasEdge("owner"))
}

// WhereHasOwnerWith applies a predicate to check if query has an edge owner with a given conditions (other predicates).
func (f *ProfileFilter) WhereHasOwnerWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("owner", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasTenant applies a predicate to check if query has an edge tenant.
func (f *ProfileFilter) WhereHasTenant() {
	f.Where(entql.HasEdge("tenant"))
}

// WhereHasTenantWith applies a predicate to check if query has an edge tenant with a given conditions (other predicates).
func (f *ProfileFilter) WhereHasTenantWith(preds ...predicate.Tenant) {
	f.Where(entql.HasEdgeWith("tenant", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (tq *TenantQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TenantQuery builder.
func (tq *TenantQuery) Filter() *TenantFilter {
	return &TenantFilter{config: tq.config, predicateAdder: tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TenantMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TenantMutation builder.
func (m *TenantMutation) Filter() *TenantFilter {
	return &TenantFilter{config: m.config, predicateAdder: m}
}

// TenantFilter provides a generic filtering capability at runtime for TenantQuery.
type TenantFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TenantFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *TenantFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(tenant.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *TenantFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(tenant.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *TenantFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(tenant.FieldUpdateTime))
}

// WhereName applies the entql string predicate on the name field.
func (f *TenantFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(tenant.FieldName))
}

// WhereHasMembers applies a predicate to check if query has an edge members.
func (f *TenantFilter) WhereHasMembers() {
	f.Where(entql.HasEdge("members"))
}

// WhereHasMembersWith applies a predicate to check if query has an edge members with a given conditions (other predicates).
func (f *TenantFilter) WhereHasMembersWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("members", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasMemberProfiles applies a predicate to check if query has an edge memberProfiles.
func (f *TenantFilter) WhereHasMemberProfiles() {
	f.Where(entql.HasEdge("memberProfiles"))
}

// WhereHasMemberProfilesWith applies a predicate to check if query has an edge memberProfiles with a given conditions (other predicates).
func (f *TenantFilter) WhereHasMemberProfilesWith(preds ...predicate.Profile) {
	f.Where(entql.HasEdgeWith("memberProfiles", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{config: uq.config, predicateAdder: uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{config: m.config, predicateAdder: m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *UserFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(user.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *UserFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(user.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *UserFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(user.FieldUpdateTime))
}

// WhereName applies the entql string predicate on the name field.
func (f *UserFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(user.FieldName))
}

// WhereHasProfiles applies a predicate to check if query has an edge profiles.
func (f *UserFilter) WhereHasProfiles() {
	f.Where(entql.HasEdge("profiles"))
}

// WhereHasProfilesWith applies a predicate to check if query has an edge profiles with a given conditions (other predicates).
func (f *UserFilter) WhereHasProfilesWith(preds ...predicate.Profile) {
	f.Where(entql.HasEdgeWith("profiles", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasTenants applies a predicate to check if query has an edge tenants.
func (f *UserFilter) WhereHasTenants() {
	f.Where(entql.HasEdge("tenants"))
}

// WhereHasTenantsWith applies a predicate to check if query has an edge tenants with a given conditions (other predicates).
func (f *UserFilter) WhereHasTenantsWith(preds ...predicate.Tenant) {
	f.Where(entql.HasEdgeWith("tenants", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
