package counters

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

//Todo test all methods

func TestCreate(t *testing.T) {
	var cases = []struct {
		name        string
		setup       func() *InMemDB
		counterName string
		expectErr   error
	}{
		{
			"Successfully create counter",
			func() *InMemDB {
				return New()
			},
			"unique",
			nil,
		},
		{
			"Error; Counter already exist",
			func() *InMemDB {
				return &InMemDB{
					counters: map[string]int{"not_unique": 0},
				}
			},
			"not_unique",
			errors.New("counter already exist"),
		},
	}
	for _, c := range cases {
		db := c.setup()
		require.Equal(t, c.expectErr, db.Create(c.counterName))
	}
}
