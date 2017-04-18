package stdio_test

import (
	"github.com/zjohl/go-simple-database-challenge/stdio"

	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stdio Socket", func() {
	Describe("#BufferInput", func() {
		var socket stdio.Socket

		BeforeEach(func() {
			socket = stdio.Socket{
				In: make(chan string),
			}
		})

		Context("when passed stdin input", func() {
			It("sends them to its channel", func() {
				socket.BufferInput()

				fmt.Fprint(GinkgoWriter, "input")

				Eventually(socket.In).Should(Receive(Equal("input")))
			})
		})

	})

	Describe("#BufferOutput", func() {
		var socket stdio.Socket

		BeforeEach(func() {
			socket = stdio.Socket{
				Out: make(chan string),
			}
		})

		AfterEach(func() {

		})

		Context("when passed stdin input", func() {
			It("sends them to its channel", func() {
				stdout := make(chan string)

				go func() {
					var output string
					fmt.Scanln(&output)
					stdout <- output
				}()

				socket.BufferOutput()

				socket.Out <- "output"

				Eventually(stdout).Should(Receive(Equal("output")))
			})
		})
	})
})
