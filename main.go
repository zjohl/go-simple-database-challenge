package main

import (
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/zjohl/go-simple-database-challenge/cmd"
	"github.com/zjohl/go-simple-database-challenge/controller"
	"github.com/zjohl/go-simple-database-challenge/database"
	"github.com/zjohl/go-simple-database-challenge/stdio"
)

func main() {

	logger := lager.NewLogger("simple-db")

	stdioSocket := stdio.NewSocket(logger)
	database := database.NewStore(logger)
	parser := cmd.NewParser(logger)
	controller := controller.NewController(stdioSocket, parser, database, logger)

	controller.Start()

	for {
		time.Sleep(time.Second)
	}
}
