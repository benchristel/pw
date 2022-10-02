package pw

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// NOT SECURE. Do not use outside tests.
var cheapestscrypt = Scryptor{Complexity: 1}
var cheapscrypt = Scryptor{Complexity: 2}

var _ = Describe("a generated password for a site", func() {
	pw := SitePassword{
		Master: "master password",
		Salt: "the salt",
		Domain: "example.com",
		DeriveKey: cheapestscrypt.Key,
	}

	It("is 20 characters long", func() {
		Expect(len(pw.String())).To(Equal(20))
	})

	It("contains a capital letter", func() {		
		Expect(pw).To(MatchRegexp("[A-Z]"))
	})

	It("contains a digit", func() {
		Expect(pw).To(MatchRegexp("[0-9]"))
	})

	It("contains a symbol", func() {
		// This is just to satisfy websites' bogus password constraints.
		// There is plenty of entropy without choosing a random symbol.
		Expect(pw).To(ContainSubstring("!"))
	})

	It("is based on the master password", func() {
		copy := pw
		copy.Master = "different"
		Expect(copy.String()).NotTo(Equal(pw.String()))
	})

	It("is based on the salt", func() {
		copy := pw
		copy.Salt = "different"
		Expect(copy.String()).NotTo(Equal(pw.String()))
	})

	It("is based on the domain", func() {
		copy := pw
		copy.Domain = "different"
		Expect(copy.String()).NotTo(Equal(pw.String()))
	})

	It("is based on the encryption algorithm", func() {
		copy := pw
		copy.DeriveKey = cheapscrypt.Key
		Expect(copy.String()).NotTo(Equal(pw.String()))
	})
})
