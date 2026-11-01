package main

import (
	"errors"
	"fmt"
	"strings"
)

type UserAgent struct {
	UserAgent string
}

func NewUserAgent(ua string) (UserAgent, error) {
	// separate user agent in parts
	parts := strings.Fields(ua)
	if len(parts) != 2 {
		fmt.Println(parts)
		return UserAgent{}, errors.New("UserAgent format is invalid or doesn't exist")
	}

	// split user agent
	splitUserAgent := strings.Split(parts[1], "\r\n")
	return UserAgent{
		UserAgent: strings.TrimSpace(splitUserAgent[0]),
	}, nil
}
