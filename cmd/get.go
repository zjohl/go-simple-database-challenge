package cmd

type Get struct {
	Key string
}

func (s *Get) Execute() (string, error) {
	return "", nil
}
