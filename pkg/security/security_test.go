package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	password := "123456"
	hashed, err := Hash(password)
	assert.NoError(t, err, "Hash should not return error")
	assert.NotEmpty(t, hashed, "Hash should return a non-empty hash")

	// O hash gerado deve ser válido para a senha original
	err = CheckPassword(password, string(hashed))
	assert.NoError(t, err, "CheckPassword should succeed with correct password")
}

func TestCheckPassword(t *testing.T) {
	password := "123456"
	hashed, _ := Hash(password)

	// Senha correta
	err := CheckPassword(password, string(hashed))
	assert.NoError(t, err, "CheckPassword should succeed with correct password")

	// Senha incorreta
	err = CheckPassword("aaaaaaaaaa", string(hashed))
	assert.Error(t, err, "CheckPassword should fail with incorrect password")
}
