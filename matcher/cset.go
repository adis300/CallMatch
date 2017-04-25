package matcher

// A very simple example about how to use concurrent-safe SETs (using string as keys) in GO

import (
	"fmt"
	"sync"
)

// CSet type def
type CSet struct {
	m map[string]bool
	sync.RWMutex
}

// New creates a new CSet struct
func New() *CSet {
	return &CSet{
		m: make(map[string]bool),
	}
}

/*
func main() {
	// Initialize our Set
	s := New()

	// Add example items
	s.Add("item1")
	s.Add("item1") // duplicate item
	s.Add("item2")
	fmt.Printf("%d items\n", s.Len())

	// Clear all items
	s.Clear()
	if s.IsEmpty() {
		fmt.Printf("0 items\n")
	}

	s.Add("item2")
	s.Add("item3")
	s.Add("item4")

	// Check for existence
	if s.Has("item2") {
		fmt.Println("item2 does exist")
	}

	// Remove some of our items
	s.Remove("item2")
	s.Remove("item4")
	fmt.Println("list of all items:", s.List())
}
*/
// Add add
func (s *CSet) Add(item string) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

// Remove deletes the specified item from the map
func (s *CSet) Remove(item string) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

// Pop returns an item and removes the item from map
func (s *CSet) Pop() string {
	if len(s.m) > 0 {
		s.Lock()
		defer s.Unlock()
		var item string
		for k := range s.m {
			item = k
			break
		}
		delete(s.m, item)
		return item
	}
	return ""
}

// Has looks for the existence of an item
func (s *CSet) Has(item string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// Len returns the number of items in a set.
func (s *CSet) Len() int {
	return len(s.m)
}

// Clear removes all items from the set
func (s *CSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = make(map[string]bool)
}

// IsEmpty checks for emptiness
func (s *CSet) IsEmpty() bool {
	return len(s.m) == 0
}

// List returns a slice of all items
func (s *CSet) List() []string {
	s.RLock()
	defer s.RUnlock()
	var list = make([]string, 0)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
