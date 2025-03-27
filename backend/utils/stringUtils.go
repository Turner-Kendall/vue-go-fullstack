package utils

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func ProcessUsername(s string) (string, error) {

	if s == "" {
		return "", errors.New("input string cannot be empty")
	}
	if len(s) > 49 {
		return "", errors.New("input string cannot be more than 50 characters")
	}
	if len(s) < 3 {
		return "", errors.New("input string cannot be less than 3 characters")
	}

	return strings.ToLower(s), nil

}

func ProcessPassword(p string) ([]byte, error) {
	if p == "" {
		return nil, errors.New("input string cannot be empty")
	}
	if len(p) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}
	if len(p) > 25 {
		return nil, errors.New("password must be less than 25 characters")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("password could not be generated")
	}

	return passwordHash, nil

}
