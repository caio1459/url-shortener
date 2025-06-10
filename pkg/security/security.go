package security

import "golang.org/x/crypto/bcrypt"

//Criptografa a senha
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//Compara a senha Criptografada com a senha que está sendo mandada
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
