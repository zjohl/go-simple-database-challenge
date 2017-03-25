package cmd

import (
	"errors"
	"strings"
)

type Parser struct {}

type Command interface {
	Execute() (string, error)
}

func (p *Parser) Parse(toParse string) (Command, error){
	components := strings.Split(toParse, " ")

	switch components[0] {
	case "SET":
		if len(components) == 3 {
			return &Set{Key: components[1], Value: components[2]}, nil
		}
	}
	return nil, errors.New("Unrecognized command")
}
