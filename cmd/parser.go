package cmd

import "errors"

type Parser struct {}

type Command interface {
	Execute() (string, error)
}

func (p *Parser) Parse(toParse string) (*Command, error){
	return nil, errors.New("Unrecognized command")
}