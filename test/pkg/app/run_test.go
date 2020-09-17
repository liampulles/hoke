package app_test

import (
	"testing"

	"github.com/liampulles/hoke/pkg/app"
	"github.com/stretchr/testify/assert"
)

func TestRun_SmokeTest(t *testing.T) {
	// Exercise SUT
	actual := app.Run([]string{"hoke"})

	// Verify results
	assert.Equal(t, 0, actual)
}
