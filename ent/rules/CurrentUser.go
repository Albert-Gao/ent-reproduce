package rules

import (
	"context"
	"errors"
)

type CurrentUser struct {
	UserId string

	// the current active tenant
	CurrentTenantId string

	// all tenants the user joins
	TenantIds []string
}

func GetCurrentUserFromContext(c context.Context) (CurrentUser, error) {
	tempUser := c.Value("currentUser")

	if tempUser == nil {
		return CurrentUser{}, errors.New("can not get currentUser from context")
	}

	user, ok := tempUser.(CurrentUser)

	if !ok {
		return CurrentUser{}, errors.New("can not cast the type of currentUser")
	}

	return user, nil
}
