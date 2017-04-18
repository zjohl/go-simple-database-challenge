package database

import (
	"errors"

	"code.cloudfoundry.org/lager"
)

func NewStore(logger lager.Logger) *Store {
	return &Store{
		keyMap:         make(map[string]string),
		occurrencesMap: make(map[string]int),
		logger:         logger,
	}
}

type Store struct {
	keyMap         map[string]string
	occurrencesMap map[string]int
	logger         lager.Logger
}

const NoValueErrorMessage = "no value for provided key"

func (s *Store) Set(key, value string) string {
	prevValue := s.keyMap[key]
	s.keyMap[key] = value
	s.occurrencesMap[value] = s.occurrencesMap[value] + 1
	return prevValue
}

func (s *Store) Get(key string) (string, error) {
	val, ok := s.keyMap[key]
	if !ok {
		return "", errors.New(NoValueErrorMessage)
	}
	return val, nil
}

func (s *Store) Unset(key string) (string, error) {
	prevValue, ok := s.keyMap[key]
	if !ok {
		return "", errors.New(NoValueErrorMessage)
	}
	delete(s.keyMap, key)
	s.occurrencesMap[prevValue] = s.occurrencesMap[prevValue] - 1
	return prevValue, nil
}

func (s *Store) NumEqualTo(value string) int {
	return s.occurrencesMap[value]
}
