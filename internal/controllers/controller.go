package controllers

import "url-shortener/internal/interfaces"

type Controllers struct {
	URL *URLController
	// User *UserController
}

func NewControllers(urlUsecase interfaces.URLUsecase) *Controllers {
	return &Controllers{
		URL: NewURLController(urlUsecase),
		// inicialize outros controllers aqui
	}
}
