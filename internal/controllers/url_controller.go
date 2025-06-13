package controllers

import (
	"net/http"
	"time"
	"url-shortener/internal/dtos"
	"url-shortener/internal/interfaces"

	"github.com/gin-gonic/gin"
)

type URLController struct {
	uc interfaces.URLUsecase
}

func NewURLController(uc interfaces.URLUsecase) *URLController {
	return &URLController{uc: uc}
}

func (c *URLController) Shorten(ctx *gin.Context) {
	var req dtos.URLRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	slug, err := c.uc.Shorten(req.URL, req.ExpireInMinutes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"short_url": ctx.Request.Host + "/" + slug})
}

func (c *URLController) Redirect(ctx *gin.Context) {
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
