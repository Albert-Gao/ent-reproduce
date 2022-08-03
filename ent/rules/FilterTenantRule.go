package rules

import (
	"context"

	"enttry/entgen/privacy"

	"entgo.io/ent/entql"
)

func FilterTenantRule() privacy.QueryMutationRule {
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

		ids := stringSliceToInterfaceSlice(currentUser.TenantIds)

		tf.Where(entql.FieldIn("id", ids...))

		// Skip to the next rules (equivalent to return nil).
		return privacy.Skip
	})
}
