package cset

// A very simple example about how to use concurrent-safe SETs (using string as keys) in GO

import (
	"sync"
)

// CMap type def
type CMap struct {
	m map[string]interface{}
	sync.RWMutex
}

// NewMap creates a new CSet struct
func NewMap() *CMap {
	return &CMap{
		m: make(map[string]interface{}),
	}
}

// Set add
func (s *CMap) Set(key string, value interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[key] = value
}

func (s *CMap) Get(key string) interface{} {
	s.RLock()
	defer s.RUnlock()
	if val, ok := s.m[key]; ok {
		return val
	}
	return nil
}

// Remove deletes the specified item from the map
func (s *CMap) Remove(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, key)
}

// Has looks for the existence of an item
func (s *CMap) Has(item string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// Count returns the number of items in a set.
func (s *CMap) Count() int {
	return len(s.m)
}

// Clear removes all items from the set
func (s *CMap) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = make(map[string]interface{})
}

// IsEmpty checks for emptiness
func (s *CMap) IsEmpty() bool {
	return len(s.m) == 0
}

// Keys returns a slice of all items
func (s *CMap) Keys() []string {
	s.RLock()
	defer s.RUnlock()
	var list = make([]string, 0)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
