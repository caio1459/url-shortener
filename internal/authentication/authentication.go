package authentication

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint64, email string) (string, error) {
	//Gera permissoes do token
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["user_id"] = id
	permissions["email"] = email //
	// Cria o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	// Assina o token
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

// Extrai o token do header
func extractToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	// Verifica se o token possui mais de uma palavra e pega sempre a segunda. EX: Bearer <token>
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

// retorna a chave de verificação do token
func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	// Verifica se o método de assinatura utilizado pertence à família correta
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %s", token.Header["alg"])
	}
	return []byte(os.Getenv("SECRET_KEY")), nil
}

func ExtractUserInfo(c *gin.Context) (uint64, string, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, "", err
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extrai o user_id e faz a conversão para uint64
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["user_id"]), 10, 64)
		if err != nil {
			return 0, "", err
		}

		// Extrai o user_type e faz a conversão para uint8
		email, ok := permissions["email"].(string)
		if !ok {
			return 0, "", errors.New("tipo de usuário inválido no token")
		}
		return userID, email, nil
	}
	return 0, "", errors.New("token inválido")
}
