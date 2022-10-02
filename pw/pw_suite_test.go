package pw

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPw(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "pw suite")
}
