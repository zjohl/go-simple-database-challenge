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

		Context("When given an valid command to parse", func() {
			Context("when given valid arguements", func() {
				It("returns an executable command", func() {
					command, err := parser.Parse("SET a 1")

					Expect(err).NotTo(HaveOccurred())
					Expect(command).To(Equal(&cmd.Set{
						Key: "a",
						Value:"1",
					}))
				})
			})

			Context("when given invalid arguements", func() {
				It("returns an error", func() {
					_, err := parser.Parse("SET")

					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})
