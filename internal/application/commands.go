package commands

import (
	"context"

	domain "github.com/mateusmacedo/users-service/internal/domain"
)

type SignUpCommand struct {
	Name string
}

type SignUpCommandHandler interface {
	Handle(ctx context.Context, cmd *SignUpCommand) (domain.User, error)
}

type SignUpCommandHandlerImpl struct {
	repository domain.SaveUserRepository
	identity   domain.IdentityService
}

func (h *SignUpCommandHandlerImpl) Handle(ctx context.Context, cmd *SignUpCommand) (domain.User, error) {
	id, err := h.identity.Generate(ctx)
	if err != nil {
		return domain.User{}, err
	}

	user, err := domain.NewUser(id, cmd.Name)
	if err != nil {
		return domain.User{}, err
	}

	user, err = h.repository.Save(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return *user, nil
}

func NewSignUpCommandHandler(repository domain.SaveUserRepository, identity domain.IdentityService) *SignUpCommandHandlerImpl {
	return &SignUpCommandHandlerImpl{repository: repository, identity: identity}
}
