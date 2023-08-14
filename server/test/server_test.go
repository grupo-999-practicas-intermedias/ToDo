package test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestConnectServer(t *testing.T) {
	// Define the struct
	tests := []struct {
		description  string
		route        string
		method       string
		expected     string
		expectedCode int
	}{
		{
			description:  "get HTTP status 200, when route is found",
			route:        "/",
			method:       "GET",
			expected:     "Hello, World!",
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 404, when route is not found",
			route:        "/not-found",
			method:       "GET",
			expected:     "Cannot GET /not-found",
			expectedCode: 404,
		},
	}

	// define fiber app
	app := fiber.New()

	// create the route with Get
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})
	// iterate through the tests
	for _, tt := range tests {
		// create a request test
		req := httptest.NewRequest(tt.method, tt.route, nil)

		// perform the request test
		resp, _ := app.Test(req, 1)

		// compare the actual response with the expected response
		assert.Equal(t, tt.expectedCode, resp.StatusCode, tt.description)
	}
}
