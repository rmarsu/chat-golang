package service

import (
	chat "chat"
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "hueiw&fhueg3rw9fh43"

type service struct {
	chat.Repository
	timeout time.Duration
}

func NewService(repository chat.Repository) chat.Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s service) CreateUser(c context.Context, req *chat.CreateUserReq) (*chat.CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := Hash(req.Password)
	if err != nil {
		return nil, err
	}

	u := &chat.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}
	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	res := &chat.CreateUserRes{
		ID:       r.ID,
		Username: r.Username,
		Email:    r.Email,
	}
	return res, nil
}

type MyJWTToken struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s service) Login(c context.Context, req *chat.LoginRequest) (*chat.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &chat.LoginResponse{}, err

	}
	err = Check(req.Password, u.Password)
	if err != nil {
		return &chat.LoginResponse{}, err

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTToken{
		ID:       u.ID,
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	ss, err := token.SignedString([]byte(secretkey))
	if err != nil {
		return &chat.LoginResponse{}, err
	}
	return &chat.LoginResponse{AccessToken: ss, Username: u.Username, ID: u.ID}, nil

}
