package http_test

import (
	"testing"

	"github.com/liampulles/hoke/pkg/http"
	"github.com/stretchr/testify/assert"
)

func TestGetReturns2xx_WhenGivenInvalidUrl_ShouldFail(t *testing.T) {
	// Exercise SUT
	actual := http.GetReturns2xx("not-an-url")

	// Verify results
	assert.Equal(t, false, actual)
}

func TestGetReturns2xx_WhenGiven404Url_ShouldFail(t *testing.T) {
	// Exercise SUT
	actual := http.GetReturns2xx("https://github.com/not-an-url")

	// Verify results
	assert.Equal(t, false, actual)
}

func TestGetReturns2xx_WhenGiven200Url_ShouldPass(t *testing.T) {
	// Exercise SUT
	actual := http.GetReturns2xx("https://github.githubassets.com/images/icons/twitter.png")

	// Verify results
	assert.Equal(t, true, actual)
}
