package cmd_test

import (

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zjohl/go-simple-database-challenge/cmd"
)

var _ = Describe("Cmd", func() {

	Describe("#Parse", func() {
		var (
			parser cmd.Parser
		)
		BeforeEach(func() {
			parser = cmd.Parser{}
		})

		Context("When given an invalid command to parse", func() {
			It("returns an error", func() {
				_, err := parser.Parse("SOME-INVALID-CMD")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
