package graph

import (
	"context"
	"strconv"
)

// Resolver struct to hold in-memory data
type Resolver struct {
	users []*User
}

// Query Resolver
func (r *Resolver) Users(ctx context.Context) ([]*User, error) {
	return r.users, nil
}

// Mutation Resolver
func (r *Resolver) CreateUser(ctx context.Context, name string) (*User, error) {
	user := &User{
		ID:   strconv.Itoa(len(r.users) + 1),
		Name: name,
	}
	r.users = append(r.users, user)
	return user, nil
}
