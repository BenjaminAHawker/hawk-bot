package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Define the Deps type
type Deps struct {
	// Add fields as needed
}

// Define the HelloCommand type
type HelloCommand struct {
	Deps *Deps
}

func NewHelloCommand(deps *Deps) *HelloCommand {
	return &HelloCommand{Deps: deps}
}

func TestNewHelloCommand(t *testing.T) {
	// Mock dependencies
	mockDeps := &Deps{}

	// Call the function
	helloCommand := NewHelloCommand(mockDeps)

	// Assert that the returned HelloCommand is not nil
	assert.NotNil(t, helloCommand)

	// Assert that the Deps field is correctly set
	assert.Equal(t, mockDeps, helloCommand.Deps)
}
