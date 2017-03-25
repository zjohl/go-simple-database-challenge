package cmd

type Set struct {
	Key string
	Value string
}

func (s *Set ) Execute() (string, error) {
	return "", nil
}
