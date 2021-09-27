package mygrpc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceLastName(t *testing.T) {
	s := NewService()
	ctx := context.Background()

	lastName, err := s.GetLastName(ctx, "Lionel")
	assert.Nil(t, err, "no error expected")
	assert.Equal(t, "Messi", lastName, "this lastName was not expected")

	lastName, err = s.GetLastName(ctx, "someone")
	assert.NotNil(t, err, "error was expected")
	assert.Equal(t, "", lastName, "lastName was not expected")
	assert.Equal(t, "unknown person", err.Error(), "errortext was not expected")
}
