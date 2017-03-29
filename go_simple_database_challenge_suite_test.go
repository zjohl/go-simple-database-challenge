package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"testing"
)

func TestGoSimpleDatabaseChallenge(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoSimpleDatabaseChallenge Suite")
}

var simpleDatabasePath string

var _ = BeforeSuite(func() {
	var err error
	simpleDatabasePath, err = Build("github.com/zjohl/go-simple-database-challenge")
	Expect(err).ShouldNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	CleanupBuildArtifacts()
})