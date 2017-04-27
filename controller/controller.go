package controller

import (
	"code.cloudfoundry.org/lager"
	"github.com/zjohl/go-simple-database-challenge/cmd"
	"github.com/zjohl/go-simple-database-challenge/database"
	"github.com/zjohl/go-simple-database-challenge/stdio"
)

type Controller struct {
	stdioSocket *stdio.Socket
	parser      *cmd.Parser
	store       *database.Store
	logger      lager.Logger
}

func NewController(stdioSocket *stdio.Socket, parser *cmd.Parser, store *database.Store, logger lager.Logger) *Controller {
	return &Controller{
		stdioSocket: stdioSocket,
		parser:      parser,
		store:       store,
		logger:      logger,
	}
}

func (c *Controller) Start() {
	c.stdioSocket.BufferInput()
	c.stdioSocket.BufferOutput()
	for {
		select {
		case input := <-c.stdioSocket.In:
			command, err := c.parser.Parse(input)
			if err != nil {
				panic(err)
			}

			output, err := command.Execute()
			if err != nil {
				panic(err)
			}

			c.stdioSocket.Out <- output
		}
	}
}
