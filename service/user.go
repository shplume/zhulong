package service

import (
	"context"
	"fmt"
)

func CreateUser(account, password string, role int) error {
	ctx := context.Background()
	if _, err := client.User.Create().SetAccount(account).SetPassword(password).SetRole(role).Save(ctx); err != nil {
		return fmt.Errorf("failed creating user: %w", err)
	}

	return nil
}
