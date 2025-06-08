package interfaces

import "url-shortener/internal/domains"

type URLRepository interface {
	Save(url *domains.URL) error
	FindByID(id string) (*domains.URL, error)
}

type URLUsecase interface {
	Shorten(original string, expireMinutes int) (string, error)
	Resolve(slug string) (*domains.URL, error)
}
