package usecases

import (
	"time"
	"url-shortener/internal/domains"
	"url-shortener/internal/interfaces"
	"url-shortener/pkg/sluggen"
)

type URLUsecase struct {
	repo interfaces.URLRepository
}

func NewURLUsecase(repo interfaces.URLRepository) *URLUsecase {
	return &URLUsecase{repo: repo}
}

// Cria um novo URL encurtado
func (uc *URLUsecase) Shorten(original string, expireInMinutes int) (string, error) {
	ID := sluggen.GenerateSlug(6)

	var expireAt *time.Time
	if expireInMinutes > 0 {
		// Define a data de expiração com base no tempo atual e na duração especificada
		t := time.Now().Add(time.Duration(expireInMinutes) * time.Minute)
		expireAt = &t
	}

	url := &domains.URL{
		ID:        ID,
		Original:  original,
		CreatedAt: time.Now(),
		ExpireAt:  expireAt,
	}

	return ID, uc.repo.Save(url)
}

func (uc *URLUsecase) Resolve(slug string) (*domains.URL, error) {
	return uc.repo.FindByID(slug)
}
