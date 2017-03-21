package database

import "errors"

type Store struct {
	KeyMap map[string]string
}

const NoValueErrorMessage = "no value for provided key"

func (s *Store) Set(key, value string) string {
	prevValue := s.KeyMap[key]
	s.KeyMap[key] = value
	return prevValue
}

func (s *Store) Get(key string) (string, error) {
	val, ok := s.KeyMap[key]
	if !ok {
		return "", errors.New(NoValueErrorMessage)
	}
	return val, nil
}

func (s *Store) Unset(key string) (string, error) {
	prevValue, ok := s.KeyMap[key]
	if !ok {
		return "", errors.New(NoValueErrorMessage)
	}
	delete(s.KeyMap, key)
	return prevValue, nil
}
