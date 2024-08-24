package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func testRouter() *echo.Echo {
	e := echo.New()

	Router(e.Group("/user"))

	return e
}

func TestHelloHandler(t *testing.T) {
	router := testRouter()

	req := httptest.NewRequest("GET", "/user", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	// assert.Equal(t, "Hello, Echo World!!", rec.Body.String())
}

func TestUserPost(t *testing.T) {
	router := testRouter()

	req := httptest.NewRequest("POST", "/user", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, Echo World!!", rec.Body.String())
}
