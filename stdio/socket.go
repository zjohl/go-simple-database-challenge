package stdio

import (
	"fmt"

	"os"

	"code.cloudfoundry.org/lager"
)

type Socket struct {
	In     chan string
	Out    chan string
	stdin  *os.File
	stdout *os.File
	logger lager.Logger
}

func NewSocket(logger lager.Logger, in *os.File, out *os.File) *Socket {
	return &Socket{
		In:     make(chan string),
		Out:    make(chan string),
		stdin:  in,
		stdout: out,
		logger: logger,
	}
}

func (s *Socket) BufferInput() {
	go func() {
		var input string
		for {
			fmt.Fscan(s.stdin, &input)
			if input != "" {
				s.In <- input
			}
		}
	}()
}

func (s *Socket) BufferOutput() {
	go func(passed chan string) {
		for {
			output := <-passed
			if output != "" {
				fmt.Println(output)
			}
		}
	}(s.Out)
}
