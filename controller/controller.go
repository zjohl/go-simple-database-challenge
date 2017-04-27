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

}
