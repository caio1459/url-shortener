package usecases

import (
	"errors"
	"testing"
	"url-shortener/internal/domains"
)

// Mock do URLRepository
type mockURLRepo struct {
	savedURL   *domains.URL
	saveErr    error
	findResult *domains.URL
	findErr    error
}

func (m *mockURLRepo) Save(url *domains.URL) error {
	m.savedURL = url
	return m.saveErr
}
func (m *mockURLRepo) FindByID(id string) (*domains.URL, error) {
	return m.findResult, m.findErr
}

func TestShorten_Success(t *testing.T) {
	mockRepo := &mockURLRepo{}
	uc := NewURLUsecase(mockRepo)

	original := "https://example.com"
	expire := 10

	slug, err := uc.Shorten(original, expire)
	if err != nil {
		t.Fatalf("esperava nil, obteve erro: %v", err)
	}
	if slug == "" {
		t.Error("esperava slug não vazio")
	}
	if mockRepo.savedURL == nil {
		t.Fatal("esperava URL salva no repositório")
	}
	if mockRepo.savedURL.Original != original {
		t.Errorf("esperava original %s, obteve %s", original, mockRepo.savedURL.Original)
	}
	if mockRepo.savedURL.ExpireAt == nil {
		t.Error("esperava ExpireAt definido")
	}
}

func TestShorten_NoExpire(t *testing.T) {
	mockRepo := &mockURLRepo{}
	uc := NewURLUsecase(mockRepo)

	slug, err := uc.Shorten("https://test.com", 0)
	if err != nil {
		t.Fatalf("esperava nil, obteve erro: %v", err)
	}
	if slug == "" {
		t.Error("esperava slug não vazio")
	}
	if mockRepo.savedURL.ExpireAt != nil {
		t.Error("esperava ExpireAt nil")
	}
}

func TestShorten_RepoError(t *testing.T) {
	mockRepo := &mockURLRepo{saveErr: errors.New("erro repo")}
	uc := NewURLUsecase(mockRepo)

	_, err := uc.Shorten("https://fail.com", 5)
	if err == nil {
		t.Error("esperava erro do repositório")
	}
}

func TestResolve_Success(t *testing.T) {
	expected := &domains.URL{ID: "abc123", Original: "https://x.com"}
	mockRepo := &mockURLRepo{findResult: expected}
	uc := NewURLUsecase(mockRepo)

	url, err := uc.Resolve("abc123")
	if err != nil {
		t.Fatalf("esperava nil, obteve erro: %v", err)
	}
	if url != expected {
		t.Error("esperava URL encontrada")
	}
}

func TestResolve_NotFound(t *testing.T) {
	mockRepo := &mockURLRepo{findErr: errors.New("não encontrado")}
	uc := NewURLUsecase(mockRepo)

	_, err := uc.Resolve("notfound")
	if err == nil {
		t.Error("esperava erro para slug inexistente")
	}
}
