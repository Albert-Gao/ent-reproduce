package tests

import (
	"context"
	"testing"

	"enttry/ent/rules"
	"enttry/entgen"
	"enttry/entgen/enttest"
	"enttry/entgen/migrate"

	_ "github.com/mattn/go-sqlite3"
)

func getEntTestOptions(t *testing.T) []enttest.Option {
	return []enttest.Option{
		enttest.WithOptions(entgen.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	}
}

func GetEntTestClient(t *testing.T) *entgen.Client {
	t.Helper()
	opts := getEntTestOptions(t)

	ent := enttest.Open(
		t,
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
		opts...,
	)

	return ent
}

func getAuthContext(currentUser rules.CurrentUser) context.Context {
	authCtx := context.Background()
	authCtx = context.WithValue(authCtx, "currentUser", currentUser)
	return authCtx
}

const (
	TestUser1Id    = "user_test_1"
	TestUser2Id    = "user_test_2"
	TestTenant1Id  = "tenant_test_1"
	TestTenant2Id  = "tenant_test_2"
	TestProfile1Id = "profile_test_1"
	TestProfile2Id = "profile_test_2"
	TestProfile3Id = "profile_test_3"
)

// user1 -> tenant1 / profile1
func getTestUser1Data(ent *entgen.Client, ctx context.Context) (*entgen.Tenant, *entgen.User, *entgen.Profile) {
	t1 := ent.Tenant.Create().SetID(TestTenant1Id).SetName("testTenant1").SaveX(ctx)
	u1 := ent.User.Create().SetID(TestUser1Id).AddTenants(t1).SaveX(ctx)
	p1 := ent.Profile.Create().SetID(TestProfile1Id).SetNickName("p1").SetOwner(u1).SetTenant(t1).SaveX(ctx)

	return t1, u1, p1
}

// user2 -> tenant1 / profile3 && tenant2 / profile2
func getTestUser2Data(ent *entgen.Client, ctx context.Context) (*entgen.Tenant, *entgen.User, *entgen.Profile, *entgen.Profile) {
	t2 := ent.Tenant.Create().SetID(TestTenant2Id).SetName("testTenant2").SaveX(ctx)
	u2 := ent.User.Create().SetID(TestUser2Id).AddTenantIDs(TestTenant1Id, TestTenant2Id).SaveX(ctx)
	p2 := ent.Profile.Create().SetID(TestProfile2Id).SetNickName("p2").SetOwner(u2).SetTenant(t2).SaveX(ctx)
	p3 := ent.Profile.Create().SetID(TestProfile3Id).SetNickName("p3").SetOwner(u2).SetTenantID(TestTenant1Id).SaveX(ctx)

	return t2, u2, p2, p3
}
