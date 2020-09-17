package app_test

import (
	"testing"

	"github.com/liampulles/hoke/pkg/app"
	"github.com/stretchr/testify/assert"
)

func TestRun_WhenGivenValidArgs_ShouldPass(t *testing.T) {
	// Exercise SUT
	actual := app.Run([]string{"hoke", "https://github.com"})

	// Verify results
	assert.Equal(t, 0, actual)
}

func TestRun_WhenGivenInvalidArgs_ShouldFail(t *testing.T) {
	// Exercise SUT
	actual := app.Run([]string{"hoke", "-invalid"})

	// Verify results
	assert.Equal(t, 1, actual)
}

func TestRun_WhenMissingUrl_ShouldFail(t *testing.T) {
	// Exercise SUT
	actual := app.Run([]string{"hoke"})

	// Verify results
	assert.Equal(t, 2, actual)
}
