package sluggen

import (
	"math/rand"
)

// Contém os caracteres permitidos para o slug
var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// Gera uma string aleatória (slug) com o tamanho especificado
func GenerateSlug(length int) string {
	b := make([]rune, length) // cria um slice de runes com o tamanho desejado
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))] // seleciona um caractere aleatório do charset
	}
	return string(b) // retorna o slug como string
}
