package restful

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	writer := httptest.NewRecorder()
	RootHandler(writer, req)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Contains(t, writer.Body.String(), "Hello from Graph Collab")

	notAllowedMethods := []string{"HEAD", "POST", "PUT", "DELETE", "CONNECT", "PATCH"}
	for _, method := range notAllowedMethods {
		req, err := http.NewRequest(method, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		writer := httptest.NewRecorder()
		RootHandler(writer, req)

		assert.Equal(t, http.StatusMethodNotAllowed, writer.Code)
		assert.Contains(t, writer.Body.String(), "Method not allowed")
	}
}

func TestHealthzHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	writer := httptest.NewRecorder()
	HealthzHandler(writer, req)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Contains(t, writer.Body.String(), "Healthy")

	notAllowedMethods := []string{"HEAD", "POST", "PUT", "DELETE", "CONNECT", "PATCH"}
	for _, method := range notAllowedMethods {
		req, err := http.NewRequest(method, "/healthz", nil)
		if err != nil {
			t.Fatal(err)
		}

		writer := httptest.NewRecorder()
		RootHandler(writer, req)

		assert.Equal(t, http.StatusMethodNotAllowed, writer.Code)
		assert.Contains(t, writer.Body.String(), "Method not allowed")
	}
}
