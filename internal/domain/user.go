package domain

import "errors"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUser(id string, name string) (*User, error) {
	if id == "" {
		return nil, errors.New("empty id field")
	}
	if name == "" {
		return nil, errors.New("empty name feild")
	}
	return &User{
		ID:   id,
		Name: name,
	}, nil
}

func (u *User) UpdateName(name string) {
	u.Name = name
}

func (u *User) SetID(id string) {
	u.ID = id
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetID() string {
	return u.ID
}
