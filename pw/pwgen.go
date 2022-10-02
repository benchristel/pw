package pw

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
)

type SitePassword struct {
	Master    string
	Salt      string
	Domain    string
	DeriveKey func(string, []byte) []byte
}

func (sp SitePassword) String() string {
	key := sp.DeriveKey(sp.Master+sp.Domain, hash([]byte(sp.Salt)))
	return base64.StdEncoding.EncodeToString(hash(key))[0:18] + "1!"
}

func hash(b []byte) []byte {
	h := sha256.New()
	_, err := h.Write(b)
	if err != nil {
		// Writing to a hash should never fail.
		// If it does, something has gone very wrong.
		log.Fatal(err)
	}
	return h.Sum(nil)
}