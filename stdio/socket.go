package stdio

import (
	"bufio"
	"os"
)

type Socket struct {
	In chan string
}

func (s *Socket) BufferInput() {
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			s.In <- scanner.Text()
		}
	} ()
}
