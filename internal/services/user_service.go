package services

import (
	"context"
	"errors"
	"stackoverflow-clone/ent"
	"stackoverflow-clone/ent/user"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
    client *ent.Client
}

func NewUserService(client *ent.Client) *UserService {
    return &UserService{client: client}
}

func (s *UserService) CreateUser(ctx context.Context, username, email, password, displayName string) (*ent.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    return s.client.User.Create().
        SetUsername(username).
        SetEmail(email).
        SetPasswordHash(string(hashedPassword)).
        SetDisplayName(displayName).
        SetCreatedAt(time.Now()).
        Save(ctx)
}

func (s *UserService) AuthenticateUser(ctx context.Context, email, password string) (*ent.User, error) {
    user, err := s.client.User.Query().
        Where(user.EmailEQ(email)).
        Only(ctx)

    if err != nil {
        return nil, err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return nil, errors.New("invalid password")
    }

    return user, nil
}