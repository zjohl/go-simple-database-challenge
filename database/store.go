package database

import (
	"errors"
)

func NewStore() *store {
	return &store{
		KeyMap:         make(map[string]string),
		OccurrencesMap: make(map[string]int),
	}
}

type store struct {
	KeyMap         map[string]string
	OccurrencesMap map[string]int
}

const NoValueErrorMessage = "no value for provided key"

func (s *store) Set(key, value string) string {
	prevValue := s.KeyMap[key]
	s.KeyMap[key] = value
	s.OccurrencesMap[value] = s.OccurrencesMap[value] + 1
	return prevValue
}

func (s *store) Get(key string) (string, error) {
	val, ok := s.KeyMap[key]
	if !ok {
		return "", errors.New(NoValueErrorMessage)
	}
	return val, nil
}

func (s *store) Unset(key string) (string, error) {
	prevValue, ok := s.KeyMap[key]
	if !ok {
		return "", errors.New(NoValueErrorMessage)
	}
	delete(s.KeyMap, key)
	s.OccurrencesMap[prevValue] = s.OccurrencesMap[prevValue] - 1
	return prevValue, nil
}

func (s *store) NumEqualTo(value string) int {
	return s.OccurrencesMap[value]
}
