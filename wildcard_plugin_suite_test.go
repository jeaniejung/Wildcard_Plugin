//'go test' = 'gingko'
package main

import (
	//"github.com/cloudfoundry/cli/plugin/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestWildcardPlugin(t *testing.T) { //This is what is run when 'go test' or ginko 
	//When ginkgo test fails, ginkgo's Fail(description string) function gets called
	//This function is passed to Gomega using RegisterFailHandler, the sole connection point b/w Ginkgo and Gomega.
	RegisterFailHandler(Fail)
	//RunSpace(t *testing.T, suiteDescription string) tells Ginkgo to start the test suite.
	//When any spec fails, testing.T is automatically failed by Ginkgo
	RunSpecs(t, "WildcardPlugin Suite")
}
