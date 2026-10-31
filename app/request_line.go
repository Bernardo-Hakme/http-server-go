package main

import (
	"errors"
	"strings"
)

type RequestLine struct {
	HTTPMethod    string
	RequestTarget string
	HTTPVersion   string
}

func New(rl string) (RequestLine, error) {
	// separate rl in parts (3 parts method target and version)
	parts := strings.Fields(rl)
	if len(parts) < 3 {
		return RequestLine{}, errors.New("invalid request line")
	}

	// return RequestLine struct with values
	return RequestLine{
		HTTPMethod:    parts[0],
		RequestTarget: parts[1],
		HTTPVersion:   parts[2],
	}, nil
}

func (rl *RequestLine) getParsedTaget() []string {
	// Parse the target (splitting accordingly)
	return strings.Split(rl.RequestTarget, "/")
}
