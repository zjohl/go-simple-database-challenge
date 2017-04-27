package stdio

import (
	"fmt"

	"code.cloudfoundry.org/lager"
)

type Socket struct {
	In     chan string
	Out    chan string
	logger lager.Logger
}

func NewSocket(logger lager.Logger) *Socket {
	return &Socket{
		In:     make(chan string),
		Out:    make(chan string),
		logger: logger,
	}
}

func (s *Socket) BufferInput() {
	go func() {
		var input string
		for {
			fmt.Scanln(&input)
			s.In <- input
		}
	}()
}

func (s *Socket) BufferOutput() {
	go func(passed chan string) {
		for {
			fmt.Println(<-passed)
		}
	}(s.Out)
}
