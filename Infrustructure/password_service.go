package infrastructure

import "golang.org/x/crypto/bcrypt"

type BcryptService struct{}

func (b *BcryptService) Hash(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(bytes), err
}

func (b *BcryptService) Compare(hash, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}
