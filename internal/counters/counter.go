package counters

import (
	"errors"
)

// InMemDB TODO replace with proper percistance db
type InMemDB struct {
	counters map[string]int
}

func New() *InMemDB {
	return &InMemDB{
		counters: make(map[string]int, 0),
	}
}

// Create creates a new counter
func (c *InMemDB) Create(name string) error {
	if _, ok := c.counters[name]; ok {
		return errors.New("counter already exist")
	}

	c.counters[name] = 0

	return nil
}

// Increment increments the associated counter by 1 based on the passed name
func (c *InMemDB) Increment(name string) error {
	if _, ok := c.counters[name]; !ok {
		return errors.New("counter does not exist")
	}

	c.counters[name]++

	return nil
}

// GetOne returns the counter associated with the passed name
func (c *InMemDB) GetOne(name string) (int, error) {
	if _, ok := c.counters[name]; !ok {
		return 0, errors.New("counter does not exist")
	}

	return c.counters[name], nil
}

// GetAll returns all counters
func (c *InMemDB) GetAll() map[string]int {
	return c.counters
}

// TODO resets, deletions
