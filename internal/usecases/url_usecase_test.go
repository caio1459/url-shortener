package usecases

import (
	"errors"
	"testing"
	"url-shortener/internal/domains"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do repositório
type mockURLRepository struct {
	mock.Mock
}

func (m *mockURLRepository) Save(url *domains.URL) error {
	args := m.Called(url)
	return args.Error(0)
}

func (m *mockURLRepository) FindByID(id string) (*domains.URL, error) {
	args := m.Called(id)
	if url, ok := args.Get(0).(*domains.URL); ok {
		return url, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestShorten(t *testing.T) {
	mockRepo := new(mockURLRepository)
	uc := NewURLUsecase(mockRepo)

	original := "https://github.com/stretchr/testify"
	expireInMinutes := 1

	// Espera que o método Save seja chamado com qualquer *URL
	mockRepo.On("Save", mock.AnythingOfType("*domains.URL")).Return(nil)
	slug, err := uc.Shorten(original, expireInMinutes)

	assert.NoError(t, err, "Expected no error when shortening URL")
	assert.Len(t, slug, 6, "Expected slug to be 6 characters long")
	mockRepo.AssertExpectations(t)
}

func TestResolve(t *testing.T) {
	mockRepo := new(mockURLRepository)
	uc := NewURLUsecase(mockRepo)

	slug := "abc123"
	expectedURL := &domains.URL{
		ID:       slug,
		Original: "https://github.com/stretchr/testify",
	}
	// Espera que FindByID retorne uma URL válida
	mockRepo.On("FindByID", slug).Return(expectedURL, nil)

	result, err := uc.Resolve(slug)

	assert.NoError(t, err, "Expected no error when resolving URL")
	assert.Equal(t, expectedURL, result, "Expected resolved URL to match the expected URL")
	mockRepo.AssertExpectations(t)
}

func TestResolve_NotFound(t *testing.T) {
	mockRepo := new(mockURLRepository)
	uc := NewURLUsecase(mockRepo)

	slug := "notfound123"
	mockRepo.On("FindByID", slug).Return(nil, errors.New("not found"))

	result, err := uc.Resolve(slug)

	assert.Error(t, err, "Expected error when resolving non-existent URL")
	assert.Nil(t, result, "Expected result to be nil when URL is not found")
}
