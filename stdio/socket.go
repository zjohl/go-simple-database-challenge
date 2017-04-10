package stdio

import "fmt"

type Socket struct {
	In  chan string
	Out chan string
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
