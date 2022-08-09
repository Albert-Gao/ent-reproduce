package entHelpers

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"enttry/entgen/profile"
	"enttry/entgen/tenant"
	"enttry/entgen/user"

	"entgo.io/ent/entc/integration/privacy/ent/task"
)

var prefixMap = map[string]string{
	"user":    user.Table,
	"profile": profile.Table,
	"tenant":  tenant.Table,
	"task":    task.Table,
}

// IDToType maps ID to the underlying table.
func IDToType(_ context.Context, id string) (string, error) {
	if len(id) < 2 {
		return "", errors.New("IDToType: id too short")
	}
	prefix := strings.Split(id, "_")[0]
	typ := prefixMap[prefix]
	if typ == "" {
		return "", fmt.Errorf("IDToType: could not map prefix '%s' to a type", prefix)
	}
	return typ, nil
}
