package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"url-shortener/internal/domains"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockURLUsecase simula a interface URLUsecase
type mockURLUsecase struct {
	mock.Mock
}

func (m *mockURLUsecase) Shorten(original string, expireMinutes int) (string, error) {
	args := m.Called(original, expireMinutes)
	return args.String(0), args.Error(1)
}

func (m *mockURLUsecase) Resolve(slug string) (*domains.URL, error) {
	args := m.Called(slug)
	if url, ok := args.Get(0).(*domains.URL); ok {
		return url, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestShorten(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUC := new(mockURLUsecase)
	controller := NewURLController(mockUC)

	router := gin.Default()
	router.POST("/api/shorten", controller.Shorten)

	body := `{"url": "https://github.com/stretchr/testify", "expire_in_minutes": 1}`
	mockUC.On("Shorten", "https://github.com/stretchr/testify", 1).Return("abc123", nil)

	req := httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Host = "localhost:8080"
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status code 200 OK")
	assert.Contains(t, resp.Body.String(), "abc123", "Expected response to contain the shortened URL slug")
	mockUC.AssertExpectations(t)
}

func TestRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUC := new(mockURLUsecase)
	controller := NewURLController(mockUC)

	router := gin.Default()
	router.GET("/:slug", controller.Redirect)

	slug := "abc123"
	mockUC.On("Resolve", slug).Return(&domains.URL{
		ID:       slug,
		Original: "https://github.com/stretchr/testify",
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/abc123", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusFound, resp.Code, "Expected status code 302 Found")
	assert.Equal(t, "https://github.com/stretchr/testify", resp.Header().Get("Location"), "Expected redirect to the original URL")
}
