package main_test


import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	"os"
	"log"
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Simple Database", func() {
	var (
		session       *Session
		stdinPipe     io.WriteCloser
		err           error
	)

	BeforeEach(func() {
		simpleDatabase := exec.Command(simpleDatabasePath)
		stdinPipe, err = simpleDatabase.StdinPipe()
		Expect(err).ToNot(HaveOccurred())

		session, err = Start(simpleDatabase, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("when given a valid set of data commands", func() {
		It("executes those commands and returns the expected output", func() {
			printFile(stdinPipe, "fixtures/data-command-1.txt")

			Eventually(session).Should(Exit())
			Expect(session.Out).To(gbytes.Say("10"))
			Expect(session.Out).To(gbytes.Say("NULL"))

		})
	})
})

func printFile(pipe io.WriteCloser, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Fprint(pipe, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer pipe.Close()
}
