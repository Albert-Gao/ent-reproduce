package tests

import (
	"testing"

	"enttry/ent/rules"
	"enttry/entgen/user"

	"github.com/stretchr/testify/assert"
)

func TestUserShouldBeAbleToGetAllSelfData(t *testing.T) {
	ent := GetEntTestClient(t)
	defer ent.Close()

	user1Ctx := getAuthContext(rules.CurrentUser{
		UserId:          TestUser1Id,
		CurrentTenantId: TestTenant1Id,
		TenantIds:       []string{TestTenant1Id},
	})

	user2Ctx := getAuthContext(rules.CurrentUser{
		UserId:          TestUser2Id,
		CurrentTenantId: TestTenant1Id,
		TenantIds:       []string{TestTenant1Id, TestTenant2Id},
	})

	// user1 -> tenant1 / profile1
	tenant1, user1, profile1 := getTestUser1Data(ent, user1Ctx)

	// user2 -> tenant1 / profile3 && tenant2 / profile2
	tenant2, user2, profile2, profile3 := getTestUser2Data(ent, user2Ctx)

	// user1 queries himself
	qUser1 := ent.User.Query().Where(user.ID(user1.ID)).WithProfiles().WithTenants().OnlyX(user1Ctx)

	// Assert user1 can get all the connections from himself
	assert.Len(t, qUser1.Edges.Tenants, 1)
	assert.Equal(t, qUser1.Edges.Tenants[0].ID, tenant1.ID)
	assert.Len(t, qUser1.Edges.Profiles, 1)
	assert.Equal(t, qUser1.Edges.Profiles[0].ID, profile1.ID)

	// user2 queries himself
	qUser2 := ent.User.Query().Where(user.ID(user2.ID)).WithProfiles().WithTenants().OnlyX(user2Ctx)

	// make sure user2 can get all the connections from himself
	assert.Len(t, qUser2.Edges.Tenants, 2)
	assert.Equal(t, qUser2.Edges.Tenants[0].ID, tenant1.ID)
	assert.Equal(t, qUser2.Edges.Tenants[1].ID, tenant2.ID)
	assert.Len(t, qUser2.Edges.Profiles, 2)
	assert.Equal(t, qUser2.Edges.Profiles[0].ID, profile3.ID)
	assert.Equal(t, qUser2.Edges.Profiles[1].ID, profile2.ID)
}

func TestUserShouldNotBeAbleToGetNonSameTenantData1(t *testing.T) {
	ent := GetEntTestClient(t)
	defer ent.Close()

	user1Ctx := getAuthContext(rules.CurrentUser{
		UserId:          TestUser1Id,
		CurrentTenantId: TestTenant1Id,
		TenantIds:       []string{TestTenant1Id},
	})

	user2Ctx := getAuthContext(rules.CurrentUser{
		UserId:          TestUser2Id,
		CurrentTenantId: TestTenant1Id,
		TenantIds:       []string{TestTenant1Id, TestTenant2Id},
	})

	// user1 -> tenant1 / profile1
	getTestUser1Data(ent, user1Ctx)

	// user2 -> tenant1 / profile3 && tenant2 / profile2
	_, user2, _, profile3 := getTestUser2Data(ent, user2Ctx)

	// user1 -> tenant1 || user2 -> tenant1 and tenant2
	// when user1 tries to get user2,
	// user1 should only get user2.profile (which links to tenant1)
	user2Got1 := ent.User.Query().Where(user.ID(user2.ID)).WithProfiles().OnlyX(user1Ctx)
	assert.Len(t, user2Got1.Edges.Profiles, 1)
	assert.Equal(t, user2Got1.Edges.Profiles[0].ID, profile3.ID)
}

func TestUserShouldNotBeAbleToGetNonSameTenantData2(t *testing.T) {
	ent := GetEntTestClient(t)
	defer ent.Close()

	user1Ctx := getAuthContext(rules.CurrentUser{
		UserId:          TestUser1Id,
		CurrentTenantId: TestTenant1Id,
		TenantIds:       []string{TestTenant1Id},
	})

	user2Ctx := getAuthContext(rules.CurrentUser{
		UserId:          TestUser2Id,
		CurrentTenantId: TestTenant1Id,
		TenantIds:       []string{TestTenant1Id, TestTenant2Id},
	})

	// user1 -> tenant1 / profile1
	tenant1, _, _ := getTestUser1Data(ent, user1Ctx)

	// user2 -> tenant1 / profile3 && tenant2 / profile2
	_, user2, _, _ := getTestUser2Data(ent, user2Ctx)

	// user1 -> tenant1 || user2 -> tenant1 and tenant2
	// when user1 tries to get user2,
	// user1 should only get user2.tenants (which links to tenant1)
	user2Got1 := ent.User.Query().Where(user.ID(user2.ID)).WithTenants().OnlyX(user1Ctx)
	assert.Len(t, user2Got1.Edges.Tenants, 1)
	assert.Equal(t, user2Got1.Edges.Tenants[0].ID, tenant1.ID)
}
