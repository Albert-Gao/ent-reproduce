package rules

import (
	"context"

	"enttry/entgen/privacy"
	"enttry/entgen/profile"

	"entgo.io/ent/entql"
)

func stringSliceToInterfaceSlice(stringSlice []string) []interface{} {
	var list []interface{}

	for _, id := range stringSlice {
		list = append(list, id)
	}

	return list
}

/*
 FilterProfileRule is a query rule that filters out entities that have tenantId which is not in the currentUser.TenantIds.

 for example, when userA is querying userB.profiles, userB.profiles might contain profile which is in another tenant.

 adding this rule would filter userB.profiles to make sure userA can only see userB.profiles which have the same tenantId.

 ownerIdFieldName is the owner of the entity, for profile, it's owner_id, for task, it's creator_id
*/
func FilterProfileRule() privacy.QueryRule {
	// ProfileFilter is an interface to wrap Where()
	// predicate that is used by `Profile` schema.
	type WhereFilter interface {
		Where(p entql.P)
	}

	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		currentUser, err := GetCurrentUserFromContext(ctx)
		if err != nil {
			return privacy.Denyf(err.Error())
		}

		tf, ok := f.(WhereFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}

		tf.Where(entql.Or(
			entql.FieldEQ(profile.FieldOwnerID, currentUser.UserId),
			entql.FieldEQ(profile.FieldTenantID, currentUser.CurrentTenantId),
		))

		// Skip to the next rules (equivalent to return nil).
		return privacy.Skip
	})
}
