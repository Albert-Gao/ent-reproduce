package rules

import (
	"context"
	"errors"
)

type CurrentUser struct {
	UserId          string
	CurrentTenantId string
	TenantIds       []string
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
