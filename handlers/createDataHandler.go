package handlers

import (
	"context"
	"fmt"
	"net/http"

	"enttry/entgen"

	"github.com/segmentio/ksuid"
)

type InsertUser struct {
	Ent *entgen.Client
}

func Id(name string) string {
	return name + "_" + ksuid.New().String()
}

func (i InsertUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	tenant1 := i.Ent.Tenant.Create().SetID(Id("tenant")).SetName("tenant 1").SaveX(ctx)

	user1 := i.Ent.User.Create().
		SetID(Id("user")).SetName("user1").
		AddTenantIDs(tenant1.ID).SaveX(ctx)

	user2 := i.Ent.User.Create().
		SetID(Id("user")).SetName("user2").
		AddTenantIDs(tenant1.ID).SaveX(ctx)

	user3 := i.Ent.User.Create().
		SetID(Id("user")).SetName("user3").
		AddTenantIDs(tenant1.ID).SaveX(ctx)

	profile1 := i.Ent.Profile.Create().
		SetID(Id("profile")).SetNickName("profile1").
		SetTenantID(tenant1.ID).SetOwnerID(user1.ID).SaveX(ctx)

	profile2 := i.Ent.Profile.Create().
		SetID(Id("profile")).SetNickName("profile2").
		SetTenantID(tenant1.ID).SetOwnerID(user1.ID).SaveX(ctx)

	profile3 := i.Ent.Profile.Create().
		SetID(Id("profile")).SetNickName("profile3").
		SetTenantID(tenant1.ID).SetOwnerID(user1.ID).SaveX(ctx)

	profile4 := i.Ent.Profile.Create().
		SetID(Id("profile")).SetNickName("profile4").
		SetTenantID(tenant1.ID).SetOwnerID(user2.ID).SaveX(ctx)

	profile5 := i.Ent.Profile.Create().
		SetID(Id("profile")).SetNickName("profile5").
		SetTenantID(tenant1.ID).SetOwnerID(user3.ID).SaveX(ctx)

	fmt.Print(profile1, profile2, profile3, profile4, profile5)
}
