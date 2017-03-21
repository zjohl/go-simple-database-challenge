package database

import "errors"

type Store struct {
	KeyMap map[string]string
}

func (s *Store) Set(key, value string) string {
	prevValue := s.KeyMap[key]
	s.KeyMap[key] = value
	return prevValue
}

func (s *Store) Get(key string) (string, error) {
	val, ok := s.KeyMap[key]
	if !ok {
		return "", errors.New("no value for provided key")
	}
	return val, nil
}
