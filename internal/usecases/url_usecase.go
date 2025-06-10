package usecases

import (
	"time"
	"url-shortener/internal/domains"
	"url-shortener/internal/interfaces"
	"url-shortener/pkg/sluggen"
)

type urlUsecase struct {
	repo interfaces.URLRepository
}

func NewURLUsecase(repo interfaces.URLRepository) interfaces.URLUsecase {
	return &urlUsecase{repo: repo}
}

// Resolve implements interfaces.URLUsecase.
func (u *urlUsecase) Resolve(slug string) (*domains.URL, error) {
	return u.repo.FindByID(slug)
}

// Shorten implements interfaces.URLUsecase.
func (u *urlUsecase) Shorten(original string, expireMinutes int) (string, error) {
	ID := sluggen.GenerateSlug(6)

	var expireAt *time.Time
	if expireMinutes > 0 {
		// Define a data de expiração com base no tempo atual e na duração especificada
		t := time.Now().Add(time.Duration(expireMinutes) * time.Minute)
		expireAt = &t
	}

	url := &domains.URL{
		ID:        ID,
		Original:  original,
		CreatedAt: time.Now(),
		ExpireAt:  expireAt,
	}

	return ID, u.repo.Save(url)
}
