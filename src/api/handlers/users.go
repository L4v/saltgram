package handlers

import (
	"fmt"
	"log"
	"net/http"
	"saltgram/protos/users/prusers"
)

type Users struct {
	l  *log.Logger
	uc prusers.UsersClient
}

func NewUsers(l *log.Logger, uc prusers.UsersClient) *Users {
	return &Users{l: l, uc: uc}
}

var ErrorJWSNotFound = fmt.Errorf("jws not found")

func getUserJWS(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) <= 7 {
		return "", ErrorJWSNotFound
	}
	// NOTE(Jovan): Trimming first 7 characters from "Bearer <jws>"
	return authHeader[7:], nil
}
