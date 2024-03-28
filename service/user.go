package service

import (
	"context"
	"fmt"

	"github.com/shplume/zhulong/ent"
	"github.com/shplume/zhulong/ent/user"
)

func CreateUser(account, password string, role int) error {
	ctx := context.Background()
	if _, err := client.User.Create().SetAccount(account).SetPassword(password).SetRole(role).Save(ctx); err != nil {
		return fmt.Errorf("create user error: %w", err)
	}

	return nil
}

func QueryUser(account string) (*ent.User, error) {
	ctx := context.Background()
	user, err := client.User.Query().Where(user.Account(account)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("user query error: %w", err)
	}

	return user, nil
}
