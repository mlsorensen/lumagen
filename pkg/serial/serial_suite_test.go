package serial_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSerial(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Serial Suite")
}
