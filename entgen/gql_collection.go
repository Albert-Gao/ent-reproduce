// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"database/sql/driver"
	"enttry/entgen/tenant"
	"enttry/entgen/user"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pr *ProfileQuery) CollectFields(ctx context.Context, satisfies ...string) (*ProfileQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pr, nil
	}
	if err := pr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pr, nil
}

func (pr *ProfileQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "owner":
			var (
				path  = append(path, field.Name)
				query = &UserQuery{config: pr.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pr.withOwner = query
		case "tenant":
			var (
				path  = append(path, field.Name)
				query = &TenantQuery{config: pr.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pr.withTenant = query
		}
	}
	return nil
}

type profilePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ProfilePaginateOption
}

func newProfilePaginateArgs(rv map[string]interface{}) *profilePaginateArgs {
	args := &profilePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*ProfileWhereInput); ok {
		args.opts = append(args.opts, WithProfileFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TenantQuery) CollectFields(ctx context.Context, satisfies ...string) (*TenantQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TenantQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "members":
			var (
				path  = append(path, field.Name)
				query = &UserQuery{config: t.config}
			)
			args := newUserPaginateArgs(fieldArgs(ctx, new(UserWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newUserPager(args.opts)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			if !hasCollectedField(ctx, append(path, edgesField)...) || args.first != nil && *args.first == 0 || args.last != nil && *args.last == 0 {
				if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
					query := query.Clone()
					t.loadTotal = append(t.loadTotal, func(ctx context.Context, nodes []*Tenant) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID string `sql:"tenant_id"`
							Count  int    `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							joinT := sql.Table(tenant.MembersTable)
							s.Join(joinT).On(s.C(user.FieldID), joinT.C(tenant.MembersPrimaryKey[0]))
							s.Where(sql.InValues(joinT.C(tenant.MembersPrimaryKey[1]), ids...))
							s.Select(joinT.C(tenant.MembersPrimaryKey[1]), sql.Count("*"))
							s.GroupBy(joinT.C(tenant.MembersPrimaryKey[1]))
						})
						if err := query.Select().Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[string]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							nodes[i].Edges.totalCount[0] = &n
						}
						return nil
					})
				}
				continue
			}
			if (args.after != nil || args.first != nil || args.before != nil || args.last != nil) && hasCollectedField(ctx, append(path, totalCountField)...) {
				query := query.Clone()
				t.loadTotal = append(t.loadTotal, func(ctx context.Context, nodes []*Tenant) error {
					ids := make([]driver.Value, len(nodes))
					for i := range nodes {
						ids[i] = nodes[i].ID
					}
					var v []struct {
						NodeID string `sql:"tenant_id"`
						Count  int    `sql:"count"`
					}
					query.Where(func(s *sql.Selector) {
						joinT := sql.Table(tenant.MembersTable)
						s.Join(joinT).On(s.C(user.FieldID), joinT.C(tenant.MembersPrimaryKey[0]))
						s.Where(sql.InValues(joinT.C(tenant.MembersPrimaryKey[1]), ids...))
						s.Select(joinT.C(tenant.MembersPrimaryKey[1]), sql.Count("*"))
						s.GroupBy(joinT.C(tenant.MembersPrimaryKey[1]))
					})
					if err := query.Select().Scan(ctx, &v); err != nil {
						return err
					}
					m := make(map[string]int, len(v))
					for i := range v {
						m[v[i].NodeID] = v[i].Count
					}
					for i := range nodes {
						n := m[nodes[i].ID]
						nodes[i].Edges.totalCount[0] = &n
					}
					return nil
				})
			} else {
				t.loadTotal = append(t.loadTotal, func(_ context.Context, nodes []*Tenant) error {
					for i := range nodes {
						n := len(nodes[i].Edges.Members)
						nodes[i].Edges.totalCount[0] = &n
					}
					return nil
				})
			}
			query = pager.applyCursors(query, args.after, args.before)
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(tenant.MembersPrimaryKey[1], limit, pager.orderExpr(args.last != nil))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query, args.last != nil)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			t.withMembers = query
		case "memberprofiles", "memberProfiles":
			var (
				path  = append(path, field.Name)
				query = &ProfileQuery{config: t.config}
			)
			args := newProfilePaginateArgs(fieldArgs(ctx, new(ProfileWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newProfilePager(args.opts)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			if !hasCollectedField(ctx, append(path, edgesField)...) || args.first != nil && *args.first == 0 || args.last != nil && *args.last == 0 {
				if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
					query := query.Clone()
					t.loadTotal = append(t.loadTotal, func(ctx context.Context, nodes []*Tenant) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID string `sql:"tenant_id"`
							Count  int    `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(tenant.MemberProfilesColumn, ids...))
						})
						if err := query.GroupBy(tenant.MemberProfilesColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[string]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							nodes[i].Edges.totalCount[1] = &n
						}
						return nil
					})
				}
				continue
			}
			if (args.after != nil || args.first != nil || args.before != nil || args.last != nil) && hasCollectedField(ctx, append(path, totalCountField)...) {
				query := query.Clone()
				t.loadTotal = append(t.loadTotal, func(ctx context.Context, nodes []*Tenant) error {
					ids := make([]driver.Value, len(nodes))
					for i := range nodes {
						ids[i] = nodes[i].ID
					}
					var v []struct {
						NodeID string `sql:"tenant_id"`
						Count  int    `sql:"count"`
					}
					query.Where(func(s *sql.Selector) {
						s.Where(sql.InValues(tenant.MemberProfilesColumn, ids...))
					})
					if err := query.GroupBy(tenant.MemberProfilesColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
						return err
					}
					m := make(map[string]int, len(v))
					for i := range v {
						m[v[i].NodeID] = v[i].Count
					}
					for i := range nodes {
						n := m[nodes[i].ID]
						nodes[i].Edges.totalCount[1] = &n
					}
					return nil
				})
			} else {
				t.loadTotal = append(t.loadTotal, func(_ context.Context, nodes []*Tenant) error {
					for i := range nodes {
						n := len(nodes[i].Edges.MemberProfiles)
						nodes[i].Edges.totalCount[1] = &n
					}
					return nil
				})
			}
			query = pager.applyCursors(query, args.after, args.before)
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(tenant.MemberProfilesColumn, limit, pager.orderExpr(args.last != nil))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query, args.last != nil)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			t.withMemberProfiles = query
		}
	}
	return nil
}

type tenantPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TenantPaginateOption
}

func newTenantPaginateArgs(rv map[string]interface{}) *tenantPaginateArgs {
	args := &tenantPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*TenantWhereInput); ok {
		args.opts = append(args.opts, WithTenantFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "profiles":
			var (
				path  = append(path, field.Name)
				query = &ProfileQuery{config: u.config}
			)
			args := newProfilePaginateArgs(fieldArgs(ctx, new(ProfileWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newProfilePager(args.opts)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			if !hasCollectedField(ctx, append(path, edgesField)...) || args.first != nil && *args.first == 0 || args.last != nil && *args.last == 0 {
				if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
					query := query.Clone()
					u.loadTotal = append(u.loadTotal, func(ctx context.Context, nodes []*User) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID string `sql:"owner_id"`
							Count  int    `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(user.ProfilesColumn, ids...))
						})
						if err := query.GroupBy(user.ProfilesColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[string]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							nodes[i].Edges.totalCount[0] = &n
						}
						return nil
					})
				}
				continue
			}
			if (args.after != nil || args.first != nil || args.before != nil || args.last != nil) && hasCollectedField(ctx, append(path, totalCountField)...) {
				query := query.Clone()
				u.loadTotal = append(u.loadTotal, func(ctx context.Context, nodes []*User) error {
					ids := make([]driver.Value, len(nodes))
					for i := range nodes {
						ids[i] = nodes[i].ID
					}
					var v []struct {
						NodeID string `sql:"owner_id"`
						Count  int    `sql:"count"`
					}
					query.Where(func(s *sql.Selector) {
						s.Where(sql.InValues(user.ProfilesColumn, ids...))
					})
					if err := query.GroupBy(user.ProfilesColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
						return err
					}
					m := make(map[string]int, len(v))
					for i := range v {
						m[v[i].NodeID] = v[i].Count
					}
					for i := range nodes {
						n := m[nodes[i].ID]
						nodes[i].Edges.totalCount[0] = &n
					}
					return nil
				})
			} else {
				u.loadTotal = append(u.loadTotal, func(_ context.Context, nodes []*User) error {
					for i := range nodes {
						n := len(nodes[i].Edges.Profiles)
						nodes[i].Edges.totalCount[0] = &n
					}
					return nil
				})
			}
			query = pager.applyCursors(query, args.after, args.before)
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(user.ProfilesColumn, limit, pager.orderExpr(args.last != nil))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query, args.last != nil)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			u.withProfiles = query
		case "tenants":
			var (
				path  = append(path, field.Name)
				query = &TenantQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.withTenants = query
		}
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]interface{}) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput interface{}, path ...string) map[string]interface{} {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	for _, name := range path {
		var field *graphql.CollectedField
		for _, f := range graphql.CollectFields(oc, fc.Field.Selections, nil) {
			if f.Name == name {
				field = &f
				break
			}
		}
		if field == nil {
			return nil
		}
		cf, err := fc.Child(ctx, *field)
		if err != nil {
			args := field.ArgumentMap(oc.Variables)
			return unmarshalArgs(ctx, whereInput, args)
		}
		fc = cf
	}
	return fc.Args
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput interface{}, args map[string]interface{}) map[string]interface{} {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = &c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}