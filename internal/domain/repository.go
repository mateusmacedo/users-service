package domain

import "context"

type IdentityService interface {
	Generate(ctx context.Context) (string, error)
}

type SaveUserRepository interface {
	Save(ctx context.Context, user *User) (*User, error)
}

type GetUserRepository interface {
	Get(ctx context.Context, id string) (*User, error)
}

type DeleteUserRepository interface {
	Delete(ctx context.Context, id string) error
}

type ListUserRepository interface {
	List(ctx context.Context, filter map[string]interface{}, limit int, offset int) ([]*User, error)
}
