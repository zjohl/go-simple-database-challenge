package stdio_test

import (
	"github.com/zjohl/go-simple-database-challenge/stdio"

	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stdio Socket", func() {
	FDescribe("#BufferInput", func() {
		var socket stdio.Socket

		BeforeEach(func() {
			socket = stdio.Socket{
				In: make(chan string),
			}
		})

		AfterEach(func() {

		})

		Context("when passed stdin input", func() {
			It("sends them to its channel", func() {
				socket.BufferInput()

				fmt.Fprint(GinkgoWriter, "input")

				Eventually(socket.In).Should(Receive())
			})
		})

	})

	Describe("#BufferOutput", func() {
		It("", func() {
			Fail("Not implemented")
		})
	})
})
