package rover

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRover(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rover Suite")
}
