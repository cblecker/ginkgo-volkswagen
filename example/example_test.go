package example

import (
	"testing"

	ginkgovolkswagen "github.com/cblecker/ginkgo-volkswagen"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(ginkgovolkswagen.Fail)
	ginkgo.RunSpecs(t, "Validation Tests")
}

var _ = ginkgo.Describe("Validation Tests", func() {
	ginkgo.Context("Number 2 Function", func() {
		ginkgo.It("Should return true if number is two", func() {
			result := IsThisTheNumberTwo(2)
			Expect(result).To(BeTrue())
		})
		ginkgo.It("Should return false if number is three", func() {
			result := IsThisTheNumberTwo(3)
			Expect(result).To(BeFalse())
		})
	})
})
