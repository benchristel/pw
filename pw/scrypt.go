package pw

import (
	"log"

	"golang.org/x/crypto/scrypt"
)

type Scryptor struct {
	Complexity uint8
}

func (s Scryptor) Key(plaintext string, salt []byte) []byte {
	dk, err := scrypt.Key([]byte(plaintext), salt, 1<<s.Complexity, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	return dk
}
