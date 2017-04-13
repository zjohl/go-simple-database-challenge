package cmd

import (
	"errors"
	"strings"
)

type Parser struct{}

type Command interface {
	Execute() (string, error)
}

func (p *Parser) Parse(toParse string) (Command, error) {
	components := strings.Split(toParse, " ")

	switch strings.ToUpper(components[0]) {
	case "SET":
		if len(components) == 3 {
			return &Set{Key: components[1], Value: components[2]}, nil
		} else {
			return nil, errors.New("wrong number of arguements to command")
		}
	case "GET":
		if len(components) == 2 {
			return &Get{Key: components[1]}, nil
		} else {
			return nil, errors.New("wrong number of arguements to command")
		}
	}
	return nil, errors.New("unrecognized command")
}
