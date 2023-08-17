package test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"server/api-go-test/model"
	"server/api-go-test/routes"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {

	// define the struct
	tests := []struct {
		description  string
		route        string
		method       string
		expectedCode int
		body         []byte
	}{
		{
			description:  "get HTTP status 201, the todo was created",
			route:        "/todos",
			method:       "POST",
			expectedCode: 201,
			body:         []byte(`{"id":"3","title":"test","description":"test","completed":false}`),
		},
		{
			description:  "get HTTP status 400, the json is invalid",
			route:        "/todos",
			method:       "POST",
			expectedCode: 400,
			body:         []byte(`{"id":"1","title":"test","description":"test","completed":"false"}`),
		},
		{
			description:  "get HTTP status 500, todo was not created",
			route:        "/todos",
			method:       "POST",
			expectedCode: 500,
			body:         []byte(`{"id":"3","title":"test","description":"test","COMPLETED":false}`),
		},
	}

	// create fiber app
	app := fiber.New()

	// create the redis client mock
	client, mock := redismock.NewClientMock()
	defer client.Close() // close the mock database connection after test finishes

	// create the route with Post
	app.Post("/todos", routes.CreateTodo(client))

	// iterate through the tests
	for _, tt := range tests {

		if tt.expectedCode == 201 {
			var todo model.Todo
			json.Unmarshal(tt.body, &todo)
			todoJSON, _ := json.Marshal(todo)
			mock.ExpectSet(todo.ID, todoJSON, 0).SetVal("OK")
		}

		// create the request
		req := httptest.NewRequest(tt.method, tt.route, bytes.NewReader(tt.body))

		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, 1)

		assert.Equal(t, tt.expectedCode, resp.StatusCode, tt.description)

	}
}

// TestUpdateTodo is a function that tests the UpdateTodo function
func TestUpdateTodo(t *testing.T) {

	// define the struct
	tests := []struct {
		description  string
		route        string
		method       string
		expectedCode int
		body         []byte
	}{
		{
			description:  "get HTTP status 200, the todo was updated",
			route:        "/todos/1",
			method:       "PUT",
			expectedCode: 200,
			body:         []byte(`{"id":"1","title":"test","description":"test","completed":false}`),
		},
		{
			description:  "get HTTP status 400, the json is invalid",
			route:        "/todos/1",
			method:       "PUT",
			expectedCode: 400,
			body:         []byte(`{"id":"1","title":"test","description":"test","completed":"false"}`),
		},
		{
			description:  "get HTTP status 500, todo was not updated",
			route:        "/todos/1",
			method:       "PUT",
			expectedCode: 500,
			body:         []byte(`{"id":"1","title":"test","description":"test","COMPLETED":false}`),
		},
	}

	// create fiber app
	app := fiber.New()

	// create the redis client mock
	client, mock := redismock.NewClientMock()
	defer client.Close() // close the mock database connection after test finishes

	// create the route with Put
	app.Put("/todos/:id", routes.UpdateTodo(client))

	// iterate through the tests
	for _, tt := range tests {

		if tt.expectedCode == 200 {
			var todo model.Todo
			json.Unmarshal(tt.body, &todo)
			todoJSON, _ := json.Marshal(todo)
			mock.ExpectSet(todo.ID, todoJSON, 0).SetVal("OK")
		}

		// create the request
		req := httptest.NewRequest(tt.method, tt.route, bytes.NewReader(tt.body))

		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, 1)

		assert.Equal(t, tt.expectedCode, resp.StatusCode, tt.description)

	}
}

// TestDeleteTodo is a function that tests the DeleteTodo function
func TestDeleteTodo(t *testing.T) {

	// define the struct
	tests := []struct {
		description  string
		route        string
		param        string
		method       string
		expectedCode int
	}{
		{
			description:  "get HTTP status 200, the todo was deleted",
			route:        "/todos/",
			param:        "1",
			method:       "DELETE",
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 500, todo was not deleted",
			route:        "/todos/",
			param:        "",
			method:       "DELETE",
			expectedCode: 404,
		},
		{
			description:  "get HTTP status 200, the todo was deleted",
			route:        "/todos/",
			param:        "1",
			method:       "DELETE",
			expectedCode: 500,
		},
	}

	// create fiber app
	app := fiber.New()

	// create the redis client mock
	client, mock := redismock.NewClientMock()
	defer mock.ClearExpect()
	defer client.Close() // close the mock database connection after test finishes

	// create the route with Delete
	app.Delete("/todos/:id", routes.DeleteTodo(client))

	// iterate through the tests
	for _, tt := range tests {

		if tt.expectedCode == 200 {
			mock.ExpectDel(tt.param).SetVal(1)
		}

		// create the request
		req := httptest.NewRequest(tt.method, tt.route+tt.param, nil)

		resp, _ := app.Test(req, 1)

		assert.Equal(t, tt.expectedCode, resp.StatusCode, tt.description)

	}
}
