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
				Expect(err).To(MatchError("unrecognized command"))
			})
		})

		Context("When given a SET command to parse", func() {
			Context("when given valid arguements", func() {
				It("returns an executable command", func() {
					command, err := parser.Parse("SET a 1")

					Expect(err).NotTo(HaveOccurred())
					Expect(command).To(Equal(&cmd.Set{
						Key:   "a",
						Value: "1",
					}))
				})
			})

			Context("when given too many arguements", func() {
				It("returns an error", func() {
					_, err := parser.Parse("SET a 1 12 a 9 65")

					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("wrong number of arguements to command"))
				})
			})

			Context("when given too few arguements", func() {
				It("returns an error", func() {
					_, err := parser.Parse("SET")

					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("wrong number of arguements to command"))
				})
			})
		})

		Context("When given a GET command to parse", func() {
			Context("when given valid arguements", func() {
				It("returns an executable command", func() {
					command, err := parser.Parse("GET a")

					Expect(err).NotTo(HaveOccurred())
					Expect(command).To(Equal(&cmd.Get{
						Key: "a",
					}))
				})
			})

			Context("when given too many arguements", func() {
				It("returns an error", func() {
					_, err := parser.Parse("GET a 1")

					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("wrong number of arguements to command"))

				})
			})

			Context("when given too few arguements", func() {
				It("returns an error", func() {
					_, err := parser.Parse("GET")

					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("wrong number of arguements to command"))
				})
			})
		})
	})
})
