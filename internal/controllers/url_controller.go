package controllers

import (
	"net/http"
	"time"
	"url-shortener/internal/dtos"
	"url-shortener/internal/usecases"

	"github.com/gin-gonic/gin"
)

type URLContoller struct {
	uc *usecases.URLUsecase
}

func NewURLContoller(uc *usecases.URLUsecase) *URLContoller {
	return &URLContoller{uc: uc}
}

func (c *URLContoller) Shorten(ctx *gin.Context) {
	var req dtos.URLRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	slug, err := c.uc.Shorten(req.URL, req.ExpireInMinutes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"short_url": ctx.Request.Host + "/" + slug})
}

func (c *URLContoller) Redirect(ctx *gin.Context) {
	slug := ctx.Param("slug")
	url, err := c.uc.Resolve(slug)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	if url.ExpireAt != nil && url.ExpireAt.Before(time.Now()) {
		ctx.JSON(http.StatusGone, gin.H{"error": "URL expired"})
		return
	}

	ctx.Redirect(http.StatusFound, url.Original)
}
