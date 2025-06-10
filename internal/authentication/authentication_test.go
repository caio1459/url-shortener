package authentication

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAndExtractUserInfo_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Setenv("SECRET_KEY", "secreto")

	// Gera token válido
	token, err := GenerateToken(1, "user@email.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Cria contexto de teste do Gin com header Authorization
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Extrai as informações do token
	userID, email, err := ExtractUserInfo(c)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), userID)
	assert.Equal(t, "user@email.com", email)
}
func TestExtractUserInfo_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Setenv("SECRET_KEY", "secreto")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	userID, email, err := ExtractUserInfo(c)
	assert.Error(t, err)
	assert.Equal(t, uint64(0), userID)
	assert.Equal(t, "", email)
}

func TestExtractUserInfo_MissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Sem header Authorization
	userID, email, err := ExtractUserInfo(c)
	assert.Error(t, err)
	assert.Equal(t, uint64(0), userID)
	assert.Equal(t, "", email)
}
