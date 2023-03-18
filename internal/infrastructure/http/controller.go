package http_transport

import (
	"context"

	commands "github.com/mateusmacedo/users-service/internal/application"
)

type SingUpUserRequest struct {
	Name string `json:"name"`
}

type SingUpUserResponse struct {
	ID string `json:"id"`
}

type SingUpUserController struct {
	handler commands.SignUpCommandHandler
}

func NewSingUpUserController(handler commands.SignUpCommandHandler) *SingUpUserController {
	return &SingUpUserController{handler: handler}
}

func (c *SingUpUserController) Handle(ctx context.Context, req *SingUpUserRequest) (*SingUpUserResponse, error) {
	cmd := &commands.SignUpCommand{Name: req.Name}
	user, err := c.handler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &SingUpUserResponse{ID: user.GetID()}, nil
}
