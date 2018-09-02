package job

import (
	"fmt"
	"sync"
)

type Store interface {
	SetState(string, bool)
	GetState(string) (bool, error)
}

type Service struct {
	Name     string
	CheckJob Job `yaml:"checkJob"`
}

type ServicesStateStore struct {
	mux   sync.Mutex
	state map[string]bool
}

func NewServicesStateStore() *ServicesStateStore {
	return &ServicesStateStore{
		state: map[string]bool{},
	}
}

func (s *ServicesStateStore) SetState(serviceName string, state bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.state[serviceName] = state
}

func (s *ServicesStateStore) GetState(serviceName string) (bool, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	val, ok := s.state[serviceName]
	if !ok {
		return false, fmt.Errorf("Service not found")
	}
	return val, nil
}
